package snapshot

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	conf "casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/pkg/gogs"
	"casicloud.com/ylops/marco/pkg/log"
	"casicloud.com/ylops/marco/pkg/models"
	"casicloud.com/ylops/marco/pkg/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	ssh2 "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

var (
	descFormat = "id:%s,desc:%s"
)

//SnapShotOption 初始化Git需要的相关配置
type SnapShotOption struct {
	// author for access
	Username string
	// key pem for sshkey
	PEMBytes []byte
	PEMPass  string
	// password for key
	Password    string
	APIURL      string
	AccessToken string
}

func (o *SnapShotOption) getSSHAuth() (transport.AuthMethod, error) {
	auth, err := ssh2.NewPublicKeys(o.Username, []byte(o.PEMBytes), o.PEMPass)
	auth.HostKeyCallbackHelper = ssh2.HostKeyCallbackHelper{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return auth, err
}

func (o *SnapShotOption) getHTTPAuth() (transport.AuthMethod, error) {
	return &http.BasicAuth{
		Username: o.Username,
		Password: o.Password,
	}, nil
}

func (o *SnapShotOption) createClient() *gogs.Client {
	return gogs.NewClient(
		o.APIURL,
		o.AccessToken,
	)
}

type WorkTree struct {
	LocalRepo  string
	TargetDir  string
	RestoreDir string
}

//HasRepo 判断是否以及有git库
func (wt *WorkTree) HasRepo() bool {
	gitDir := filepath.Join(wt.LocalRepo, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return false
	}

	return true
}

func (wt *WorkTree) CleanGitDir() error {
	gitDir := filepath.Join(wt.LocalRepo, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return err
	}
	return os.RemoveAll(gitDir)
}

//GitSnapshot 通过git来管理快照并同步到远程仓库且以目录为单位进行快照的创建与恢复
// 目前只支持 ssh和git的协议
type GitSnapshot struct {
	Cluster *models.Cluster
	Config  *conf.Config
	opt     *SnapShotOption
	wt      *WorkTree
	client  *gogs.Client
	mt      sync.Mutex
}

func NewGitSnapShot(cluster *models.Cluster,
	config *conf.Config,
	opt *SnapShotOption) *GitSnapshot {
	client := opt.createClient()

	snapDir := config.GetSnapshotDir()
	localReop := filepath.Join(snapDir, "local")
	targetDir := config.GetPrefix()
	restorDir := filepath.Join(snapDir, "restore")
	wt := &WorkTree{
		LocalRepo:  localReop,
		TargetDir:  targetDir,
		RestoreDir: restorDir,
	}
	g := &GitSnapshot{
		client:  client,
		Cluster: cluster,
		Config:  config,
		wt:      wt,
		opt:     opt,
	}

	return g
}

func (g *GitSnapshot) Take(info *Snapinfo) error {
	// 检查远程Remote是否存在,不能存在就调用API创建一个信息
	remoteURL, err := g.GetRemoteURL()
	if err == gogs.ErrNotFound {
		// 调用GitServerOperator创建一个新的
		_, err := g.client.CreateRepo(gogs.CreateRepoOption{
			Name:        g.Cluster.ID,
			Description: fmt.Sprintf("macro cluster: %s", g.Cluster.ID),
			Private:     true,
		})

		if err != nil {
			return err
		}
	}
	// 初始化Git
	auth, err := g.opt.getHTTPAuth()
	if err != nil {
		return err
	}
	//清空localrepo
	os.RemoveAll(g.wt.LocalRepo)

	var repo *git.Repository
	repo, err = git.PlainClone(g.wt.LocalRepo, false, &git.CloneOptions{
		Auth:       auth,
		NoCheckout: true,
		Progress:   os.Stdout,
		URL:        remoteURL,
	})

	if err == transport.ErrEmptyRemoteRepository {
		repo, err = git.PlainInit(g.wt.LocalRepo, false)
		if err != nil {
			return err
		}
		_, err = repo.CreateRemote(&config.RemoteConfig{
			Name: "origin",
			URLs: []string{remoteURL},
		})
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	//复制
	err = utils.Copy(g.wt.TargetDir, g.wt.LocalRepo)
	if err != nil {
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	status, err := w.Status()
	for n := range status {
		log.Debugf("repo add %v\n", n)
		w.Add(n)
	}

	desc := fmt.Sprintf(descFormat, info.ID, info.Description)
	commit, err := w.Commit(desc, &git.CommitOptions{
		Author: &object.Signature{
			Name:  info.Author,
			Email: info.Email,
			When:  info.CreatedAt,
		},
	})
	if err != nil {
		return err
	}
	_, err = repo.CommitObject(commit)
	if err != nil {
		return err
	}
	err = repo.Push(&git.PushOptions{
		Auth: auth,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *GitSnapshot) GetRemoteURL() (string, error) {
	// check exists
	repo, err := g.client.GetRepo(g.opt.Username, g.Cluster.ID)
	if err != nil {
		return "", err
	}
	return repo.CloneURL, nil
}

func (g *GitSnapshot) Restore(id string) error {

}

func (g *GitSnapshot) List() (Snapinfos, error) {
	remoteURL, err := g.GetRemoteURL()
	fmt.Println(remoteURL)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

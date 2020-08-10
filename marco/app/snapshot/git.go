package snapshot

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	conf "casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/pkg/gogs"
	"casicloud.com/ylops/marco/pkg/models"
	"github.com/go-git/go-git/v5/plumbing/transport"
	ssh2 "github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

//SnapShotOption 初始化Git需要的相关配置
type SnapShotOption struct {
	// author for access
	Username string
	// key pem for sshkey
	PEMBytes []byte
	// password for key
	Password    string
	APIURL      string
	AccessToken string
}

func (o *SnapShotOption) getAuth() (transport.AuthMethod, error) {
	return ssh2.NewPublicKeys(o.Username, []byte(o.PEMBytes), o.Password)
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
	fmt.Println(remoteURL)
	return nil
}

func (g *GitSnapshot) GetRemoteURL() (string, error) {
	// check exists
	repo, err := g.client.GetRepo(g.opt.Username, g.Cluster.ID)
	if err != nil {
		return "", err
	}
	return repo.SSHURL, nil
}

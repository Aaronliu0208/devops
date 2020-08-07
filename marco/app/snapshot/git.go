package snapshot

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport"
	ssh2 "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
)

//GitOption 初始化Git需要的相关配置
type GitOption struct {
	// author for access
	Username string
	// key pem for sshkey
	PEMBytes []byte
	// password for key
	Password string
	// git remote repo url
	// git@yunludev3.htyunlu.com:shanyou/marco-snapshot.git
	RemoteURL string
	RepoName  string
}

func (o *GitOption) getAuth() (transport.AuthMethod, error) {
	return ssh2.NewPublicKeys(o.Username, []byte(o.PEMBytes), o.Password)
}

func (o *GitOption) checkRemote() error {
	remote := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{o.RemoteURL},
	})

	auth, err := o.getAuth()
	if err != nil {
		return err
	}
	_, err = remote.List(&git.ListOptions{
		Auth: auth,
	})

	return err
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
	opt *GitOption
	wt  *WorkTree
	gso GitServerOperator
}

func (g *GitSnapshot) Take() (*Snapinfo, error) {
	// 检查远程Remote是否存在,不能存在就调用API创建一个信息
	err := g.opt.checkRemote()
	if err == git.ErrRemoteNotFound {
		// 调用GitServerOperator创建一个新的
	}
	// 初始化Git

	return nil, nil
}

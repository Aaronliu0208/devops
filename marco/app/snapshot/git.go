package snapshot

import (
	"os"
	"path/filepath"

	conf "casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/pkg/gogs"
	"casicloud.com/ylops/marco/pkg/models"
	"github.com/go-git/go-git/v5"
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
}

func (g *GitSnapshot) Take(info *Snapinfo) error {
	// 检查远程Remote是否存在,不能存在就调用API创建一个信息
	var err error
	if err == git.ErrRemoteNotFound {
		// 调用GitServerOperator创建一个新的
	}
	// 初始化Git

	return nil
}

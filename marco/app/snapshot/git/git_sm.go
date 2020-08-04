package git

//GitSnapshotManager 通过git管理文件快照
type GitSnapshotManager struct {
	//本地git工作目录
	LocalDir string
	//远程地址
	RemoteRepo string
}

func New(localdir, remoteRepo string) *GitSnapshotManager {
	return nil
}
func (gs *GitSnapshotManager) Init() error {
	return nil
}

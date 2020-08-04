package snapshot

//SnapshotManager with create or restore snapshot
type SnapshotManager interface {
	Create() (*Snapshot, error)
	List() (Snapshots, error)
	DeleteByID(id string) error
	GetByID(id string) (*Snapshot, error)
	Restore(sp *Snapshot) error
}

package snapshot

type Snapshot interface {
	Take() (*Snapinfo, error)
	Restore(*Snapinfo) error
	List() (Snapinfos, error)
}

// GitServerOperator operator that manage remote repo for create and add pubkey
type GitServerOperator interface {
	CreateRepo(name string) error
	AddPublicKey(pubKey []byte) error
	IsPubKeyExists(pubKey []byte) bool
}

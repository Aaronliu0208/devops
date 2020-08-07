package snapshot

type RemoteGitClient struct {
}

func (rg *RemoteGitClient) CreateRepo(name string) error {
	return nil
}

func (rg *RemoteGitClient) AddPublicKey(pubKey []byte) error {
	return nil
}

func (rg *RemoteGitClient) IsPubKeyExists(pubKey []byte) bool {
	return false
}

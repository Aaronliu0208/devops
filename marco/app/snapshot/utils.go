package snapshot

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
)

func CheckRemote(remoteURL string, o *SnapShotOption) error {
	remote := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteURL},
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

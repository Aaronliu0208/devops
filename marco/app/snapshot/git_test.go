package snapshot

import (
	"fmt"
	"testing"

	"casicloud.com/ylops/marco/pkg/gogs"
)

func TestGogsClient(t *testing.T) {
	client := gogs.NewClient(
		"http://172.20.4.45:10080",
		"ea4ede54549411cb1a73e07248cc143265b4496d",
	)
	defer client.DeleteRepo("test", "test1")

	keys, err := client.ListMyPublicKeys()
	if err != nil {
		t.Fatal(err)
	}

	for _, k := range keys {
		fmt.Printf("%v\n", k)
	}

	_, err = client.CreateRepo(gogs.CreateRepoOption{
		Name:        "test1",
		Description: "test repo for gogs client create",
		Private:     true,
	})

	if err != nil {
		t.Fatal(err)
	}

	repos, err := client.ListMyRepos()
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range repos {
		fmt.Printf("%v\n", r)
	}

}

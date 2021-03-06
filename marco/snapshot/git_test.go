package snapshot

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	conf "casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/pkg/gogs"
	"casicloud.com/ylops/marco/pkg/models"
	"casicloud.com/ylops/marco/pkg/utils"
	"github.com/stretchr/testify/assert"
)

//事先需要在gogs上建好clusterid对应的repo
var mockClusterID = "xxx-xxx"

func getGitSnapshot() *GitSnapshot {
	cluster := &models.Cluster{
		ID: mockClusterID,
	}
	_, filename, _, _ := runtime.Caller(0)
	wd := filepath.Dir(filename)
	config := &conf.Config{
		Workspace: filepath.Join(wd, "../bin"),
	}

	s := fmt.Sprintf("%s/git/test/id_rsa", wd)
	sshKey, _ := ioutil.ReadFile(s)
	opt := &SnapShotOption{
		Username:    "test",
		APIURL:      "http://172.20.4.45:10080",
		AccessToken: "ea4ede54549411cb1a73e07248cc143265b4496d",
		PEMBytes:    sshKey,
		PEMPass:     "1234.asd",
		Password:    "1234.asd",
	}
	return NewGitSnapShot(cluster, config, opt)
}
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

func TestGetRemoteURL(t *testing.T) {
	g := getGitSnapshot()

	url, err := g.GetRemoteURL()
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, strings.Contains(url, "http://172.20.4.45:10080/test/xxx-xxx.git"))
}

func TestGetRemoteURLFail(t *testing.T) {
	g := getGitSnapshot()
	g.Cluster.ID = utils.NewID()
	_, err := g.GetRemoteURL()
	assert.Equal(t, err, gogs.ErrNotFound)
}

func TestJoinTargetPath(t *testing.T) {
	g := getGitSnapshot()
	defer os.RemoveAll(g.wt.LocalRepo)
	err := utils.Copy(g.wt.TargetDir, g.wt.LocalRepo)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTaskSnap(t *testing.T) {
	g := getGitSnapshot()
	defer os.RemoveAll(g.wt.LocalRepo)
	info := &Snapinfo{
		ID:          utils.NewID(),
		CreatedAt:   time.Now(),
		Author:      "shanyou",
		Email:       "shanyou@htyunwang.com",
		Description: "create snapshot",
	}

	err := g.Take(info)

	if err != nil {
		t.Fatal(err)
	}
}

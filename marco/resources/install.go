package resources

import (
	"context"
)

//RestyInstaller install resty on disk
type RestyInstaller struct {
	WorkDir      string
	Prefix       string
	BuildOptions []string
}

func (r *RestyInstaller) install(ctx context.Context) {

}

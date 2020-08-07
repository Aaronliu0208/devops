package snapshot

import (
	"time"
)

type Snapinfo struct {
	ID          string
	CreatedAt   time.Time
	Description string
	Author      string
	// snapshot的元信息，包括git地址，commitid, author等
	Meta map[string]string
}

type Snapinfos []*Snapinfo

type Snapshot interface {
	Take(info *Snapinfo) error
	Restore(id string) error
	List() (Snapinfos, error)
}

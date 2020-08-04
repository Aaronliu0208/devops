package snapshot

import "time"

type Snapshot struct {
	ID          string
	CreatedAt   time.Time
	Description string
	// snapshot的元信息，包括git地址，commitid, author等
	Meta map[string]string
}

type Snapshots []*Snapshot

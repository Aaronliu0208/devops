package models

// Certificate respresent nginx ssl certificate with sni
type Certificate struct {
	ID      string `json:"id"`
	SNI     string `json:"sni"`
	Content []byte `json:"content"`
}

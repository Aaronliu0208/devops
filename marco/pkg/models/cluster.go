package models

import (
	"time"

	"casicloud.com/ylops/marco/pkg/nginx"
)

//Cluster represent cluster of marco instance
// 一个cluster代表一组站点的集合并且享有统一的Nginx配置，一组相同upstreams
type Cluster struct {
	ID          string `json:"id"`
	Name        string
	Description string
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Config      *nginx.Config
	Sites       []Site
	Upstreams   []Upstream
}

//GenerateConfig with cluster
func (c *Cluster) GenerateConfig() (string, error) {
	emptyblk := &nginx.EmptyBlock{}
	emptyblk.AddInterface(c.Config)
	return emptyblk.String(), nil
}

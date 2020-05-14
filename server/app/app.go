package app

import (
	"context"

	"casicloud.com/ylops/backend/config"
	"casicloud.com/ylops/backend/pkg/logger"
)

//Run 主程序入口
func Run(ctx context.Context) error {
	logger.Printf(ctx, "HTTP server is running at %s:%d.", config.C.HTTP.Host, config.C.HTTP.Port)
	return nil
}

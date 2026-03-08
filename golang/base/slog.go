package main

import (
	"log/slog"
	"os"
)

func slogSolution() {
	slog.Info("用户登录", "userID", 1001, "ip", "192.168.1.1")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Error("数据库连接失败",
		slog.String("db_name", "mysql_main"),
		slog.Int("retry_count", 3),
	)
}

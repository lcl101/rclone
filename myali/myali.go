package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/rclone/rclone/backend/all"
	"github.com/rclone/rclone/cmd"
	_ "github.com/rclone/rclone/cmd/all"
	_ "github.com/rclone/rclone/lib/plugin"
	"github.com/rclone/rclone/myali/config"
	"github.com/rclone/rclone/myali/webdav"
)

func parseParams(conf config.RcloneMountConfig) []string {
	cmd := make([]string, 0)
	cmd = append(cmd, "mount")
	if conf.Remote != "" {
		cmd = append(cmd, conf.Remote)
	}
	if conf.Mountpoint != "" {
		cmd = append(cmd, conf.Mountpoint)
	}
	if conf.CacheDir != "" {
		cmd = append(cmd, "--cache-dir", conf.CacheDir)
	}
	cmd = append(cmd, "--vfs-cache-mode", "writes")
	if conf.CacheMaxAge > 0 {
		var maxAge time.Duration
		maxAge = time.Duration(conf.CacheMaxAge) * time.Second
		fmt.Println("maxAge=", maxAge)
		cmd = append(cmd, "--vfs-cache-max-age", maxAge.String())
	}
	return cmd
}

func main() {
	if len(os.Args) < 2 {
		config.LoadConfig()
		// args := []string{"mount", "ali:/", "g:", "--cache-dir", "E:\\ali_tmp", "--vfs-cache-mode", "writes"}
		cmd.Root.SetArgs(parseParams(config.GetConfig().Mount))
		webdav.RunWebDav(config.GetConfig().Webdav)
		defer webdav.StopWebDav()
	}
	cmd.Main()
}

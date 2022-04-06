package webdav

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/rclone/rclone/myali/config"
	"github.com/rclone/rclone/myali/utils"
)

// AliWebDavConfig USAGE:
//    aliyundrive-webdav [FLAGS] [OPTIONS] --refresh-token <refresh-token>
// USAGE:
// aliyundrive-webdav [FLAGS] [OPTIONS] --refresh-token <refresh-token>
//
// FLAGS:
// -I, --auto-index    Automatically generate index.html
// -h, --help          Prints help information
// --no-trash      Delete file permanently instead of trashing it
// --read-only     Enable read only mode
// -V, --version       Prints version information
//
// OPTIONS:
// -W, --auth-password <auth-password>          WebDAV authentication password [env: WEBDAV_AUTH_PASSWORD=]
// -U, --auth-user <auth-user>                  WebDAV authentication username [env: WEBDAV_AUTH_USER=]
// --cache-size <cache-size>                Directory entries cache size [default: 1000]
// --cache-ttl <cache-ttl>                  Directory entries cache expiration time in seconds [default: 600]
// --domain-id <domain-id>                  Aliyun PDS domain id
// --host <host>                            Listen host [env: HOST=]  [default: 0.0.0.0]
// -p, --port <port>                            Listen port [env: PORT=]  [default: 8080]
// -S, --read-buffer-size <read-buffer-size>
// Read/download buffer size in bytes, defaults to 10MB [default: 10485760]
//
// -r, --refresh-token <refresh-token>          Aliyun drive refresh token [env: REFRESH_TOKEN=]
// --root <root>                            Root directory path [default: /]
// -w, --workdir <workdir>                      Working directory, refresh_token will be stored in there if specified

//CurrentDirectory 获取程序运行路径
// func CurrentDirectory11() string {
// 	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

// 	return OsPathConverter(dir)
// }

func RunWebDav(conf config.AliWebDavConfig) {
	params := parseParams(conf)
	command := exec.Command(utils.WebdavPath(), params...)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	log.Println(command.String())
	logFile, err := os.OpenFile(utils.WebdavLogsPath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Println("记录webdav执行日志文件打开失败 " + err.Error())
	} else {
		command.Stdout = logFile
		command.Stderr = logFile
	}
	command.Start()
}

func StopWebDav() error {

	if runtime.GOOS == "windows" {
		if runtime.GOOS == "windows" {
			cmd := exec.Command("taskkill.exe", "/F", "/T", "/IM", utils.WinAliyunDriveWebdav)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			log.Println(cmd.String())
			res, err := cmd.CombinedOutput()
			result := utils.ConvertToString(string(res), "gbk", "utf8")
			// result := string(res)
			log.Println(result)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			if strings.Contains(result, "错误") || strings.Contains(result, "exit") {
				return fmt.Errorf(result)
			}
		}
	}
	return nil
}

func parseParams(conf config.AliWebDavConfig) []string {
	cmd := make([]string, 0)
	if conf.AutoIndex == "Y" {
		cmd = append(cmd, "-I")
	}
	if conf.NoTrash == "Y" {
		cmd = append(cmd, "--no-trash")
	}
	if conf.ReadOnly == "Y" {
		cmd = append(cmd, "--read-only")
	}
	if len(conf.AuthPassword) > 0 {
		cmd = append(cmd, "-W", conf.AuthPassword)
	}
	if len(conf.AuthUser) > 0 {
		cmd = append(cmd, "-U", conf.AuthUser)
	}
	if len(conf.CacheSize) > 0 {
		cmd = append(cmd, "--cache-size", conf.CacheSize)
	}
	if len(conf.CacheTtl) > 0 {
		cmd = append(cmd, "--cache-ttl", conf.CacheTtl)
	}
	if len(conf.DomainId) > 0 {
		cmd = append(cmd, "--domain-id", conf.DomainId)
	}
	if len(conf.Host) > 0 {
		cmd = append(cmd, "--host", conf.Host)
	}
	if len(conf.Port) > 0 {
		cmd = append(cmd, "-p", conf.Port)
	}
	if len(conf.ReadBuffSize) > 0 {
		cmd = append(cmd, "-S", conf.ReadBuffSize)
	}
	if len(conf.RefreshToken) > 0 {
		cmd = append(cmd, "-r", conf.RefreshToken)
	}
	if len(conf.Root) > 0 {
		cmd = append(cmd, "--root", conf.Root)
	}
	if len(conf.WorkDir) > 0 {
		cmd = append(cmd, "-w", conf.WorkDir)
	}
	return cmd
}

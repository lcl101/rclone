package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/axgle/mahonia"
)

const WinAliyunDriveWebdav = "aliyundrive-webdav.exe"

//GetExecPath 获取当前路径
func CurrentDirectory() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	return string(path[0:i])
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func OsPathConverter(source string) string {
	return strings.Replace(strings.Replace(source, "\\", string(os.PathSeparator), -1), "/", string(os.PathSeparator), -1)
}

func WebdavPath() string {
	return CurrentDirectory() + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + WinAliyunDriveWebdav
}

func WebdavLogsPath() string {
	return CurrentDirectory() + string(os.PathSeparator) + "logs" + string(os.PathSeparator) + "webdav-executor-" + time.Now().Format("2006-01-02") + ".log"
}

func ConfigPath() string {
	return CurrentDirectory() + string(os.PathSeparator) + "myali.conf"
}

package utils

import (
	"os/user"
	"time"
)

var (
	HomeDir                = GetHomeDir()
	WindowsStartMenu       = HomeDir + `\AppData\Roaming\Microsoft\Windows\Start Menu\Programs`
	WindowsStartMenu2      = `C:\ProgramData\Microsoft\Windows\Start Menu\Programs`
	WindowsDeskTop         = HomeDir + `\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`
	WindowsDirs            = []string{WindowsStartMenu, WindowsStartMenu2, WindowsDeskTop}
	WindowsCredentials     = HomeDir + `\AppData\Local\Microsoft\Credentials`
	FileZillaProfilesDir   = HomeDir + `\AppData\Roaming\FileZilla\`
	FileZillaProfile1      = `\recentservers.xml`
	FileZillaProfile2      = `\sitemanager.xml`
	WinSCPProfilePath      = HomeDir + `\AppData\Roaming\`
	WinSCPProfile          = `\WinSCP.ini`
	TortoiseSVNProfilePath = HomeDir + `\AppData\Roaming\Subversion\auth\svn.simple`
)

const (
	PROCESS_VM_READ           = 0x0010
	PROCESS_QUERY_INFORMATION = 0x0400
)

type Proc struct {
	PID  int32
	Path string
	Name string
}

func GetHomeDir() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.HomeDir
}
func GetCurrentDateString(format string) string {
	return time.Now().Format(format)
}

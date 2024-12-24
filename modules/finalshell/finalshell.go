package finalshell

import (
	"PassGet/modules/utils"
	"encoding/json"
	"fmt"
	"github.com/parsiya/golnk"
	"log"
	"os"
	"strings"
)

func Get(Path string) (*ClientConfig, []ServerDetail) {
	if Path == "" {
		WorkDir := GetInstallPath()
		Path = WorkDir
	}
	ConnFileList, _, _ := utils.ListFilesAndDirs(Path + `\conn`)
	return GetClientConfig(Path), GetConnDetails(ConnFileList)
}

func GetInstallPath() string {
	WorkDir := ""
	for _, dir := range utils.WindowsDirs {
		files, _, err := utils.ListFilesAndDirs(dir)
		if err != nil {
			continue
		}
		for _, file := range files {
			if strings.Contains(file, "FinalShell.lnk") {
				f, _ := lnk.File(file)
				WorkDir = f.StringData.WorkingDir
				break
			}
		}
		if WorkDir != "" {
			break
		}
	}
	if WorkDir == "" {
		log.Fatal("Can not get install path")
	}
	return WorkDir
}
func GetClientConfig(Path string) *ClientConfig {
	ConfigFilepath := Path + `\config.json`
	data, err := os.ReadFile(ConfigFilepath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return nil
	}

	// 解析 JSON
	var C = new(ClientConfig)
	err = json.Unmarshal(data, C)
	if err != nil {
		log.Println("parse error", err)
		return nil
	}
	return C
}
func GetConnDetails(ConnFileList []string) []ServerDetail {
	ServerDetails := make([]ServerDetail, 0)
	for _, ConnFile := range ConnFileList {
		if !strings.Contains(ConnFile, "_connect_config.json") {
			continue
		}
		data, err := os.ReadFile(ConnFile)
		if err != nil {
			fmt.Println("读取文件失败:", err)
			return nil
		}
		var C = ServerDetail{}
		err = json.Unmarshal(data, &C)
		if err != nil {
			log.Println("parse error", err)
			return nil
		}
		if C.AuthenticationType == PASSWORD_AUTH {
			C.AuthType = "PASSWORD"
		} else if C.AuthenticationType == PUBLIC_KEY_AUTH {
			C.AuthType = "PUBLIC_KEY"
		}
		if C.ConectionType == SSH_AUTH {
			C.ConnType = "SSH"
		} else if C.ConectionType == RDP_AUTH {
			C.ConnType = "RDP"
		}
		if C.Password != "" {
			C.PasswordPlainText, _ = Decrypt(C.Password)
		}
		ServerDetails = append(ServerDetails, C)
	}
	return ServerDetails
}

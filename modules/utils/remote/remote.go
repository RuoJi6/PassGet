package remote

import (
	"PassGet/modules/utils"
	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/sys/windows"
	"path/filepath"
	"strconv"
	"strings"
	"unsafe"
)

func OpenProcess(pid int32) (windows.Handle, error) {
	handle, err := windows.OpenProcess(utils.PROCESS_VM_READ|utils.PROCESS_QUERY_INFORMATION, false, uint32(pid))
	if err != nil {
		//	return 0, fmt.Errorf("failed to open process: %v", err)
	}
	return handle, nil
}
func ReadMemory(handle windows.Handle, address uintptr, size uint32) ([]byte, error) {
	buffer := make([]byte, size)
	var bytesRead uintptr
	err := windows.ReadProcessMemory(handle, address, &buffer[0], uintptr(size), &bytesRead)
	if err != nil {
		//	return nil, fmt.Errorf("failed to read memory: %v", err)
	}
	return buffer, nil
}
func SearchMemory(handle windows.Handle, pattern []byte, todesk bool) ([]uintptr, []byte, error) {
	var results []uintptr
	var memoryInfo windows.MemoryBasicInformation
	var datas []byte
	if todesk {
		datas = make([]byte, 0)
	}
	address := uintptr(0)
	for {
		err := windows.VirtualQueryEx(handle, address, &memoryInfo, unsafe.Sizeof(memoryInfo))
		if err != nil || memoryInfo.RegionSize == 0 {
			break
		}

		if memoryInfo.State == windows.MEM_COMMIT && (memoryInfo.Protect&windows.PAGE_READWRITE) != 0 {
			data, err := ReadMemory(handle, memoryInfo.BaseAddress, uint32(memoryInfo.RegionSize))
			if err == nil {
				if todesk {
					datas = append(datas, data...)
				}
				for i := 0; i < len(data)-len(pattern); i++ {
					if MatchPattern(data[i:i+len(pattern)], pattern) {
						results = append(results, memoryInfo.BaseAddress+uintptr(i))
					}
				}
			}
		}
		address = memoryInfo.BaseAddress + uintptr(memoryInfo.RegionSize)
	}
	if todesk {
		return results, datas, nil
	}
	return results, nil, nil
}
func MatchPattern(data, pattern []byte) bool {
	for i := range pattern {
		if data[i] != pattern[i] {
			return false
		}
	}
	return true
}
func ExtractBetween(value, startDelim, endDelim string) string {
	start := strings.Index(value, startDelim)
	if start == -1 {
		return ""
	}
	start += len(startDelim)

	end := strings.Index(value[start:], endDelim)
	if end == -1 {
		return ""
	}

	return value[start : start+end]
}
func IsNumeric(s string) bool {
	_, err := strconv.Atoi(strings.TrimSpace(s))
	return err == nil
}
func GetProcessList(name string) []utils.Proc {
	processes, err := process.Processes()
	if err != nil {
		return nil
	}
	Procs := make([]utils.Proc, 0)
	for _, proc := range processes {
		procName, err := proc.Name()
		if err != nil {
			continue
		}

		// 如果进程名称匹配
		if strings.EqualFold(procName, name) {
			// 获取进程的可执行文件路径
			path, err := proc.Exe()
			if err != nil {
				continue
			}
			// 将 PID 和路径添加到结果中
			Procs = append(Procs, utils.Proc{PID: proc.Pid, Path: filepath.Dir(path), Name: procName})
		}
	}

	// 如果没有找到任何匹配的进程，返回 nil
	if len(Procs) == 0 {
		return nil
	}

	return Procs
}

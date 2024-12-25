package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func CheckIsInSlice(Items []string, Item string) bool {
	for _, item := range Items {
		if item == Item {
			return true
		}
	}
	return false
}
func GetSliceLastOne(Items []string) (Item string) {
	return Items[len(Items)-1]
}
func ListFilesAndDirs(directory string) ([]string, []string, error) {
	var files []string
	var dirs []string

	// 使用 filepath.Walk 遍历目录
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 忽略根目录本身
		if path == directory {
			return nil
		}

		// 判断是文件还是文件夹
		if info.IsDir() {
			dirs = append(dirs, path) // 如果是目录，添加到 dirs 列表
		} else {
			files = append(files, path) // 如果是文件，添加到 files 列表
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return files, dirs, nil
}
func CheckIsAdmin() bool {
	_, err := user.Current()
	return err == nil
}
func OutPutToFile(content string, dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 目录不存在，创建目录
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/results.txt", dir), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入内容
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

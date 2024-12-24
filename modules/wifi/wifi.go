package wifi

import (
	"fmt"
	"os/exec"
	"strings"
)

type Wifi struct {
	SSID     string
	Password string
}

// getWifiPasswords 获取已连接过的 Wi-Fi 网络名称及其密码
func getWifiPasswords() ([]Wifi, error) {
	cmd := exec.Command("netsh", "wlan", "show", "profiles")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run netsh command: %v", err)
	}
	profiles := make([]Wifi, 0)
	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "User Profile") {
			data := Wifi{
				SSID: strings.TrimSpace(strings.Split(line, ":")[1]),
			}
			password, err := getWifiPassword(data.SSID)
			if err != nil {
				continue
			}
			data.Password = password
			profiles = append(profiles, data)
		}
	}
	return profiles, nil
}

// getWifiPassword 获取指定网络的密码
func getWifiPassword(profileName string) (string, error) {
	// 执行 netsh wlan show profile "profileName" key=clear 命令来查看密码
	cmd := exec.Command("netsh", "wlan", "show", "profile", profileName, "key=clear")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run netsh command for profile %s: %v", profileName, err)
	}

	// 解析输出，查找密码字段
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Key Content") {
			// 提取密码
			password := strings.TrimSpace(strings.Split(line, ":")[1])
			return password, nil
		}
	}
	return "", fmt.Errorf("password not found for profile %s", profileName)
}

func Get() []Wifi {
	// 获取连接过的 Wi-Fi 网络及其密码
	profiles, err := getWifiPasswords()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return profiles
}

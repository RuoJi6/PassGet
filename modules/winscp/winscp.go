package winscp

import (
	"PassGet/modules/utils"
	"fmt"
	"github.com/go-ini/ini"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type ServerDetail struct {
	Name          string
	HostName      string
	PortNumber    string
	UserName      string
	Password      string
	PasswordPlain string
}

const PW_MAGIC = 0xA3
const PW_FLAG = 0xFF

type Flags struct {
	flag          rune
	remainingPass string
}

func (S *ServerDetail) DecryptNextCharacterWinSCP(Password string) Flags {
	var Flag Flags
	bases := "0123456789ABCDEF"

	// Find the first and second character values
	firstval := strings.Index(bases, string(Password[0])) * 16
	secondval := strings.Index(bases, string(Password[1]))
	added := firstval + secondval
	Flag.flag = rune(((^(added ^ PW_MAGIC) % 256) + 256) % 256) // decrypt the character
	Flag.remainingPass = Password[2:]                           // Remaining password
	return Flag
}
func Get(Path string) []ServerDetail {
	return GetSavedData(Path)
}
func GetSavedData(Path string) []ServerDetail {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Martin Prikryl\WinSCP 2\Sessions`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
	defer k.Close()
	if err != nil {
		fmt.Println("Can Not Found Data From Registry.")
		ServerDetails := GetSaveDataFromFile(Path)
		if ServerDetails != nil {
			return ServerDetails
		}
		fmt.Println("Can Not Found Data From File.")
		return nil
	}
	Sessions, err := k.ReadSubKeyNames(0)
	if err == nil && len(Sessions) > 1 {
		ServerDetails := make([]ServerDetail, 0)
		for _, Session := range Sessions {
			if strings.Contains(Session, "Default") {
				continue
			}
			Server := ServerDetail{
				Name: Session,
			}
			S, err := registry.OpenKey(registry.CURRENT_USER, `Software\Martin Prikryl\WinSCP 2\Sessions\`+Session, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
			defer k.Close()
			if err != nil {
				continue
			}
			Server.HostName, _, _ = S.GetStringValue(`HostName`)
			Server.PortNumber, _, _ = S.GetStringValue(`PortNumber`)
			Server.UserName, _, _ = S.GetStringValue(`UserName`)
			Server.Password, _, _ = S.GetStringValue(`Password`)
			Server.Decrypt()
			ServerDetails = append(ServerDetails, Server)
		}
		return ServerDetails
	} else {
		fmt.Println("Can Not Found Data From Registry.")
		return nil
	}
}
func GetSaveDataFromFile(Path string) []ServerDetail {
	if Path == "" {
		Path = utils.WinSCPProfilePath
	}
	file, _, err := utils.ListFilesAndDirs(Path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if utils.CheckIsInSlice(file, utils.WinSCPProfile) {
		fmt.Println("No Data File.")
		return nil
	}
	Sections := getSection(Path)
	if len(Sections) == 0 {
		fmt.Println("No Data File.")
		return nil
	}
	cfg, err := ini.Load(Path + utils.WinSCPProfile)
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		return nil
	}
	ServerDetails := make([]ServerDetail, 0)
	for _, v := range Sections {
		Server := ServerDetail{
			Name: strings.Split(v, `\`)[1],
		}
		Server.HostName = cfg.Section(v).Key("HostName").String()
		Server.PortNumber = cfg.Section(v).Key("PortNumber").String()
		Server.UserName = cfg.Section(v).Key("UserName").String()
		Server.Password = cfg.Section(v).Key("Password").String()
		Server.Decrypt()
		ServerDetails = append(ServerDetails, Server)
	}
	return ServerDetails
}
func getSection(Path string) []string {
	data, err := ioutil.ReadFile(Path + utils.WinSCPProfile) // 假设配置文件为 config.ini
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// 正则表达式匹配所有 [Session.*] 部分
	re := regexp.MustCompile(`\[Session.*?\]`)

	// 查找所有匹配的部分
	matches := re.FindAllString(string(data), -1)
	final_datas := make([]string, 0)
	// 输出匹配结果
	for _, match := range matches {
		final_datas = append(final_datas, match[1:len(match)-1])
	}
	return final_datas
}
func (S *ServerDetail) Decrypt() {
	if S.Password == "" {
		S.PasswordPlain = ""
		return
	}
	var length rune
	var clearpwd string
	unicodeKey := S.UserName + S.HostName
	Flag := S.DecryptNextCharacterWinSCP(S.Password)
	storedFlag := Flag.flag

	if storedFlag == PW_FLAG {
		Flag = S.DecryptNextCharacterWinSCP(Flag.remainingPass)
		Flag = S.DecryptNextCharacterWinSCP(Flag.remainingPass)
		length = Flag.flag
	} else {
		length = Flag.flag
	}

	Flag = S.DecryptNextCharacterWinSCP(Flag.remainingPass)
	Flag.remainingPass = Flag.remainingPass[Flag.flag*2:]

	// Loop to get the password
	for i := 0; i < int(length); i++ {
		Flag = S.DecryptNextCharacterWinSCP(Flag.remainingPass)
		clearpwd += string(Flag.flag)
	}

	// If the storedFlag is PW_FLAG, check for the unicodeKey in the clear password
	if storedFlag == PW_FLAG {
		if strings.HasPrefix(clearpwd, unicodeKey) {
			clearpwd = clearpwd[len(unicodeKey):] // Remove unicodeKey prefix
		} else {
			clearpwd = "" // Invalid password
		}
	}
	S.PasswordPlain = clearpwd
}

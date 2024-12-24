package navicat

import (
	"PassGet/modules/utils"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"strconv"
	"strings"
)

var (
	base = `SOFTWARE\PremiumSoft`
)

func Check() (SubKeys []string, Exists bool, err error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, base, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
	if err != nil {
		return nil, false, err
	}
	defer k.Close()

	// 读取字符串值
	Keys, err := k.ReadSubKeyNames(0)
	if err != nil {
		return nil, false, err
	}
	SubKeys = GetTypeKeys(Keys)
	return SubKeys, true, nil
}
func Get() ([]ServerDetail, error) {
	SubKeys, exist, err := Check()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !exist {
		log.Println("Navicat Client Not Found")
		return nil, nil
	}
	ServerKeys, _ := GetServerKeys(SubKeys)
	return GetDetails(ServerKeys)
}
func GetTypeKeys(Keys []string) []string {
	TypeKeys := make([]string, 0)
	for _, Key := range Keys {
		k, err := registry.OpenKey(registry.CURRENT_USER, base+`\`+Key, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
		if err != nil {
			continue
		}
		defer k.Close()
		data, err := k.ReadSubKeyNames(0)
		if err != nil {
			continue
		}
		if utils.CheckIsInSlice(data, `Servers`) {
			TypeKeys = append(TypeKeys, fmt.Sprintf(`%s\%s`, base, Key))
		}
	}
	return TypeKeys
}
func GetServerKeys(SubKeys []string) (ServerKeys map[string][]string, err error) {
	ServerKeys = make(map[string][]string, 0)
	for _, Key := range SubKeys {
		k, err := registry.OpenKey(registry.CURRENT_USER, Key+`\Servers`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
		defer k.Close()
		if err != nil {
			continue
		}
		Type := utils.GetSliceLastOne(strings.Split(Key, `\`))
		Type = strings.Replace(Type, "Navicat", "", -1)
		if Type == "" {
			Type = "MYSQL"
		}
		Keys, err := k.ReadSubKeyNames(0)
		for _, key := range Keys {
			ServerKeys[Type] = append(ServerKeys[Type], fmt.Sprintf(`%s\Servers\%s`, Key, key))
		}
	}
	return ServerKeys, nil
}
func GetDetails(ServerKeys map[string][]string) (Details []ServerDetail, err error) {
	Details = make([]ServerDetail, 0)
	for Type, ServerKey := range ServerKeys {
		for _, Key := range ServerKey {
			k, err := registry.OpenKey(registry.CURRENT_USER, Key, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
			defer k.Close()
			if err != nil {
				log.Println(err)
				continue
			}
			Server := ServerDetail{
				Type: Type,
			}
			Server.Password, _, _ = k.GetStringValue(`PWD`)
			if Server.Password != "" {
				Server.PasswordPlainText, _ = decrypt(Server.Password)
			}
			if Server.Type == "SQLITE" {
				Server.DataBaseFileName, _, _ = k.GetStringValue(`DatabaseFileName`)
				Details = append(Details, Server)
				continue
			}
			Server.Host, _, _ = k.GetStringValue(`Host`)
			Port, _, _ := k.GetIntegerValue(`Port`)
			Server.Port = strconv.Itoa(int(Port))
			Server.UserName, _, _ = k.GetStringValue(`UserName`)
			if Server.Type == "MYSQL" {
				Details = append(Details, Server)
				continue
			}
			if Server.Type == "PG" || Server.Type == "MSSQL" {
				Server.InitialDatabase, _, _ = k.GetStringValue(`InitialDatabase`)
				Details = append(Details, Server)
				continue
			}
			if Server.Type == "MONGODB" {
				Server.AuthSource, _, _ = k.GetStringValue(`AuthSource`)
				Details = append(Details, Server)
				continue
			}
			if Server.Type == "ORACLE" {
				Server.OraServiceNameType, _, _ = k.GetStringValue(`OraServiceNameType`)
				Server.InitialDatabase, _, _ = k.GetStringValue(`InitialDatabase`)
				Details = append(Details, Server)
				continue
			}
		}
	}
	return Details, nil
}

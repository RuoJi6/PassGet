package navicat

type ServerDetail struct {
	Type               string `json:"type"`
	Host               string `json:"host"`
	Port               string `json:"port"`
	UserName           string `json:"userName"`
	Password           string `json:"password"`
	PasswordPlainText  string `json:"passwordPlainText"`
	OraServiceNameType string `json:"oraServiceNameType"`
	InitialDatabase    string `json:"initialDatabase"`
	DataBaseFileName   string `json:"dataBaseFileName"`
	AuthSource         string `json:"authSource"`
}

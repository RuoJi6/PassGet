package finalshell

var (
	PASSWORD_AUTH   = 1
	PUBLIC_KEY_AUTH = 2
	SSH_AUTH        = 100
	RDP_AUTH        = 101
)

type ServerDetail struct {
	UserName           string `json:"user_name"`
	ConectionType      int    `json:"conection_type"`
	ConnType           string
	AuthType           string
	Description        string `json:"description"`
	AuthenticationType int    `json:"authentication_type"`
	Password           string `json:"password"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	Name               string `json:"name"`
	SecretKeyId        string `json:"secret_key_id"`
	PasswordPlainText  string
}
type ClientConfig struct {
	SecretKeyList []struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Id       string `json:"id"`
		KeyData  string `json:"key_data"`
	} `json:"secret_key_list"`
}

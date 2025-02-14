package configs

import (
	"encoding/json"
	"os"
)

type Configs struct {
	FileLc         string `json:"file_lc"`
	DataSource     string `json:"data_source"`
	Port           string `json:"port"`
	AccessSecret   string `json:"access_secret,omitempty"`
	ExpireAccess   string `json:"expire_access,omitempty"`
	RefreshSecret  string `json:"refresh_secret,omitempty"`
	ExpireRefresh  string `json:"expire_refresh,omitempty"`
	KeyAES128      string `json:"keyAES128"`
	FromEmail      string `json:"fromEmail"`
	PasswordEmail  string `json:"passwordEmail"`
	SmtpHost       string `json:"smtpHost,omitempty"`
	SmtpPort       string `json:"smtpPort,omitempty"`
	AddressRedis   string `json:"addressRedis"`
	PasswordRedis  string `json:"passwordRedis"`
	DatabasedIndex string `json:"databased_index"`
	ConfigPathFile string `json:"configPathFile,omitempty"`
	HostAccess     string `json:"hostAccess,omitempty"`
}

var config *Configs

func Get() *Configs {
	return config
}
func LoadConfig(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	byteValue, err := os.ReadFile(configFile.Name())
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
}

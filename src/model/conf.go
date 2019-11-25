package model

import "encoding/json"

type Configuration struct {
	Address                string                 `yaml:"address"`
	Database               Database               `yaml:"database"`
	EmailServer            EmailServer            `yaml:"emailServer"`
	DingTalkAuthentication DingTalkAuthentication `yaml:"dingtalk"`
	Env                    string                 `yaml:"env"`
}

type EmailServer struct {
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	SMTPServer string `yaml:"smtpServer"`
	Port       int    `yaml:"port"`
	SSL        bool   `yaml:"ssl"`
}

type Database struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
	Log          bool   `yaml:"log"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

type DingTalkAuthentication struct {
	AgentID           int64  `yaml:"agentID"`
	CorporationID     string `yaml:"corpID"`
	CorporationSecret string `yaml:"corpSecret"`
}

func (c *Configuration) String() string {
	if dataInBytes, err := json.Marshal(&c); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func (es *EmailServer) String() string {
	if dataInBytes, err := json.Marshal(&es); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func (d *Database) String() string {
	if dataInBytes, err := json.Marshal(&d); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func (dta *DingTalkAuthentication) String() string {
	if dataInBytes, err := json.Marshal(&dta); err == nil {
		return string(dataInBytes)
	}

	return ""
}

package models

type YamlStruct struct {
	WeChat struct{
		AgentID string `yaml:"agentid"`
		Color string `yaml:"color"`
		CorpID string `yaml:"corpid"`
		Corpsecret string `yaml:"corpsecret"`
		ToUser	string `yaml:"touser"`
	}
	DingDing struct{
		AgentID string `yaml:"agentid"`
		Color string `yaml:"color"`
		CorpID string `yaml:"corpid"`
		Corpsecret string `yaml:"corpsecret"`
	}
	Email struct{
		ToEmail string `yaml:"toemail"`
		FromEmail string `yaml:"fromemail"`
		EmailSubject string `yaml:"emailsubject"`
	}
	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DbName   string `yaml:"dbname"`
	}
	PtQueryDigest struct {
		Enable   bool   `yaml:"enable"`
		Ptcommand string `yaml:"ptcommand"`
		LogPath  string `yaml:"logpath"`
		SlowName string `yaml:"slowname"`
		Values struct {
			OrderBy string `yaml:"orderby"`
			Limit   string `yaml:"limit"`
			Filter  string `yaml:"filter"`
		}
	}
}

/*
type StringToJson struct {
	SlowLogValues string `json:"slowlogvalues"`
}*/

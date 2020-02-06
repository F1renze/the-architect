package config

type ServiceConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type WebConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
}

func GetWebConfig(cfgName string, cli CMSClient) (*WebConfig, error) {
	cfg := new(WebConfig)
	err := cli.Scan(cfgName, cfg)
	return cfg, err
}

func GetSrvConfig(cfgName string, cli CMSClient) (*ServiceConfig, error) {
	cfg := new(ServiceConfig)
	err := cli.Scan(cfgName, cfg)
	return cfg, err
}

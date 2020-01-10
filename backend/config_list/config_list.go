package config_list

import "backend/restapi/operations"

type Config func(api *operations.BroodClashAPI)

type ConfigList []Config

func (cl *ConfigList) Register(c Config) {
	*cl = append(*cl, c)
}

func (cl *ConfigList) Init(api *operations.BroodClashAPI) {
	for _, c := range *cl {
		c(api)
	}
}

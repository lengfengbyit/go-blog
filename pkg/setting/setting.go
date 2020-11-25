package setting

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			if _, err := os.Stat(config); os.IsNotExist(err) {
				return nil, errors.New(fmt.Sprintf("config path: %s is not exist", config))
			}

			vp.AddConfigPath(config)
		}
	}
	
	//vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}



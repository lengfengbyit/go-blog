package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
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

	s := &Setting{vp}
	s.WatchSettingChange()

	return s, nil
}

// WatchSettingChange 监控配置文件是否修改
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
			fmt.Println("config update:", in.Name, in.Op, in.String())
		})
	}()
}

package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// NewSetting 是新增全局设置
// 参数configs是配置文件所在路径
func NewSetting(configs ...string) (*Setting, error) {
	// 读取配置文件
	vp := viper.New()
	// vp.SetConfigFile("configs/config.yaml")
	// vp.AddConfigPath("configs/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// 配置文件热更新
	st := &Setting{vp}
	st.WatchSettingChange()
	return st, nil
}

// WatchSettingChange 可以使配置文件热更新
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}


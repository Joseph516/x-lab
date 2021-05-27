package settings

import "time"

var sections = make(map[string]interface{})

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int // gorm中传入int，而不是time.Second
	MaxOpenConns int
}

type AppSetting struct {
	UploadServerUrl      string
	UploadSavePath       string
	UploadImageAllowExts []string
	UploadImageMaxSize   int
}

// ReadSection 将配置文件中指定键k对应的配置项，保存至值v中
// k为键，对应config文件夹中的每一类
// v对应global.setting中的变量
func (s *Setting) ReadSection(k string, v interface{}) error {
	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	// 将vp中对应key的数据，解析到v对应的结构体中
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	// 配置热更新，将配置文件中的所有配置项保存在sections中
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

// ReloadAllSection 依次重新载入所有的sections配置项
func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

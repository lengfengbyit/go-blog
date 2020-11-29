package setting

import "time"

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	LogDateSuffix   bool
	ContextTimeout  time.Duration
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	Port         int
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type UploadSetting struct {
	SavePath       string   // 文件的保存路径
	ServerUrl      string   // 用户展示文件的服务地址
	ImageMaxSize   int      // 单位 M
	ImageAllowExts []string // 允许的图片后缀
}

type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSetting struct {
	Enable   bool
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

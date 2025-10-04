package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// setting 实例化工厂方法
// 初始化 viper 对象并返回
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	// 指定配置文件的类型
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

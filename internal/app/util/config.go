package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver        string `mapstructure:"DB_DRIVER"`
	DBSource        string `mapstructure:"DB_SOURCE"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	PageAccessToken string `mapstructure:"PAGE_ACCESS_TOKEN"`
	VerifyToken     string `mapstructure:"VERIFY_TOKEN"`
	FacebookApi     string `mapstructure:"FACEBOOK_API"`
}

const (
	DefaultScript = "HELP"
	UnknownScript = "UNKNOWN"

	DefaultSubject = "VAT_LY"

	MessageText     = "TEXT"
	MessagePostback = "POSTBACK"

	ProcessGuideAction    = "ProcessGuideAction"
	ProcessPracticeAction = "ProcessPracticeAction"
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

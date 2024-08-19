package formSpree

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   `yaml:"server"`
	JWT      `yaml:"jwt"`
	Database `yaml:"dataBase"`
	Mail     MailCreds `yaml:"mail"` // Mail configuration settings.
}

type Server struct {
	Port           int      `yaml:"port"`
	ExposedHeaders []string `yaml:"exposedHeaders"`
	AllowedHeaders []string `yaml:"allowedHeaders"`
	AllowedMethods []string `yaml:"allowedMethods"`
	AllowedOrigins []string `yaml:"allowedOrigins"`
}

type JWT struct {
	Key string `yaml:"key"`
}

type Database struct {
	DB string `yaml:"db"`
}

type MailCreds struct {
	TemplatePath string `yaml:"templatePath"` // Template represents the email template.
	User         string `yaml:"user"`         // User represents the email user.
	Secret       string `yaml:"secret"`       // Secret represents the email password or secret.
	Host         string `yaml:"host"`         // Host represents the email server host.
	Port         int    `yaml:"port"`         // Port represents the email server port.
}

func GlobalConfig() *Config {
	config := Config{}
	path := "/home/ubuntu/formSpree/"
	path="./"
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		// If there is an error reading the configuration file, panic with the error.
		log.Panic(err)
	}
	// Unmarshal the configuration file into the Config struct.
	if err := viper.Unmarshal(&config); err != nil {
		log.Panic(err)
	}
	// Return a pointer to the loaded Config struct.
	return &config
}

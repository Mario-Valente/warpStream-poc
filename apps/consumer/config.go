package main

import "github.com/spf13/viper"

type Config struct {
	TopicName string `yaml:"topic_name"`
	BrokerURL string `yaml:"broker_url"`
}

func LoadConfig() *Config {
	viper.AutomaticEnv()

	cfg := &Config{
		TopicName: viper.GetString("TOPIC_NAME"),
		BrokerURL: viper.GetString("BROKER_URL"),
	}

	return cfg
}

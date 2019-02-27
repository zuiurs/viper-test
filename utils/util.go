package utils

import (
	"github.com/spf13/viper"
)

func ShowUser() string {
	return viper.GetString("global.vp-user")
}

func ShowPassword() string {
	return viper.GetString("global.vp-password")
}

package utils

import (
	"testing"

	"github.com/spf13/viper"
)

const (
	TestUser     = "test"
	TestPassword = "abc"
)

func init() {
	viper.Set("global.vp-user", TestUser)
	viper.Set("global.vp-password", TestPassword)
}

func TestShowUser(t *testing.T) {
	if u := ShowUser(); u != TestUser {
		t.Errorf("User is not match: %#v", u)
	}
}

func TestShowPassword(t *testing.T) {
	if p := ShowPassword(); p != TestPassword {
		t.Errorf("Password is not match: %#v", p)
	}
}

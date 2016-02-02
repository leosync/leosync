package lingualeo

import (
	"errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
)

type config struct {
	Email    string
	Password string
}

func createConfig() (conf *config) {
	var errs []error = nil

	email := viper.Get("lingualeo.login")
	if email == nil {
		errs = append(errs, errors.New("LinguaLeo email not defined"))
	}
	password := viper.Get("lingualeo.password")
	if password == nil {
		errs = append(errs, errors.New("LinguaLeo password not defined"))
	}
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	conf = &config{
		Email:    cast.ToString(email),
		Password: cast.ToString(password),
	}
	return conf
}

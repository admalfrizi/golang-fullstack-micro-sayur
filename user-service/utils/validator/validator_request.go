package validator

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type Validator struct {
	Validator *validator.Validate
	Translator ut.Translator
}

func NewValidator() *Validator {
	en := en.New()
	uni := ut.New(en,en)
	trans, found := uni.GetTranslator("en")

	if !found {
		log.Fatalf("[NewValidator-1] Translator not found")
	}

	validate := validator.New()

	return &Validator{
		Validator: validate,
		Translator: trans,
	}
}

func (v *Validator) Validate(i interface{}) error {
	err := v.Validator.Struct(i)

	if err != nil {
		object, _ := err.(validator.ValidationErrors)

		for _, e := range object {
			log.Infof("[Validate-1] %s: %s", e.Field(), e.Translate(v.Translator))

			return errors.New(e.Translate(v.Translator))
		}
	}

	return nil
}
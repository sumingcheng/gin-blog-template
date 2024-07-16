package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var trans ut.Translator

func init() {
	if err := transInit("zh"); err != nil {
		LogRus.Fatal("Failed to initialize translator:", err)
	}
}

func transInit(locale string) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()              // Chinese translator
		enT := en.New()              // English translator
		uni := ut.New(enT, zhT, enT) // Universal translator

		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// Register translation based on locale
		switch locale {
		case "zh":
			return chTranslations.RegisterDefaultTranslations(v, trans)
		case "en":
			return enTranslations.RegisterDefaultTranslations(v, trans)
		default:
			return enTranslations.RegisterDefaultTranslations(v, trans)
		}
	}

	return fmt.Errorf("failed to assert Validator")
}

func TranslateErrors(err error) string {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var errMessages []string
		for _, e := range errs {
			translatedMsg := e.Translate(trans)
			errMessages = append(errMessages, translatedMsg)
		}
		return strings.Join(errMessages, ", ")
	}
	return err.Error()
}

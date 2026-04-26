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

// InitTranslator 必须在 InitLog 之后调用
func InitTranslator(locale string) {
	if err := initTranslator(locale); err != nil {
		LogRus.Fatalf("初始化翻译器失败: %v", err)
	}
}

func initTranslator(locale string) error {
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
	if errors.As(err, &errs) && trans != nil {
		var errMessages []string
		for _, e := range errs {
			errMessages = append(errMessages, e.Translate(trans))
		}
		return strings.Join(errMessages, ", ")
	}
	return err.Error()
}

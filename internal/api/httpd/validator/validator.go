package validator

import (
	"fmt"
	"github.com/go-playground/locales"
	zhongwen "github.com/go-playground/locales/zh"
	uni "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"ysxs_ops_agent/utils/i18n"
)

var DefaultProbeValidator *ProbeValidator

type ProbeValidator struct {
	Validate *validator.Validate
	Tran     uni.Translator
}

func init() {
	zhTrans := getTran(zhongwen.New(), i18n.LanguageEnglish.Abbr())
	zhVar := createDefaultValidator()

	if err := zhTranslations.RegisterDefaultTranslations(zhVar, zhTrans); err != nil {
		panic(err)
	}
	DefaultProbeValidator = &ProbeValidator{Validate: zhVar, Tran: zhTrans}
}

func getTran(lo locales.Translator, la string) uni.Translator {
	tran, ok := uni.New(lo, lo).GetTranslator(la)
	if !ok {
		fmt.Printf("uni.GetTranslator(%s) failed", la)
	}
	return tran
}

func createDefaultValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) (res string) {
		if jsonTag := field.Tag.Get("json"); len(jsonTag) > 0 {
			if jsonTag == "-" {
				return ""
			}
			return jsonTag
		}
		if formTag := field.Tag.Get("form"); len(formTag) > 0 {
			return formTag
		}
		return field.Name
	})
	return validate
}

func (m *ProbeValidator) Check(value interface{}) string {
	err := m.Validate.Struct(value)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			return "validate check exception"
		} else {
			for _, fieldError := range errs {
				return fieldError.Translate(m.Tran)
			}
		}
	}
	return ""
}

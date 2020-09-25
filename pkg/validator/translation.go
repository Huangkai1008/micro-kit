package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"

	"github.com/Huangkai1008/micro-kit/pkg/message"
)

func registerTranslation(v *validator.Validate, locale string) (trans ut.Translator, err error) {
	zhTrans, enTrans := zh.New(), en.New()
	uniTrans := ut.New(enTrans, zhTrans, enTrans)
	trans, ok := uniTrans.GetTranslator(locale)
	if !ok {
		err = fmt.Errorf("uniTrans.GetTranslator(%s) failed", locale)
		return nil, errors.Wrap(err, message.TransGetTranslatorError)
	}
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return
}

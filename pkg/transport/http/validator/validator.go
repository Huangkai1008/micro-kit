package validator

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	kiterr "github.com/Huangkai1008/micro-kit/pkg/error"
)

type CustomValidator struct {
	validate *validator.Validate
	trans    ut.Translator
	locale   string
}

func New(locale string) (*CustomValidator, error) {
	validate := validator.New()
	trans, err := registerTranslation(validate, locale)
	if err != nil {
		return nil, err
	}
	return &CustomValidator{
		validate: validate,
		locale:   locale,
		trans:    trans,
	}, nil
}

func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.validate.Struct(i); err != nil {
		var builder strings.Builder
		errs := err.(validator.ValidationErrors)
		for idx, err := range errs {
			msg := err.Translate(v.trans)
			builder.WriteString(msg)
			if idx != len(errs)-1 {
				builder.WriteString(", ")
			}
		}
		return kiterr.NewValidationError(builder.String())
	}
	return nil
}

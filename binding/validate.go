package binding

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"github.com/go-playground/validator/v10"
)

// use a single instance , it caches struct info
var (
	uni *ut.UniversalTranslator
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	trans    ut.Translator
}

var _ binding.StructValidator = (*DefaultValidator)(nil)

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			errs := err.(validator.ValidationErrors)
			terr := errs.Translate(v.trans)
			for _, v := range terr {
				return errors.New(v)
			}
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		local := zh.New()
		uni = ut.New(local, local)
		v.trans, _ = uni.GetTranslator("zh")
		_ = zhTranslations.RegisterDefaultTranslations(v.validate, v.trans)
		v.translateOverride()
	})
}

func (v *DefaultValidator) translateOverride() {
	err := v.validate.RegisterTranslation("required", v.trans, func(ut ut.Translator) error {
		return ut.Add("required", "缺少{0}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	if err != nil {
		log.Println("defaultValidator translateOverride error:", err)
	}

	v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

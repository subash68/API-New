package controller

import (
	"reflect"
	"regexp"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

const (
	alphaNumericRegexSpaceString = "^[a-zA-Z0-99\\s]+$"
)

var (
	alphaNumericSpaceRegex = regexp.MustCompile(alphaNumericRegexSpaceString)
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}
	return nil
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		config := &validator.Config{TagName: "binding"}
		v.validate = validator.New(config)
		v.validate.RegisterValidation("alphanumspace", IsAlphanumSpace)
		v.validate.RegisterValidation("validateusername", ValidateUsrName)
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

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

// IsAlphanumSpace ...
func IsAlphanumSpace(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return alphaNumericSpaceRegex.MatchString(field.String())
}

// ValidateUsrName ...
func ValidateUsrName(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if name, ok := field.Interface().(string); ok {

		if valid, _ := regexp.MatchString("[a-zA-Z][a-zA-Z ]+[a-zA-Z]$", name); !valid {
			return false
		}

	}
	return true
}

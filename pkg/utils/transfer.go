package utils

import (
	"reflect"
)

type TransformInterface interface {
	Transform() map[string]interface{}
}

func Transform(data interface{}) interface{} {
	reflectType := reflect.TypeOf(data)
	reflectKind := reflectType.Kind()
	reflectValue := reflect.ValueOf(data)

	if reflectKind == reflect.Ptr {
		reflectType = reflectType.Elem()
		reflectKind = reflectType.Kind()
		reflectValue = reflectValue.Elem()
	}

	if reflectKind == reflect.Slice {
		var resSlice []map[string]interface{}

		for i := 0; i < reflectValue.Len(); i++ {
			obj, ok := reflectValue.Index(i).Interface().(TransformInterface)
			if ok {
				resSlice = append(resSlice, obj.Transform())
			}
		}

		return resSlice
	} else if reflectKind == reflect.Struct {
		var res map[string]interface{}

		obj, ok := reflectValue.Interface().(TransformInterface)
		if ok {
			res = obj.Transform()
		}

		return res
	}

	return nil
}

type CallBack func(data map[string]interface{}) map[string]interface{}

// data 类型 *[]map[string]interface{}
func MapTransform(data interface{}, fn CallBack) interface{} {
	_data := data.(*[]map[string]interface{})

	var res []map[string]interface{}

	for _, v := range *_data {
		res = append(res, fn(v))
	}

	return res
}

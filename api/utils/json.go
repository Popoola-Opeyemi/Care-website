package utils

import (
	"encoding/json"
	"reflect"

	"github.com/labstack/gommon/log"
)

// IsPointer ...
func IsPointer(val interface{}) bool {
	typ := reflect.TypeOf(val)
	return typ.Kind().String() == "ptr"
}

// SaveImageInJson ...
func SaveImageInJson(env *Enviroment, node interface{}) error {

	if IsDict(node) {
		dict := node.(map[string]interface{})
		for k := range dict {
			if IsList(dict[k]) {
				err := SaveImageInJson(env, dict[k])
				if err != nil {
					return err
				}
			} else if IsDict(dict[k]) {
				log.Info("k", k)
				if k == "image" {
					img := dict[k].(map[string]interface{})
					imgJSON, err := json.Marshal(dict[k])
					if err != nil {
						return err
					}

					asset, err := SaveImageData(env, imgJSON)
					if err != nil {
						return err
					}

					img["data"] = asset.Data
				} else {
					err := SaveImageInJson(env, dict[k])
					if err != nil {
						return err
					}
				}
			}
		}
	} else if IsList(node) {
		list := node.([]interface{})
		for i := 0; i < len(list); i++ {
			err := SaveImageInJson(env, list[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func IsDict(v interface{}) bool {
	_, ok := v.(map[string]interface{})
	if !ok {
		_, ok = v.(*map[string]interface{})
	}
	return ok
}

func IsList(v interface{}) bool {
	_, ok := v.([]interface{})
	return ok
}

package renderer

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/yigithakverdi/concord/internal/utils"
	"github.com/yigithakverdi/concord/internal/utils/globals"
)

func CreatePropertyAsString(key string, value string) string {
	return key + "=" + value + "\n"
}

func CreateNestedPropertiesAsString(parent []string, data reflect.Value) {
	switch data.Kind() {
	case reflect.Slice:
		for i := 0; i < data.Len(); i++ {
			value := data.Index(i)
			fmt.Printf("Data %v \n", i)
			CreateNestedPropertiesAsString(parent, reflect.ValueOf(value.Interface()))
		}
	case reflect.Map:
		for _, key := range data.MapKeys() {
			value := data.MapIndex(key)
			parent = append(parent, strings.ToLower(key.String()))
			CreateNestedPropertiesAsString(parent, reflect.ValueOf(value.Interface()))
			parent = parent[:len(parent)-1]
		}
	case reflect.Struct:
		t := data.Type()
		for i := 0; i < data.NumField(); i++ {
			field := data.Field(i)
			fieldName := t.Field(i).Name
			parent = append(parent, strings.ToLower(fieldName))
			CreateNestedPropertiesAsString(parent, field)
			parent = parent[:len(parent)-1]
		}
	default:
		globals.GlobalProperties[strings.TrimSpace(strings.Join(parent, "."))] = data.String()
	}
}

func RenderApplicationProperties() {
	f, err := os.OpenFile(globals.DefaultEnvPropertiesFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	utils.Check(err)
	for key, value := range globals.GlobalProperties {
		property := CreatePropertyAsString(key, value)
		_, err = f.WriteString(property)
		utils.Check(err)
	}
	defer f.Close()
}

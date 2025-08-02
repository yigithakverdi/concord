package main

import (

	// "github.com/yigithakverdi/concord/internal/utils/globals"

	"io/ioutil"
	"reflect"

	"github.com/yigithakverdi/concord/internal/utils/globals"
	"github.com/yigithakverdi/concord/pkg/parser"
	"github.com/yigithakverdi/concord/pkg/renderer"
	"gopkg.in/yaml.v3"
)

func main() {
	parser.ParseApplicationProperties()

	// defaultValues, _ := renderer.GetDefaultValues()
	// fmt.Println("Default Server Value: ", defaultValues.Server.Port)

	// for key, value := range globals.GlobalProperties {
	// 	fmt.Println("Parsed Property:", key, "=", value)
	// }

	content, err := ioutil.ReadFile(globals.DefaultValuesFileLocation)
	if err != nil {
		panic(err)
	}

	var data renderer.Data
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		panic(err)
	}

	renderer.CreateNestedPropertiesAsString([]string{}, reflect.ValueOf(data.Config))
	renderer.RenderApplicationProperties()
}

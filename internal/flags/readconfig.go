package flags

import (
	"fmt"
	"reflect"
	"strings"
)

func readConfig(path string) error {
	configValues, err := getUserConfig(path)
	if err != nil {
		return err
	}

	userConfig := reflect.Indirect(reflect.ValueOf(configValues))

	for i := range userConfig.NumField() {
		fmt.Println(userConfig.Field(i).Type().Name())

		longest := 0
		for x := range userConfig.Type().Field(i).Type.NumField() {
			if length := len(userConfig.Type().Field(i).Type.Field(x).Name); length > longest {
				longest = length
			}
		}
		for x := range userConfig.Type().Field(i).Type.NumField() {
			diff := longest - len(userConfig.Type().Field(i).Type.Field(x).Name)
			fmt.Println(userConfig.Type().Field(i).Type.Field(x).Name+strings.Repeat(" ", diff)+":", userConfig.Field(i).Field(x))
		}
		fmt.Println("\n")
	}

	return nil
}

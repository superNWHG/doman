package flags

import (
	"fmt"
	"reflect"

	"github.com/superNWHG/doman/internal/config"
)

func readconfig(path string) error {
	configValues, err := config.ReadConfig(path)
	if err != nil {
		return err
	}

	userConfig := reflect.Indirect(reflect.ValueOf(configValues))

	for i := range userConfig.NumField() {
		fmt.Println(userConfig.Type().Field(i).Name)
		for x := 0; x < userConfig.Type().Field(i).Type.NumField(); x++ {
			fmt.Println(userConfig.Type().Field(i).Type.Field(x).Name, ":", userConfig.Field(i).Field(x))
		}
		fmt.Println("\n")
	}

	return nil
}

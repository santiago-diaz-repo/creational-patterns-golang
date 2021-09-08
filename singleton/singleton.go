package singleton

import (
	"fmt"
)

type instance struct {
}

var singleInstance *instance

func GetInstance() *instance {
	if singleInstance == nil {
		fmt.Println("Creating instance")
		singleInstance = &instance{}
	}

	return singleInstance
}

package lib

import (
	"fmt"
	"log"
	"reflect"
)

func PrintType(v interface{}) {
	vValue := reflect.ValueOf(v)
	vType := vValue.Type()
	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)
		fieldValue := vValue.Field(i)
		fmt.Printf("%s: %v\n", field.Name, fieldValue.Interface())
	}
}

func TestLog() {
	log.Println("======================================================================================================")
}

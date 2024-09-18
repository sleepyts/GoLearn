package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func showAll(any ...interface{})   {
	for _, v := range any {
		t, val := reflect.TypeOf(v).Elem(), reflect.ValueOf(v)
		fmt.Printf("Type: %s, Value: %v\n", t, val.Elem())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("Field: %s, Type: %s, Value: %v Tag: %v\n", field.Name, field.Type, val.Elem().Field(i),field.Tag)
		}
	}
}
func InterfaceTest(){
	comment:= Comment{
		Id: 1,
		Content: "This is a comment",
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
		ReplyOnId: -1,
	}
	arr,_:=json.Marshal(comment)
	fmt.Printf("%s",arr)
	json.Unmarshal(arr,&comment)
	fmt.Println(comment)
}
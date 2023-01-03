package repository

import (
	"log"
	"reflect"
	"testing"
)

func TestNilStringCompare(t *testing.T) {
	var nilStr string = ""
	str := "not nil"
	log.Println(&nilStr == nil, &str == nil, reflect.ValueOf(nilStr).IsZero())
}

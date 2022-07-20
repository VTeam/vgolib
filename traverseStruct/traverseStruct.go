package traverseStruct

import (
	"fmt"
	"reflect"
	"strings"
)

func splitTag(s string, tagType string) string {
	if s == "" {
		panic("empty s")
	}

	tags := strings.Split(s, " ")
	if len(tags) < 1 {
		return ""
	}

	for i := 0; i < len(tags); i++ {
		tagSubStr := strings.Split(tags[i], ":")
		if tagSubStr[0] == tagType {
			r := strings.Split(tagSubStr[1], ",")
			if len(r) > 1 {
				return r[0] + "\""
			}
			return tagSubStr[1]
		}
	}
	return ""
}

func DFSTraverseStruct(input reflect.Type) {
	if input.Kind() != reflect.Struct {
		input = input.Elem()
	}
	if input.Kind() != reflect.Struct {
		panic("not struct")
	}

	var tagTotal, tagNum int

	for i := 0; i < input.NumField(); i++ {
		if string(input.Field(i).Tag) != "" {
			tagTotal++
		}
	}

	for i := 0; i < input.NumField(); i++ {
		switch input.Field(i).Type.Kind() {
		case reflect.Ptr:
			if input.Field(i).Type.Elem().Kind() == reflect.Struct {
				if tag := string(input.Field(i).Tag); tag != "" {
					tagNum++
					fmt.Printf("%s:{", splitTag(tag, "json"))
					DFSTraverseStruct(input.Field(i).Type)
					fmt.Printf("},")
				}
			}
		case reflect.Struct:
			if tag := string(input.Field(i).Tag); tag != "" {
				tagNum++
				fmt.Printf("%s:{", splitTag(tag, "json"))
				DFSTraverseStruct(input.Field(i).Type)
				fmt.Printf("},")

			}

		default:
			if tag := string(input.Field(i).Tag); tag != "" {
				tagNum++
				if tagNum < tagTotal {
					fmt.Printf("%s:null,", splitTag(tag, "json"))
				} else {
					fmt.Printf("%s:null", splitTag(tag, "json"))
				}
			}

		}
	}
}

func BFSTraverseStruct(input reflect.Type) {

	var (
		q []reflect.Type
	)

	if input.Kind() != reflect.Struct {
		input = input.Elem()
	}
	if input.Kind() != reflect.Struct {
		panic("input not struct")
	}

	if input.NumField() < 1 {
		panic("empty struct")
	}

	q = append(q, input)
	var level int
	for len(q) > 0 {
		level++
		fmt.Println(level)
		for i := 0; i < len(q); i++ {
			out := q[0]
			q = q[1:]
			for j := 0; j < out.NumField(); j++ {

				switch out.Field(j).Type.Kind() {
				case reflect.Ptr:
					if out.Field(j).Type.Elem().Kind() == reflect.Struct {
						fmt.Println(splitTag(string(out.Field(j).Tag), "json"))
						q = append(q, out.Field(j).Type.Elem())
					}
				case reflect.Struct:
					q = append(q, out.Field(j).Type)
					fmt.Println(splitTag(string(out.Field(j).Tag), "json"))

				default:
					fmt.Println(splitTag(string(out.Field(j).Tag), "json"))

				}
			}
		}
	}
}

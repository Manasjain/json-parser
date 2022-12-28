package main

import (
	"fmt"
)

type Phase int

var index int = 0

const (
	OPEN  Phase = 0
	KEY   Phase = 1
	VALUE Phase = 2
	CLOSE Phase = 3
)

var s string = "{'abc':{'d':'ef','r':'er'},'lk':'ep'}"

func main() {
	var phase Phase = 0
	mp := parseJson([]byte(s), phase)
	fmt.Println(mp)
}

func parseJson(s []byte, phase Phase) map[string]interface{} {
	var key []byte
	var value interface{}
	mp := make(map[string]interface{})

	for index < len(s) {
		// fmt.Printf("%v\n", string(s[index]))
		// if s[index] == '\'' {
		// 	fmt.Printf("%v", phase)
		// 	break
		// }
		switch phase {
		default:
			panic("unknown phase")
		case OPEN:
			if s[index] != '{' {
				panic("invalid Json at index")
			}
			index++
			phase = KEY
		case KEY:
			key = []byte{}
			if s[index] != '\'' {
				panic("invalid json at index")
			}
			index++
			for index < len(s) && s[index] != '\'' {
				key = append(key, s[index])
				index++
			}
			// fmt.Printf("%v\n", string(key))
			if index == len(s) || len(key) == 0 {
				panic("invalid json at index")
			}
			index++
			if s[index] != ':' {
				panic("invalid json at index")
			}
			index++
			phase = VALUE
		case VALUE:
			switch c := s[index]; c {
			default:
				panic("invalid json at index")
			case '\'':
				a := []byte{}
				index++
				for index < len(s) && s[index] != '\'' {
					a = append(a, s[index])
					index++
				}
				value = string(a)
				if index == len(s) {
					panic("invalid json at index")
				}
				index++
				mp[string(key)] = value
			case '{':
				phase = OPEN
				mp[string(key)] = parseJson(s, phase)
			}
			switch c := s[index]; c {
			case '}':
				phase = CLOSE
			case ',':
				phase = KEY
			default:
				panic("invalid json at index")
			}
			index++
		case CLOSE:
			return mp
		}
	}
	return mp
}

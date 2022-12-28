package main

import (
	"fmt"
)

type jsonObject string

/*
This is a basic & brute force implementation of JSON parser as stated in the problem statment below
For more advance implementation refer to json-parser.go file.
*/

/*
It is coming from untrusted source (meaning validation of json is required)
the key will always be string the value can be string or another key value pair.
Sample input {'abc':{'d':'ef','r':'er'}} -- map.get("abc").get("d") should return "ef".
No other type i.e. integer or boolean or array in the json
Validation and parsing must in done simultaneously
In case of invalid json string throw exception

Ideate
- every object starts & ends with curly braces and is combination of "key", ":", "value"
- key - starts & ends with single quotes
- value
	- can start & end with quotes
	- can be an object
- every object is comma seperated except the last one
*/
func (o jsonObject) Length() int {
	return len(string(o))
}

func main() {
	s := jsonObject("{'abc':{'d':'ef','r':'er'}}")
	var iter = 1
	ans := s.GetObject(&iter)
	fmt.Printf("%+v", ans)
}

func (s jsonObject) GetObject(iter *int) map[string]interface{} {

	mp := make(map[string]interface{})
	var key string
	// fmt.Printf("%v-%v\n", s.Length(), *iter)
	for *iter < s.Length() {
		// fmt.Printf("%v\n", string(s[*iter]))
		key = s.GetString(iter)
		*iter++
		if string(s[*iter]) == "'" {
			mp[key] = s.GetString(iter)
		} else if s[*iter] == '{' {
			*iter++
			mp[key] = s.GetObject(iter)
		}
		if s[*iter] == '}' {
			*iter++
			return mp
		}
		*iter++
		// fmt.Printf("%+v\n", mp)
	}
	return mp
}

func (s jsonObject) GetString(iter *int) string {
	var value string
	var quotesCount = 0
	for quotesCount != 2 {
		if string(s[*iter]) == "'" {
			quotesCount++
		} else {
			value += string(s[*iter])
		}
		*iter++
	}
	// fmt.Printf("%v\n", value)
	return value
}

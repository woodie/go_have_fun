///usr/bin/true; exec /usr/bin/env go run "$0" "$@" ; exit
package main

import (
	"fmt"
	"netpress.com/fun"
)

func use_string_map() {
	string_map := make(map[string]string)
  string_map["foo"] = "bar"
  string_map["bar"] = "baz"
	fmt.Println(fun.map_keys[string, string](string_map))
	fmt.Println(fun.map_values[string, string](string_map))
	fmt.Println(fun.map_include[string, string](string_map, "foo"))
	fmt.Println(fun.map_include[string, string](string_map, "baz"))
}

func use_string_set() {
	string_set:= make(map[string]bool)
  string_set["this"] = true
  string_set["that"] = true
	fmt.Println(fun.map_keys[string, bool](string_set))
	fmt.Println(fun.map_include[string, bool](string_set, "this"))
	fmt.Println(fun.map_include[string, bool](string_set, "more"))
}

func main() {
	use_string_map()
	use_string_set()
}

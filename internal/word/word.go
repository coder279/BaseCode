package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string)string{
	return strings.ToUpper(s)
}

func ToLower(s string)string{
	return strings.ToLower(s)
}

func UndercoreToUpperCamelCase(s string)string{
	s = strings.Replace(s,"_"," ",-1)
	s = strings.Title(s)
	return strings.Replace(s," ","",-1)
}

func UndercoreToLowerCamelCase(s string)string{
	s = UndercoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelCaseToUnderscore(s string)string{
	var output []rune
	for i,r := range s{
		if i == 0{
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r){
			output = append(output,unicode.ToLower(r))
		}
	}
	return string(output)
}

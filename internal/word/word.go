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
	str := UndercoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(str[0]))) + ToLower(str[1:])
}

func CamelCaseToUnderscore(s string)string{
	var output []rune
	for _,r := range s{
		if unicode.IsUpper(r){
			output = append(output,unicode.ToLower(r))
			output = append(output,'_')
		}else{
			output = append(output,unicode.ToLower(r))
		}
	}
	return string(output)
}

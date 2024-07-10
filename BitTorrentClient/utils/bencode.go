package utils

import(
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

//Bencode encodes a GO value into bencoded format
func Bencode(value interface{})(string, error){
	switch v:= value.(type) {
	case string:
		return strconv.Itoa(len(v)) + ":" + v, nil
	case int:
		return "i" + strconv.Itoa(v) + "e", nil
	case []interface{}:
		var buf bytes.Buffer
		buf.WriteString("l")
		for _, elem := range v {
			enc, err := Bencode(elem)
			if err!=nil {
				return "",err
			}
			buf.WriteString(enc)
		}
		buf.WriteString("e")
		return buf.String(), nil
	case map[string]interface{}:
		var buf bytes.Buffer
		buf.WriteString("d")
		for key, val := range v {
			encKey, err := Bencode(key)
			if err!= nil {
				return "", err
			}
			buf.WriteString(encKey)
			encVal, err := Bencode(val)
			if err!=nil {
				return "", err
			}
			buf.WriteString(encVal)
		}
		buf.WriteString("e")
		return buf.String(), nil
	default: 
		return "", errors.New("Unsupported Type!")
	}
	
}
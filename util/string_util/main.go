package string_util

import "encoding/base64"

func Base64StringToBytes(base64Value *string) []byte {
	return []byte(*base64Value)
}

func StringToBase64ByteArray(value *string) []byte {
	base64Value := base64.StdEncoding.EncodeToString([]byte(*value))

	return []byte(base64Value)
}

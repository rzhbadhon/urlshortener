package utils

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Shortner(num int) string {
	if num == 0{
		return string(alphabet[0])
	}
	result := ""
	for num > 0 {
		result = string(alphabet[num%62]) + result
		num = num/62
	}
	return result
}

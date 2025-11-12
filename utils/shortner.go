package utils

func Shortner(reqUrl string) (string, error) {

	num := 1234567890
	result := ""
	for num > 0 {
		result = string(reqUrl[num%62]) + result
		num /= 62
	}

	return result, nil
}

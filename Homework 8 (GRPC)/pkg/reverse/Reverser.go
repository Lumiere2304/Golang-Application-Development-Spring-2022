package reverse

func ReverseString(text string) string {
	var reversedText string
	for _, v := range text {
		reversedText = string(v) + reversedText
	}
	return reversedText
}
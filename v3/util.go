package env

func joinStringValues(values []string) string {
	return joinStrings(values, ", ", " and ")
}

func joinStrings(values []string, sepRune, sepWord string) string {
	text := ""
	for index := 0; index < len(values)-1; index++ {
		if index > 0 {
			text += sepRune
		}
		text += "'" + values[index] + "'"
	}
	text += sepWord + "'" + values[len(values)-1] + "'"
	return text
}

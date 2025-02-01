package utils

func Unleetify(s string) string {
	leetMap := map[rune]rune{
		'4': 'a',
		'@': 'a',
		'8': 'b',
		'6': 'b',
		'3': 'e',
		'1': 'i',
		'0': 'o',
		'5': 's',
		'$': 's',
		'7': 't',
		'+': 't',
		'2': 'z',
		'9': 'g',
		'#': 'h',
		'&': 'g',
		'(': 'c',
		'<': 'c',
		'[': 'c',
		'{': 'c',
		'|': 'l',
		'£': 'l',
		'µ': 'u',
		'¥': 'y',
		'¿': '?',
		'¡': 'i',
		'π': 'h', // for thunder god
	}

	runes := []rune(s)

	for i, r := range runes {
		if replacement, ok := leetMap[r]; ok {
			runes[i] = replacement
		}
	}

	return string(runes)
}

package romanConvert

import "errors"

var roman = map[string][]string{
	"th": {"", "M", "MM", "MMM"},
	"hu": {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	"te": {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	"un": {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
}

func Arab2roman(num int) (string, error) {
	if num < 1 || num > 3999 {
		return "error convert", errors.New("error convert")
	}

	th := num / 1000
	hu := num / 100 % 10
	te := num / 10 % 10
	un := num % 10
	res := roman["th"][th] + roman["hu"][hu] + roman["te"][te] + roman["un"][un]
	return res, nil
}

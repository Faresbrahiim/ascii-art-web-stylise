package asciiart

import (
	"os"
	"strings"
)

func Checkchars(s string) bool {
	for _, c := range s {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}

func MapBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile("banners/" + filename + ".txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\r\n")

	banner := make(map[rune][]string)
	charIndex := 32
	for i := 0; i < len(lines); i += 9 {
		if i+8 < len(lines) {
			banner[rune(charIndex)] = lines[i+1 : i+9]
			charIndex++
		}
	}
	return banner, nil
}

func Checknewline(inpultsplit []string) bool {
	c := 0
	for _, line := range inpultsplit {
		if len(line) != 0 {
			c++
		}
	}
	if c == 0 {
		return true
	} else {
		return false
	}
}
func CheckInput(input string) string {
	var str string
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			if i < len(input)-1 && input[i] == '\r' || input[i] == '\n' {
				str += string(input[i])
			}
		} else {
			str += string(input[i])
		}
	}

	return str
}
func Draw(banner map[rune][]string, inpultsplit []string) string {
	var output string
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			output += "\n"
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			output += "\n"
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(v); j++ {
					output += (banner[rune(v[j])][i])
				}
				output += "\n"
			}
		}
	}
	return output
}

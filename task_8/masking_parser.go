package task8

import (
	"regexp"
	"unicode/utf8"
)

var (
	nameExpression  = regexp.MustCompile(`\bname=([\p{L}\p{M}'’-]+)([ \t]+)([\p{L}\p{M}'’-]+)(([ \t]+)([\p{L}\p{M}'’-]+)(=)?)?`)
	loginExpression = regexp.MustCompile(`\blogin=([\p{L}\p{N}_.-]+)`)
	emailExpression = regexp.MustCompile(`\bemail=([^\s@]+)@([^\s]+)`)
	cardExpression  = regexp.MustCompile(`\bcard=([0-9]{4})([ -]?)([0-9]{4})([ -]?)([0-9]{4})([ -]?)([0-9]{4})\b`)
)

type MaskingParser struct{}

func NewMaskingParser() *MaskingParser {
	return &MaskingParser{}
}

func (p *MaskingParser) Mask(line string) string {
	return Mask(line)
}

func Mask(line string) string {
	line = nameExpression.ReplaceAllStringFunc(line, func(match string) string {
		return maskName(nameExpression.FindStringSubmatch(match))
	})
	line = loginExpression.ReplaceAllStringFunc(line, func(match string) string {
		return maskLogin(loginExpression.FindStringSubmatch(match))
	})
	line = emailExpression.ReplaceAllStringFunc(line, func(match string) string {
		return maskEmail(emailExpression.FindStringSubmatch(match))
	})
	return cardExpression.ReplaceAllStringFunc(line, func(match string) string {
		return maskCard(cardExpression.FindStringSubmatch(match))
	})
}

func maskLogin(parts []string) string {
	return "login=" + maskWord(parts[1])
}

func maskName(parts []string) string {
	masked := "name=" + maskWord(parts[1]) + parts[2] + parts[3]
	if parts[6] == "" {
		return masked
	}

	if parts[7] == "=" {
		return masked + parts[5] + parts[6] + parts[7]
	}
	return masked + parts[5] + maskWord(parts[6])
}

func maskEmail(parts []string) string {
	firstRune, _ := utf8.DecodeRuneInString(parts[1])
	return "email=" + string(firstRune) + "***@" + parts[2]
}

func maskCard(parts []string) string {
	return "card=" + parts[1] + parts[2] + parts[3][:2] + "**" + parts[4] + "****" + parts[6] + parts[7]
}

func maskWord(value string) string {
	runes := []rune(value)
	if len(runes) <= 1 {
		return value
	}
	if len(runes) <= 7 {
		masked := make([]rune, len(runes))
		masked[0] = runes[0]
		for i := 1; i < len(masked); i++ {
			masked[i] = '*'
		}
		return string(masked)
	}
	return string(runes[:2]) + "***" + string(runes[len(runes)-1])
}

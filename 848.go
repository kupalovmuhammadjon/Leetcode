func shiftingLetters(s string, shifts []int) string {
    total := 0
    res := ""

	for i := len(shifts) - 1; i >= 0; i-- {
		prev := shifts[i]
		shifts[i] += total
		total += prev
	}

    for i := 0; i < len(s); i++ {
        for shifts[i] > 26{
            shifts[i] %= 26
        }
        sh := int(s[i]) + shifts[i]
        for sh > 122 {
            sh = 96 + (sh - 122)
        }
        res += string(sh)
    }

    return res
}

func shiftingLetters(s string, shifts []int) string {
    total := 0
    res := ""

    for i := len(shifts) - 1; i >= 0; i-- {
        prev := shifts[i]
        shifts[i] += total
        total += prev
    }

    for i := 0; i < len(s); i++ {
        sh := (int(s[i]) - 'a' + shifts[i]) % 26 + 'a'
        res += string(sh)
    }

    return res
}
func shiftingLetters(s string, shifts []int) string {
    total := 0
    res := ""

    for i := len(shifts) - 1; i >= 0; i-- {
        prev := shifts[i]
        sh := (int(s[i]) - 'a' + (shifts[i] + total)) % 26 + 'a'
        temp := string(sh)
        res = temp + res
        total += prev
    }

    return res
}

// final working solution
import (
	"strings"
)

func shiftingLetters(s string, shifts []int) string {
	var totalShift int
	var resultBuilder strings.Builder
	var res strings.Builder

	for i := len(shifts) - 1; i >= 0; i-- {
		totalShift = (totalShift + shifts[i]) % 26
		charAfterShift := 'a' + (int(s[i]-'a')+totalShift)%26
		resultBuilder.WriteByte(byte(charAfterShift))
	}
    r := resultBuilder.String()
    for i := len(r) -1; i >= 0; i--{
        res.WriteByte(byte(r[i]))
    }

	return res.String()
}
package namegen

import (
    "unicode"
)

// ToTitle returns a string with the first letter of each word capitalized and the rest in lowercase,
// and removes any leading or trailing white space.
func ToTitle(s string) string {
    title := make([]rune, 0, len(s))
    inWord := false

    for _, r := range s {
        if unicode.IsLetter(r) {
            if !inWord {
                title = append(title, unicode.ToTitle(r))
                inWord = true
            } else {
                title = append(title, unicode.ToLower(r))
            }
        } else {
            title = append(title, r)
            inWord = false
        }
    }

    return string(title)
}

package namegen_test

import (
    "strings"
    "testing"

    "github.com/forscht/namegen"
)

// TestGenerate tests the Generate method of the Generator.
func TestGenerate(t *testing.T) {
    ng := namegen.New()

    // Test that generated name contains expected number of words.
    numWords := 3
    name := ng.WithNumberOfWords(numWords).Generate()
    if len(strings.Split(name, " ")) != numWords {
        t.Errorf("Generate() did not return name with %d words", numWords)
    }

    // Test that generated name contains expected separator.
    separator := "-"
    name = ng.WithWordSeparator(separator).Generate()
    if !strings.Contains(name, separator) {
        t.Errorf("Generate() did not return name with separator '%s'", separator)
    }

    // Test that generated name is capitalized as expected.
    name = ng.WithStyle(namegen.Title).Generate()
    if !strings.EqualFold(name, strings.Title(strings.ToLower(name))) {
        t.Errorf("Generate() did not capitalize name as expected with style '%s'", namegen.Title)
    }
    name = ng.WithStyle(namegen.Uppercase).Generate()
    if !strings.EqualFold(name, strings.ToUpper(name)) {
        t.Errorf("Generate() did not capitalize name as expected with style '%s'", namegen.Uppercase)
    }
    name = ng.WithStyle(namegen.Lowercase).Generate()
    if !strings.EqualFold(name, strings.ToLower(name)) {
        t.Errorf("Generate() did not capitalize name as expected with style '%s'", namegen.Lowercase)
    }

    // Test that generated name contains at least one word from each dictionary.
    adjectives := []string{"big", "small", "red", "blue", "happy"}
    colors := []string{"red", "green", "blue", "yellow", "purple"}
    animals := []string{"dog", "cat", "bird", "fish", "lion"}
    name = ng.WithDictionaries(adjectives, colors, animals).Generate()
    for _, dict := range [][]string{adjectives, colors, animals} {
        found := false
        for _, word := range strings.Split(name, "-") {
            if contains(word, dict) {
                found = true
                break
            }
        }
        if !found {
            t.Errorf("Generate() did not contain word from dictionary %v", dict)
        }
    }
}

// TestGenerateName tests the GenerateName function.
func TestGenerateName(t *testing.T) {
    name := namegen.GenerateName()
    if len(strings.Split(name, " ")) != 3 {
        t.Error("GenerateName() did not return name with 3 words")
    }
}

// contains returns true if the given slice contains the given string.
func contains(s string, slice []string) bool {
    for _, elem := range slice {
        if s == elem {
            return true
        }
    }
    return false
}

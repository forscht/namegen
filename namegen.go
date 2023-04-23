// Package namegen provides a simple library for generating random names.
package namegen

import (
    "math/rand"
    "strings"
    "time"

    "github.com/forscht/namegen/dictionaries"
)

// Constants for name capitalization styles
const (
    Title     = "title"
    Uppercase = "uppercase"
    Lowercase = "lowercase"
)

// Generator represents a name generator instance.
type Generator struct {
    numWords     int
    separator    string
    dictionaries [][]string
    style        string
    rand         *rand.Rand
}

// GenerateName generates a new random name with default settings.
func GenerateName() string {
    return New().Generate()
}

// Generate generates a new random name based on the settings of the generator.
func (ng *Generator) Generate() string {
    // If no dictionaries provided add default
    if len(ng.dictionaries) == 0 {
        ng.dictionaries = append(ng.dictionaries, dictionaries.Adjectives)
        ng.dictionaries = append(ng.dictionaries, dictionaries.Colors)
        ng.dictionaries = append(ng.dictionaries, dictionaries.Animals)
    }
    // If numWords is not set, set default to the number of dictionaries provided
    if ng.numWords <= 0 {
        ng.numWords = len(ng.dictionaries)
    }
    // Generate one random word from each dictionary
    words := make([]string, ng.numWords)
    for i := 0; i < ng.numWords; i++ {
        idx := i % len(ng.dictionaries)
        wordIdx := ng.rand.Intn(len(ng.dictionaries[idx]))
        words[i] = ng.dictionaries[idx][wordIdx]
    }
    name := strings.Join(words, ng.separator)

    // Update name style based on settings
    if ng.style == Title {
        name = ToTitle(name)
    } else if ng.style == Lowercase {
        name = strings.ToLower(name)
    } else if ng.style == Uppercase {
        name = strings.ToUpper(name)
    }

    // Return name
    return name
}

// New creates a new name generator instance with default settings.
func New() *Generator {
    return &Generator{
        numWords:     0,
        separator:    " ",
        dictionaries: [][]string{},
        style:        Title,
        rand:         rand.New(rand.NewSource(time.Now().UnixNano())),
    }
}

// WithNumberOfWords sets the number of words in the generated name.
func (ng *Generator) WithNumberOfWords(numWords int) *Generator {
    ng.numWords = numWords
    return ng
}

// WithWordSeparator sets the separator between words in the generated name.
func (ng *Generator) WithWordSeparator(separator string) *Generator {
    ng.separator = separator
    return ng
}

// WithDictionary sets the dictionaries to be used for generating names.
func (ng *Generator) WithDictionaries(dictionaries ...[]string) *Generator {
    ng.dictionaries = dictionaries
    return ng
}

// WithStyle sets the capitalization style for the generated name.
func (ng *Generator) WithStyle(style string) *Generator {
    ng.style = style
    return ng
}

// WithSeed sets the seed for the random number generator used for generating names.
func (ng *Generator) WithSeed(seed int64) *Generator {
    ng.rand = rand.New(rand.NewSource(seed))
    return ng
}

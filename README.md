# Namegen

Namegen is a Go package for generating random names. It provides a simple library for generating names based on different settings like the number of words, the separator between words, and the capitalization style.

## Installation
To install the namegen package, use the following command:
```shell
go get github.com/forscht/namegen
```

## Usage
To use the namegen package, import it in your Golang code and create a new Generator object. You can then customize the generator using the various With methods to set the number of words, word separator, dictionaries, style, and seed. Finally, call the Generate method to generate a new random name based on the settings of the generator.

```go
package main

import (
    "fmt"
    "time"

    "github.com/forscht/namegen"
    "github.com/forscht/namegen/dictionaries"
)

func main() {
    // Generate a new name
    name := namegen.New().
        WithNumberOfWords(3). // Number of words. Default 3
        WithWordSeparator("-"). // Words Separator, Default " "
        WithStyle(namegen.Title). // Default Title, Opt Title,UpperCase,LowerCase
        WithDictionaries(dictionaries.Adjectives, []string{"red", "blue", "green"}, dictionaries.Animals). // Use Default Dictionaries
        WithSeed(time.Now().UnixNano()). // Use custom seed
        Generate()
    fmt.Println(name) // Output: Evil-Red-Hyena

    // Generate with Default settings
    name = namegen.GenerateName()
    fmt.Println(name) // Output: Real Purple Barracuda

    // Use same settings to use multiple names
    ng := namegen.New().WithStyle(namegen.Lowercase).WithNumberOfWords(2)
    name1 := ng.Generate()
    name2 := ng.Generate()
    fmt.Println(name1) // Output: top magenta
    fmt.Println(name2) // Output: adverse burgundy
}
```

## API
Namegen provides the following methods for generating names:

- GenerateName() string: generate a new random name with default settings
- New() *Generator: creates a new generator instance with default settings

Namegen provides the following methods for customizing the Generator object:

- WithNumberOfWords(numWords int): set the number of words in the generated name
- WithWordSeparator(separator string): set the separator between words in the generated name
- WithDictionary(dictionaries ...[]string): set the dictionaries to be used for generating names
- WithStyle(style string): set the capitalization style for the generated name
- WithSeed(seed int64): set the seed for the random number generator used for generating names

## License
This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

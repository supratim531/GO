package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	languages := make(map[string]string)

	languages["Go"] = "Google's programming language"
	languages["Python"] = "An interpreted, high-level programming language"
	languages["Java"] = "A statically-typed, object-oriented programming language"
	languages["C++"] = "A statically-typed, object-oriented programming language"
	languages["JavaScript"] = "An interpreted, high-level programming language"
	languages["Ruby"] = "An interpreted, high-level programming language"
	languages["Swift"] = "A compiled, general-purpose programming language"
	languages["Kotlin"] = "A statically-typed, object-oriented programming language"
	languages["Rust"] = "A statically-typed, systems programming language"
	languages["C#"] = "A statically-typed, object-oriented programming language"
	languages["Delphi"] = "A procedural, imperative, object-oriented programming language"
	languages["VBA"] = "A procedural, object-oriented programming language"
	languages["Prolog"] = "A logic programming language"
	languages["Erlang"] = "A concurrent, functional programming language"
	languages["Haskell"] = "A purely functional programming language"
	languages["OCaml"] = "A statically-typed, functional programming language"
	languages["F#"] = "A statically-typed, functional programming language"
	languages["Clojure"] = "A functional programming language"
	languages["Lisp"] = "A multi-paradigm, functional programming language"
	languages["Scheme"] = "A functional programming language"
	languages["Julia"] = "A high-level, general-purpose programming language"
	languages["R"] = "A statistical programming language"
	languages["MATLAB"] = "A programming language for technical computing"

	fmt.Println("Total number of languages is", len(languages))
	fmt.Println("Enter the name of your favourite language: ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	language, _ := reader.ReadString('\n')
	language = strings.TrimSpace(language)
	fmt.Println("The description of", language, "is:", languages[language])

	fmt.Println("Which language you want to delete: ")
	deleteLanguage, _ := reader.ReadString('\n')
	delete(languages, strings.TrimSpace(deleteLanguage))
	fmt.Println("Total number of languages is", len(languages))

	// loops are interesting in GO
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}

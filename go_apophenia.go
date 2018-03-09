package main

// remember to only list what you're going to use
import (
    "fmt"
    "io/ioutil"
)


// make use of err because if you don't you CAN'T build ANYTHING
func check(e error) {
    if e != nil {
        // "panic" descibes my anxiety when coding in GO
        panic(e)
    }
}


// theres going to be more to make sense of this
func create_wordlist() {
    // you're required to set vars for both the file and errors
    // SO annoying
    file, err := ioutil.ReadFile("words_alpha.txt")
    fmt.Print(string(file))
    check(err)
}


// main function because that what gets ran as default
func main() {
    create_wordlist()
}

[set-up Visual Studio Code](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code)

[set-up go Dockerfile](https://stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire)

Use the [go playground](https://go.dev/play/)

Youtube - [Learn Go in 7 minutes](https://www.youtube.com/watch?v=45oXTpoJ5-I)
  [Goroutines by example](https://gobyexample.com/goroutines)
  [Channels by example](https://gobyexample.com/channels)

[Learning Go](https://go.dev/learn/)

[Official Go Docs](https://go.dev/doc/)

* Tutorial: Getting started
* Tutorial: Create a module: Code in ./1_go.dev_doc_tutorial_create-a-module
* Tutorial: Getting started with multi-module workspaces: Code in ./2_go.dev_doc_tutorial_multi-module-workspaces
* Tutorial: Developing a RESTful API with Go and Gin: Code in ./3_go.dev_doc_tutorial_restful-api-with-go-and-gin
* Tutorial: Getting started with generics: Code in ./4_go.dev_doc_tutorial_generics
* Tutorial: Getting started with fuzzing: Code in ./5_go.dev_doc_tutorial_fuzzing

# Resources I went through in order, some concurrently

[Effective Go](https://go.dev/doc/effective_go)

[Code Review Comments](https://go.dev/wiki/CodeReviewComments)

[The Language Specification](https://go.dev/ref/spec)

[Go Tour](https://go.dev/tour/)

Learning Go: An Idiomatic Approach to Real-World Go Programming  
Another edition is expected out in 2024  
  * ([source](https://github.com/learning-go-book))
  * [My notes](LearningGo-AnIdiomaticApproachToRealWorldGoProgramming/README.md)  

Review of the best Go books for 2023:
* https://www.reddit.com/r/golang/comments/11hd310/what_would_be_the_best_golang_book_to_read_in/
* https://boldlygo.tech/posts/2023-01-30-review-learning-go/
* https://boldlygo.tech/posts/2023-02-24-best-book-to-learn-go-in-2023/

"The Go Programming Language"

* What go is and has according to the Preface:
  * Apparently Go is a high-level language, but it has basically no functional programming features
  * Garbage collection
  * A package system
  * First-class functions
  * lexical scope
  * A system call interface
  * Immutable strings in which text is generally encoded in UTF-8
  * No implicit numeric conversions
  * No constructors or destructors
  * No operator overloading
  * No default parameter values
  * No inheritance
  * No generics (Although these are added in 1.18)
  * No exceptions
  * No macros
  * No function annotations
  * No thread-local storage

100 Go Mistakes and How to Avoid Them

* https://github.com/teivah/100-go-mistakes
* https://100go.co/



# Useful Resources

* [Standard Library of Go Packages](https://pkg.go.dev/std)
* [Go Wiki](https://go.dev/wiki/)
* [Go blog](https://go.dev/blog/)

## Generics

* Read [Go by Example: Generics](https://gobyexample.com/generics) in conjunction with the "Learn Go in 7 minutes" video above

## Functional Go

* [Functional Go](https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4)

## Constants

* [Constants](https://go.dev/tour/basics/15)
* [How To Use Variables and Constants in Go](https://www.digitalocean.com/community/tutorials/how-to-use-variables-and-constants-in-go)

There are boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants. Rune, integer, floating-point, and complex constants are collectively called numeric constants. [as per spec](https://go.dev/ref/spec#Constants).
That means Arrays, Slices, Maps, Structs, etc [can't be made constant](https://blog.boot.dev/clean-code/constants-in-go-vs-javascript-and-when-to-use-them/#constants-in-go), but there is a [work-around](https://blog.boot.dev/golang/golang-constant-maps-slices/) using initialisation functions.

# Books

* Learning Functional Programming in Go
* Functional Programming in Golang (to be released on 2023-04-11), but it's Pakt, so based on experience with Pakt, probably not much good
* Humans are very good at measuring when we have something to measure against. [Go for JavaScript Developers](http://www.pazams.com/Go-for-Javascript-Developers/) has comparisons that are quite helpful for those of us that have spent a long time in JavaScript
* [wiki books](https://go.dev/wiki/Books)


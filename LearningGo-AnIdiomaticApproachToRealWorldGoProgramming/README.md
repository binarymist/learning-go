While reading and doing exercises, etc of this book, I also look at quite a few other resources, as it helps to solidify concepts for me. When I use other resources, I specify what they are.

# Chapter 1

## The Go Workspace

"_Since the introduction of Go in 2009, there have been several changes in how Go developers organize their code and their dependencies. Because of this churn, there’s lots of conflicting advice, and most of it is obsolete._"

"_Go still expects there to be a single workspace for third-party Go tools installed via `go install`_"

"_By default, this workspace is located in $HOME/go, with source code for these tools stored in $HOME/go/src and the compiled binaries in $HOME/go/bin. You can use this default or specify a different workspace by setting the `$GOPATH` environment variable._"

"_Whether or not you use the default location, it’s a good idea to explicitly define `GOPATH` and to put the $GOPATH/bin directory in your executable path. Explicitly defining `GOPATH` makes it clear where your Go workspace is located and adding $GOPATH/bin to your executable path makes it easier to run third-party tools installed via go install_"

If using asdf, the `GOPATH` is defined for you. For example if you are using `golang 1.20` (as specified in your .tool-versions file), then your `GOPATH` will look like:  

```shell
> go env GOPATH
GOPATH="/home/<you>/.asdf/installs/golang/1.20/packages"
```



If using zsh, there are are couple of tweaks that need to be made to the .zshrc file, including adding the `GOPATH/bin` to your system path, this is done with the following line in the .zshrc:  
`export PATH=$(go env GOPATH)/bin:$PATH`

If using asdf, the following advice does not apply:

_~(If you are using zsh, add these lines to .zshrc instead):~_  
  
_~`export GOPATH=$HOME/go`~_  
_~`export PATH=$PATH:$GOPATH/bin`~_



## The `go` Command

### `go run` and `go build`

1. Builds the binary
2. Executes the binary
3. Deletes the binary after your program finishes

**`go run`** compiles your code into a binary and saves the file to a temporary directory, by default the systems default temporary directory. You can override this directory by modifying the `$GOTMPDIR` environment variable. To view all of your go environment variables, just run: `go env`.

If you want to locate the exact temporary directory where Go is storing these artifacts on your system, you can use the `os.TempDir()` function in your Go code to retrieve the path to the system's temporary directory. Here's an example:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	tempDir := os.TempDir()
	fmt.Printf("Temporary directory: %s\n", tempDir)
}
```

**`go build`**

`go build <packages>`  
or  
`go build -o <custom-binary-name> <packages>`

### Getting Third-Party Go Tools

Pg 5 of "Learning Go an Idiomatic approach..." says:  
The `go install` command takes an argument, which is the location of the source code repository of the project you want to install, followed by an `@` and the version of the tool you want (if you just want to get the latest version, use `@latest`). It then downloads, compiles, and installs the tool into your $GOPATH/bin directory.  
To view all of your go environment variables, just run: `go env`.  

It looks like the source is downloaded from a proxy (as defined by the value of the `$GOPROXY` environment variable) if it exists into:

1. $GOPATH/pkg/mod/cache/download/github.com/rakyll/hey/@v/v0.1.4.zip as an archive file
2. $GOPATH/pkg/mod/github.com/rakyll/hey@v0.1.4/ as the decompressed archive (the hydrated source)
3. and then compiled into the $GOPATH/bin directory

### Formatting Your Code

Also discussed in [Effective Go: Formatting](https://go.dev/doc/effective_go#formatting).

"_Go programs use tabs to indent, and it is a syntax error if the opening brace is not on the same line as the declaration or command that begins the block._"

`go fmt` "_reformats your code to match the standard format_"

VS Code with the official Golang plugin uses `goimports` which is better.

Like JavaScript, the Go lexer uses semicolon insertion (ASI). Unlike JavaScript, Go developers never add semicolons other than places such as `for` loop clauses, to separate the initialiser, condition, and continuation elements. [They are also necessary to separate multiple statements on a line, should you write code that way](https://go.dev/doc/effective_go#semicolons).

## Linting and Vetting

golint is deprecated, drop-in replacement is https://github.com/mgechev/revive .  
Use `go vet` instead.  
golangci-lint runs [multiple (configurable which)](https://golangci-lint.run/usage/linters/) tools in parallel

The recommendation is: start off using `go vet`` as a required part of your automated build process and revive as part of your code review process (since revive might have false positives and false negatives, you can’t require your team to fix every issue it reports). Once you are used to their recommendations, try out golangci-lint and tweak its settings until it works for your team.

## Choose Your Tools

The Go wiki [Go IDEs and Editors](https://go.dev/wiki/IDEsAndTextEditorPlugins) Discusses:

* Visual Studio Code  
  Uses:  
  * the [Delve](https://github.com/go-delve/delve) debugger
  * [gopls](https://github.com/golang/tools/blob/master/gopls/README.md), the official Go language server
* GoLand IDE from JetBrains
* The [Go Playground](http://play.golang.org)

## Makefiles

Go developers typically use Makefiles.

## Staying Up to Date

If using asdf, just follow the asdf documentation. You can have as many versions of Go installed as you like, and specific versions can be run from specific directories and below simply by adding a .tool-versions file (containing the binary's name and version) in the directory of your choosing.

# Chapter 2. Primitive Types and Declarations

## Built-in Types

### The Zero Value

Go assigns a default zero value to any variable that is declared but not assigned a value.






## Using `const`

Also covered in:

* [Effective Go: Initialization](https://go.dev/doc/effective_go#initialization)

Created at compile time, as opposed to variables at runtime.

## Unused Variables

[Effective Go: Unused imports and variables](https://go.dev/doc/effective_go#blank_unused) discusses temporarily reading and storing to the `_`, this allows the compilation to succeed.

## Naming Variables and Constants

Also discussed in The Go wiki [CodeReviewComments#variable-names](https://go.dev/wiki/CodeReviewComments#variable-names), which to an extent goes against what history has taught us about non-descriptive variable names. Bob Martin, Steve McConnell, other greats, even myself, have written about short, unmemorable names many times, favouring write-time convenience over read-time convenience is often a mistake.
Use descriptive names rather than adding non-executable comments to explain what something does. Although I agree that short variable names are fine in some cases, such cases are mentioned in this Go wiki section.

Initialisms are covered in:

* The Go wiki [CodeReviewComments#initialisms](https://go.dev/wiki/CodeReviewComments#initialisms)
* Pg 28, **2.1. Names** of "The Go Programming Language"

MixedCaps is also discussed in:

* The Go wiki [CodeReviewComments#mixed-caps](https://go.dev/wiki/CodeReviewComments#mixed-caps)
* [Effective Go: MixedCaps](https://go.dev/doc/effective_go#mixed-caps)

No snake_case, use PascalCase for exported identifiers, use camelCase for unexported identifiers.

# Chapter 3. Composite Types

[Effective Go: Constructors and composite literals](https://go.dev/doc/effective_go#composite_literals) discusses composite literals.

## Arrays—Too Rigid to Use Directly

Also covered in:

* [Effective Go: Arrays](https://go.dev/doc/effective_go#arrays)
* Pg 81, **4.1. Arrays** of "The Go Programming Language"
* [Go for JavaScript Developers: Types - Arrays / Slices](http://www.pazams.com/Go-for-Javascript-Developers/pages/types/#d-arrays--slices) Discusses the differences between arrays in Go vs JavaScript, including how slices fit
* [Go by Example: Arrays](https://gobyexample.com/arrays)

## Slices

Also covered in:

* [Effective Go: Slices](https://go.dev/doc/effective_go#slices)
* Pg 84, **4.2. Slices** of "The Go Programming Language"
* [Go by Example: Slices](https://gobyexample.com/slices)
* "100 Go Mistakes and How to Avoid Them" [Not understanding slice length and capacity (#20)](https://100go.co/#not-understanding-slice-length-and-capacity-20)
* "100 Go Mistakes and How to Avoid Them" [Inefficient slice initialization (#21)](https://100go.co/#inefficient-slice-initialization-21)
* "100 Go Mistakes and How to Avoid Them" [Being confused about nil vs. empty slice (#22)](https://100go.co/#being-confused-about-nil-vs-empty-slice-22)
* "100 Go Mistakes and How to Avoid Them" [Not properly checking if a slice is empty (#23)](https://100go.co/#not-properly-checking-if-a-slice-is-empty-23)
* "100 Go Mistakes and How to Avoid Them" [Not making slice copies correctly (#24)](https://100go.co/#not-making-slice-copies-correctly-24)
* "100 Go Mistakes and How to Avoid Them" [Unexpected side effects using slice append (#25)](https://100go.co/#unexpected-side-effects-using-slice-append-25)
* "100 Go Mistakes and How to Avoid Them" [Slices and memory leaks (#26)](https://100go.co/#slices-and-memory-leaks-26)
* "100 Go Mistakes and How to Avoid Them" [Slices and memory leaks (#26)](https://100go.co/#slices-and-memory-leaks-26)

Go's arrays and slices are single-dimensional, you can however create the equivalent of multi-dimensional slices or arrays. Inspiration for the following example was from [Effective Go: Two-dimensional slices](https://go.dev/doc/effective_go#two_dimensional_slices):

```go
package main

import (
	"fmt"
)

func main() {
	var slice [][][]byte
	slice = [][][]byte{
		{
			{'c'},
			{'a', 't'},
		},
		{
			{'d'},
			{'o', 'g'},
		},
	}
	fmt.Println("Contents of the three-dimensional slice:")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Level %d:\n", i+1)
		for j := 0; j < len(slice[i]); j++ {
			fmt.Println(string(slice[i][j]))
		}
		fmt.Println()
	}
}
```

The Go wiki [CodeReviewComments#declaring-empty-slices](https://go.dev/wiki/CodeReviewComments#declaring-empty-slices) also discusses nil slices vs non-nil, and when to use which.

```go
// Declares a nil slice:
var t []string
```

vs

```go
//Declares a non-nill slice:
t := []string{}
```

### `append`

Also covered in:

* The [built-in documentation](https://pkg.go.dev/builtin#append)
* The spec under [Appending to and copying slices](https://go.dev/ref/spec#Appending_and_copying_slices)
* [Effective Go: `append`](https://go.dev/doc/effective_go#append)
* Pg 88, **4.2.1. The `append` Function** of "The Go Programming Language"
* "100 Go Mistakes and How to Avoid Them" [Unexpected side effects using slice append (#25)](https://100go.co/#unexpected-side-effects-using-slice-append-25)
* "100 Go Mistakes and How to Avoid Them" [Creating data races with append (#69)](https://100go.co/#creating-data-races-with-append-69)

### `make`

The built-in function [`make`](https://pkg.go.dev/builtin#make) creates slices, maps, and channels only.

[Effective Go: Allocation with make](https://go.dev/doc/effective_go#allocation_make), has some low value examples.

## Maps

Also covered in:

* [Effective Go: Maps](https://go.dev/doc/effective_go#maps), which also covers "The comma ok Idiom", discussed soon in "Learning Go an Idiomatic approach..."
* Pg 93, **4.3. Maps** of "The Go Programming Language"
* [Go for JavaScript Developers: Types - Dictionaries](http://www.pazams.com/Go-for-Javascript-Developers/pages/types/#d-dictionaries) Discusses the differences between how and what JavaScript and Go implement and use Dictionaries in the form of JS Object, JS Map, JS WeakMap, and Go Maps
* "100 Go Mistakes and How to Avoid Them" [Inefficient map initialization (#27)](https://100go.co/#inefficient-map-initialization-27)
* "100 Go Mistakes and How to Avoid Them" [Maps and memory leaks (#28)](https://100go.co/#maps-and-memory-leaks-28)
* "100 Go Mistakes and How to Avoid Them" [Making wrong assumptions during map iterations (ordering and map insert during iteration) (#33)](https://100go.co/#making-wrong-assumptions-during-map-iterations-ordering-and-map-insert-during-iteration-33)
* [Go by Example: Maps](https://gobyexample.com/maps)
* The spec under [Map types](https://go.dev/ref/spec#Map_types) discusses the types that can be used for map keys and values
* The spec under [Deletion of map elements](https://go.dev/ref/spec#Deletion_of_map_elements) discusses deleting map elements
* The spec under [Making slices, maps and channels](https://go.dev/ref/spec#Making_slices_maps_and_channels) discusses the creation of maps using the built-in `make` function

# Chapter 4. Blocks, Shadows, and Control Structures

## Blocks

This section should have discussed Redeclaration and reassignment, thankfully we have coverage in [Effective Go: Redeclaration and reassignment](https://go.dev/doc/effective_go#redeclaration), be sure to read it.

## `if`

Also discussed in [Effective Go: if](https://go.dev/doc/effective_go#if)

## `for`, Four Ways

Also discussed in [Effective Go: for](https://go.dev/doc/effective_go#for)

## `switch`

Also covered in:

* [Effective Go: Switch](https://go.dev/doc/effective_go#switch)
* [Go for JavaScript Developers: Flow control statements - Switch](http://www.pazams.com/Go-for-Javascript-Developers/pages/flow-control-statements/#d-switch), specifically the fact that in Go, `fallthrough` must be specified if desired
* Pg 23, **1.8. Loose Ends** of "The Go Programming Language"

# Chapter 5. Functions

## Declaring and Calling Functions

### Multiple Return Values

Also covered in:

* [Effective Go: Multiple return values](https://go.dev/doc/effective_go#multiple-returns), bit of a stupid example though
* The Go Wiki [CodeReviewComments#dont-panic](https://go.dev/wiki/CodeReviewComments#dont-panic) discusses using error and multiple return values
* Pg 124, **5.3. Multiple Return Values** of "The Go Programming Language"
* [Go by Example: Multiple Return Values](https://gobyexample.com/multiple-return-values)
* [Go for JavaScript Developers: Functions - Multiple returns](http://www.pazams.com/Go-for-Javascript-Developers/pages/functions/#d-multiple-returns)

### Named Return Values

Also covered in:

* [Effective Go: Named result parameters](https://go.dev/doc/effective_go#named-results), bit of a stupid example though
* Also called named result parameters, in The Go Wiki [CodeReviewComments#named-result-parameters](https://github.com/gvolang/go/wiki/CodeReviewComments#named-result-parameters), which is mostly about how it affects godoc
* "100 Go Mistakes and How to Avoid Them" [Never using named result parameters (#43)](https://100go.co/#never-using-named-result-parameters-43)
* "100 Go Mistakes and How to Avoid Them" [Unintended side effects with named result parameters (#44)](https://100go.co/#unintended-side-effects-with-named-result-parameters-44)

## `defer`

Also covered in:

* [Effective Go: Defer](https://go.dev/doc/effective_go#defer) has an interesting last example
* [Go by Example: Defer](https://gobyexample.com/defer)
* Pg 143, **5.8. Deferred Function Calls** of "The Go Programming Language"
* "100 Go Mistakes and How to Avoid Them" [Using defer inside a loop (#35)](https://100go.co/#using-defer-inside-a-loop-35)
* "100 Go Mistakes and How to Avoid Them" [Ignoring how defer arguments and receivers are evaluated (argument evaluation, pointer, and value receivers) (#47)](https://100go.co/#ignoring-how-defer-arguments-and-receivers-are-evaluated-argument-evaluation-pointer-and-value-receivers-47)
* "100 Go Mistakes and How to Avoid Them" [Not handling defer errors (#54)](https://100go.co/#not-handling-defer-errors-54)

### Blank Returns—Never Use These!

Also called naked returns, and covered in The Go Wiki [CodeReviewComments#naked-returns](https://go.dev/wiki/CodeReviewComments#naked-returns). Doesn't say anything about not using them though.

## Go is Call By Value

The Go Wiki [CodeReviewComments#pass-values](https://go.dev/wiki/CodeReviewComments#pass-values) says: "_Don't pass pointers as function arguments just to save a few bytes. If a function refers to its argument x only as *x throughout, then the argument shouldn't be a pointer._"

"Go for Javascript Developers" [says](https://www.pazams.com/Go-for-Javascript-Developers/pages/types/): "_In Go, there are value types, reference types, and pointers. References types are slices, maps, and channels. All the rest are value types, but have the ability "to be referenced" with pointers._"

Because Slices in Go are reference types, a slice's value is a reference.
Any modifications made to the slice within the function will affect the original slice in the calling function. However, when you append an element to a slice, a new underlying array is created if the capacity of the current array is exceeded, and the original slice's reference does not change.
In the `modSlice` function, when you append 10 to `s`, it exceeds the capacity of the original array, and a new array is created to accommodate the additional element. The `s` slice inside the `modSlice` function now points to a new array, but the `s` slice in the `main` function remains unchanged because slices are passed by value (a copy of the slice header is passed) and not by reference.

# Chapter 6. Pointers

## A Quick Pointer Primer

"Learning Go an Idiomatic approach..." has an example (`var x = new(int)`) of the [`new` built-in function](https://pkg.go.dev/builtin#new), which is also covered in:

* [Effective Go: Allocation with new](https://go.dev/doc/effective_go#allocation_new)
* The spec under [Slice types](https://go.dev/ref/spec#Slice_types) also has an interesting example that the following "_two expressions are equivalent:_"
   
   ```go
   make([]int, 50, 100)
   new([100]int)[0:50]
   ```
   
   The [spec](https://go.dev/ref/spec) also has some other examples of the `new` built-in
* Pg 34, **2.3.3. The `new` Function** of "The Go Programming Language"
* [Go for JavaScript Developers: Keywords & Syntax Comparison - new keyword](http://www.pazams.com/Go-for-Javascript-Developers/pages/keywords-syntax-comparison/#d-new-keyword)

# Chapter 7. Types, Methods, and Interfaces

[Effective Go: Interface Names](https://go.dev/doc/effective_go#interface-names) discusses how to name single method interfaces.

## Methods

* The Go Wiki [CodeReviewComments#pass-values](https://go.dev/wiki/CodeReviewComments#pass-values) discusses when to pass a value receiver vs reference receiver
* [Effective Go: Initialization](https://go.dev/doc/effective_go#initialization) has an example of using a `String()` method with a constant type using the `iota` enumerator

As [Effective Go: Methods](https://go.dev/doc/effective_go#methods) mentions:

* "_methods can be defined for any named type (except a pointer or an interface); the receiver does not have to be a struct._"
* "_The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers._"

### Pointer Receivers and Value Receivers

* The Go Wiki [CodeReviewComments#receiver-names](https://go.dev/wiki/CodeReviewComments#receiver-names) discusses receiver names
* The Go Wiki [CodeReviewComments#receiver-type](https://go.dev/wiki/CodeReviewComments#receiver-type) discusses when to use value vs pointer receiver type
* "100 Go Mistakes and How to Avoid Them" [Not knowing which type of receiver to use (#42)](https://100go.co/#not-knowing-which-type-of-receiver-to-use-42) also discusses when to use value vs pointer receiver type
* Pg 158, **6.2. Methods with a Pointer Receiver** of "The Go Programming Language" discusses when to pass a value receiver vs reference receiver

The following "100 Go Mistakes and How to Avoid Them" are also worth a read:

* [Returning a nil receiver (#45)](https://100go.co/#returning-a-nil-receiver-45)
* [Ignoring how defer arguments and receivers are evaluated (argument evaluation, pointer, and value receivers) (#47)
](https://100go.co/#ignoring-how-defer-arguments-and-receivers-are-evaluated-argument-evaluation-pointer-and-value-receivers-47)

"_do not write getter and setter methods for Go structs, unless you need them to meet an interface_"

Also discussed in:

* [Effective Go: Getters](https://go.dev/doc/effective_go#Getters)
* "100 Go Mistakes and How to Avoid Them" [Overusing getters and setters (#4)](https://100go.co/#overusing-getters-and-setters-4)
* Pg 169, **6.6. Encapsulation** of "The Go Programming Language"

## Use Embedding for Composition

"_You can embed any type within a struct_"  
[Effective Go: "_Only interfaces can be embedded within interfaces_"](https://go.dev/doc/effective_go#embedding)

## A Quick Lesson on Interfaces

Also covered in:

* The Go wiki [CodeReviewComments#comment-sentences](https://go.dev/wiki/CodeReviewComments#interfaces)
* [Effective Go: Interfaces](https://go.dev/doc/effective_go#interfaces)
* "100 Go Mistakes and How to Avoid Them" [Interface pollution (#5)](https://100go.co/#interface-pollution-5), [Interface on the producer side (#6)](https://100go.co/#interface-on-the-producer-side-6)
* Pg 171, **Interfaces** of "The Go Programming Language"

## Type Assertions and Type Switches

Also covered in:

* [Effective Go: Type switch](https://go.dev/doc/effective_go#type_switch) has one example.
* Pg 205, **7.10. Type Assertions** of "The Go Programming Language"
* Pg 210, **7.13. Type Switches** of "The Go Programming Language"

# Chapter 8. Errors

In-band errors are covered in the Go wiki [CodeReviewComments#in-band-errors](https://go.dev/wiki/CodeReviewComments#in-band-errors).

## Errors Are Values

Error strings are:

* Covered in the Go wiki [CodeReviewComments#error-strings](https://go.dev/wiki/CodeReviewComments#error-strings)
* Briefly touched on in [Effective Go: Errors](https://go.dev/doc/effective_go#errors)

## `panic` and `recover`

Also covered in:

* The Go wiki [CodeReviewComments#dont-panic](https://go.dev/wiki/CodeReviewComments#dont-panic) which also links to [Effective Go: Errors](https://go.dev/doc/effective_go#errors)
* [Effective Go: Panic](https://go.dev/doc/effective_go#panic)
* [Effective Go: Recover](https://go.dev/doc/effective_go#recover)
* "100 Go Mistakes and How to Avoid Them" [Panicking (#48)](https://100go.co/#panicking-48)
* Pg 148, **5.9. Panic** of "The Go Programming Language"
* [Go for JavaScript Developers: error-handling](http://www.pazams.com/Go-for-Javascript-Developers/pages/error-handling/)

# Chapter 9. Modules, Packages and Imports

## Building Packages

The Go wiki [CodeReviewComments#examples](https://go.dev/wiki/CodeReviewComments#examples) discusses adding examples.

### Imports and Exports

A topic that is somewhat thin in the "Learning Go an Idiomatic approach..." book is printing, this has coverage in:

* [Effective Go: Printing](https://go.dev/doc/effective_go#printing)

[Effective Go: Unused imports and variables](https://go.dev/doc/effective_go#blank_unused) discusses temporarily reading and storing to the `_`, this allows the compilation to succeed.

### Naming Packages

Also covered in:

* The Go wiki [CodeReviewComments#package-names](https://go.dev/wiki/CodeReviewComments#package-names) which also links to [Effective Go: Package names](https://go.dev/doc/effective_go#package-names) and [The Go Blog Package names](https://go.dev/blog/package-names)
* Pg 289, **10.6. Packages and Naming** of "The Go Programming Language"

### Overriding a Package’s Name

The Go wiki:

* [CodeReviewComments#crypto-rand](https://go.dev/wiki/CodeReviewComments#crypto-rand) discusses math/rand and crypto/rand
* [CodeReviewComments#imports](https://go.dev/wiki/CodeReviewComments#imports) discusses aliasing imports
* [CodeReviewComments#import-dot](https://go.dev/wiki/CodeReviewComments#import-dot) discusses importing all the exported identifiers (functions, variables, types, etc) from a package into the current scope, effectively allowing you to use them without specifying the package name

Also covered in "100 Go Mistakes and How to Avoid Them" [Ignoring package name collisions (#14)](https://100go.co/#ignoring-package-name-collisions-14)

### Package Comments and godoc

Also covered in:

* The Go wiki [CodeReviewComments#comment-sentences](https://go.dev/wiki/CodeReviewComments#comment-sentences) and [CodeReviewComments#doc-comments](https://go.dev/wiki/CodeReviewComments#doc-comments) which also links to [Effective Go: Commentary](https://go.dev/doc/effective_go#commentary)
* The Go wiki [CodeReviewComments#package-comments](https://go.dev/wiki/CodeReviewComments#package-comments) and [CodeReviewComments#doc-comments](https://go.dev/wiki/CodeReviewComments#doc-comments)
* "100 Go Mistakes and How to Avoid Them" [Missing code documentation (#15)](https://100go.co/#missing-code-documentation-15)
* Pg 296, **10.7.4. Documenting Packages** of "The Go Programming Language"

### The `init` Function: Avoid if Possible

Also covered in the Go wiki [CodeReviewComments#import-blank](https://go.dev/wiki/CodeReviewComments#import-blank)

# Chapter 10. Concurrency in Go

The Go wiki [CodeReviewComments#synchronous-functions](https://go.dev/wiki/CodeReviewComments#synchronous-functions) discusses preferring synchronous functions. This of course means waiting/blocking.

Go has two styles of concurrent programming:

* Goroutines and channels, which support communicating sequential processes or CSP. A go community saying: "_Share memory by communicating; do not communicate by sharing memory_". This style is covered in **8. Goroutines and Channels** of "The Go Programming Language"
* The more traditional model of shared memory multithreding. This style is covered in **9. Concurrency with Shared Variables** of "The Go Programming Language"

## When to Use Concurreny

"_If you are not sure if concurrency will help, first write your code serially, and then write a benchmark to compare performance with a concurrent implementation._"

## Goroutines

Operating System schedules threads across CPU cores. A process can own multiple threads.
a thread can own many goroutines (lightweight processes).

* Pg 205 says "_Goroutines are lightweight processes managed by the Go runtime_"
  Just about every other source says that Goroutines are more like lightweight threads, including "The Go Programming Language" pg 218

The crux of this section is that goroutines are faster and lighter (in terms of resources) because they are managed by the Go runtime scheduler instead of the operating system (as in OS managed threads). The Go runtime scheduler is obviously closer to your code, and knows more about your code than the operating system.

I have a full [example code](./ch10_ConcurrencyInGo/goroutines/) for this.

Pg 218, **8.1. Goroutines** of "The Go Programming Language" says "_The go statement itself completes immediately:_"

```go
f() // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

## Channels

As per (https://go.dev/tour/concurrency/2) By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

[Being puzzled about channel size (#67)](https://100go.co/#being-puzzled-about-channel-size-67) of "100 Go Mistakes and How to Avoid Them" says:

* "_Carefully decide on the right channel type to use, given a problem. Only unbuffered channels provide strong synchronization guarantees_". Thus they are sometimes called "_synchronous channels_" as mentioned on pg 226 of "The Go Programming Language"
* "_You should have a good reason to specify a channel size other than one for buffered channels_"

"_Like maps, channels are reference types. When you pass a channel to a function, you are really passing a pointer to the channel_"

**8.4.2. Pipelines** of "The Go Programming Language" is a pattern/technique which has a goroutine sending to a channel, that another goroutine receives. Pg 229: If the first goroutine closes the channel, after the receiving end of the channel has drained all messages, all subsequent receive operations will proceed without blocking but will yield a zero value, unless the receive operation checks the 2nd return value which is a boolean value, conventionally called `ok`, which is true for a successful receive and false for a receive on a closed and drained channel. Using a `range` loop provides the same behaviour in a more succinct manner.  
Also discussed on pg 208, **Closing a Channel** of "Learning Go an Idiomatic approach...".

Pg 230 "_You needn’t close every channel when you’ve finished with it. It’s only necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent._"

"_A channel that the garbage collector determines to be unreachable will have its resources reclaimed whether or not it is closed._". This is also mentioned on pg 208, **Closing a Channel** of "Learning Go an Idiomatic approach...".

Pg 233, **8.4.4. Buffered Channels** of "The Go Programming Language" "_Novices are sometimes tempted to use buffered channels within a single goroutine as a queue, lured by their pleasingly simple syntax, but this is a mistake. Channels are deeply connected to goroutine scheduling, and without another goroutine receiving from the channel, a sender—and perhaps the whole program—risks becoming blocked forever. If all you need is a simple queue, make one using a slice._".

"_Unlike garbage variables, leaked goroutines are not automatically collected, so it is important to make sure that goroutines terminate themselves when no longer needed._". The sub-section "[Always Clean Up Your Goroutines](#always-clean-up-your-goroutines)" below covers this.

The last paragraph of this page (Pg 207) discusses when to use unbuffered vs buffered channels.

Pg 207, **Reading, Writing, and Buffering** of "Learning Go an Idiomatic approach..." discusses using the built-in `cap` and `len` functions to find out how many elements the channel can hold, and how many elements are currently in the channel respectively.

Pg 209, **How Channels Behave** of "Learning Go an Idiomatic approach..." discusses the synchronisation of multiple writing channels using `sync.WaitGroup`, then refers to pg 220 ([Using WaitGroups](#using-waitgroups)). Pg 227, **8.4.1. Unbuffered Channels** of "The Go Programming Language" discusses using the "done" channel pattern. "Learning Go an Idiomatic approach..." also has a section on **The Done Channel Pattern**. "The Go Programming Language" also discusses and has examples for `sync.WaitGroup`.

Pg 226, **8.4.1. Unbuffered Channels** of "The Go Programming Language" mentions "_When a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the sending goroutine._"

## select

Pg 211 of "Learning Go an Idiomatic approach..." mentions using "The Done Channel Pattern", and https://go.dev/tour/concurrency/5 actually uses it (`quit <- 0`). Also seen on pg 227 of "The Go Programming Language" (`done <- struct{}{}`).

Pg 212 of "Learning Go an Idiomatic approach..." mentions: "_If you want to implement a nonblocking read or write on a channel, use a `select` with a `default`_". There is also a warning that states: "_Having a `default` case inside a `for`-`select` loop is almost always the wrong thing to do. It will be triggered every time through the loop when there’s nothing to read or write for any of the cases. This makes your for loop run constantly, which uses a great deal of CPU._". https://go.dev/tour/concurrency/5 states: "_A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready (Pg 210 of "Learning Go an Idiomatic approach..." also mentions this)._"... Unless of course there is a `default`. Pg 245 of "The Go Programming Language" states: "_the other communications do not happen._" within the `select` block.  
Pg 245 of "The Go Programming Language" states "_A select with no cases, `select{}`, waits forever._".

The following example is from pg 245 of "The Go Programming Language". It's behaviour is deterministic.

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		fmt.Printf("Before select, iteration %d, ch now holds %d elements\n", i, len(ch)) // Added for debugging.
		select {
	  	case x := <-ch:
		  	fmt.Println(x) // "0" "2" "4" "6" "8"
		  case ch <- i:
			  fmt.Printf("ch now holds the value %d\n", i) // Added for debugging.
		}
		fmt.Printf("After select,  iteration %d, ch now holds %d elements\n", i, len(ch)) // Added for debugging.
	}
}
```

"_Increasing the buffer size of the previous example makes its output nondeterministic, because when the buffer is neither full nor empty, the select statement figuratively tosses a coin._"

The following is the same example as above, but with the channel (`ch`) being unbuffered. **This deadlocks**, why?
Pg 207 of "Learning Go an Idiomatic approach..." says:
"_Every write to an open, unbuffered channel causes the writing goroutine to pause until another goroutine reads from the same channel. Likewise, a read from an open, unbuffered channel causes the reading goroutine to pause until another goroutine writes to the same channel. This means you cannot write to or read from an unbuffered channel without at least two concurrently running goroutines._"

```go
package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		fmt.Printf("Before select, iteration %d, ch now holds %d elements\n", i, len(ch)) // Added for debugging.
		select {
	  	case x := <-ch:
		  	fmt.Println(x) // "0" "2" "4" "6" "8"
		  case ch <- i:
			  fmt.Printf("ch now holds the value %d\n", i) // Added for debugging.
		}
		fmt.Printf("After select,  iteration %d, ch now holds %d elements\n", i, len(ch)) // Added for debugging.
	}
}
```

[Expecting a deterministic behavior using select and channels (#64)](https://100go.co/#expecting-a-deterministic-behavior-using-select-and-channels-64) of "100 Go Mistakes and How to Avoid Them" is mostly about buffered channels with a size greater than 1.

The following example is made up from bits and pieces from **8.7. Multiplexing with `select`** of "The Go Programming Language":

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {}

func main() {
	abort := make(chan struct{})
	fmt.Println("Commencing countdown. Press return to abort.")

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	ticker := time.NewTicker(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}
```


## Concurrency Practices and Patterns

### Goroutines, for Loops, and Varying Variables

Any time a goroutine depends on a variable in an outer scope whose value might change, you must pass the value into the goroutine.

I think the reason this example works is because the shadowing `v` in the `for range` loop's block is part of the goroutine's closure.

```go
for _, v := range a {
  v := v
    go func() {
      ch <- v * 2
  }()
}
```

This is easier to read and makes more initial sense:

```go
for _, v := range a {
  go func(val int) {
    ch <- val * 2
  }(v)
}
```

Go 1.22 [fixes this](https://go.dev/blog/loopvar-preview). "_For Go 1.22, we plan to change for loops to make these variables have per-iteration scope instead of per-loop scope._"

### Always Clean Up Your Goroutines

Whenever you launch a goroutine, you must make sure that it will eventually exit. Unlike variables, the Go runtime can’t detect that a goroutine will never be used again. If a goroutine doesn’t exit, the scheduler will still periodically give it time to do nothing, which slows down your program. This is called a goroutine leak.

```go
func countTo(max int) <-chan int {
  ch := make(chan int)
  go func() {
    for i := 0; i < max; i++ {
      ch <- i // Goroutine blocks forever waiting for a listener.
    }
    close(ch)
  }()
  return ch
}

func main() {
  for i := range countTo(10) {
    if i > 5 { // If we exit the loop early.
      break
    }
    fmt.Println(i)
  }
}
```

Pg 219, **8.1. Goroutines** of "The Go Programming Language" says that if there are goroutines still running when the `main` function returns, "_all goroutines are abruptly terminated_".

### The Done Channel Pattern

I thought the example provided was incorrect as the `case <-done:` didn't provide a `return`, but it doesn't need to. `select` will block until either `searcher` (function) or `done` (channel) provides a value, and because all but the fastest search's do not provide values from the `searcher`, the `done` channel is closed, so the zero value of the `done` channel is read by the `select`, and the goroutine finishes.

**8.9. Cancellation** of "The Go Programming Language" covers the same topic.

### Using a Cancel Function to Terminate a Goroutine

A function that contains a goroutine that returns it's goroutine's sending channel, along side a cancel function (closure containing code which closes a done channel).

### When to Use Buffered and Unbuffered Channels

I'm not sure this section answers this question very well, at least it seems possibly a little naive? This is what it says:

"_Buffered channels are useful when you know how many goroutines you have launched, want to limit the number of goroutines you will launch, or want to limit the amount of work that is queued up._"

See the [Channels](#channels) subsection for quite a few comments around this decision.

### Backpressure

If the number (concurrent requests) of tokens (a token being a `struct{}`) exceeds a specific number, the request is dropped and an error logged. I can see how this could keep resources within an acceptable/agreed limit.

### Turning Off a case in a select

If you're using a `select` statement, and one or more of it's cases reads from a channel that's been closed, that isn't being iterated on by a `for` `range` loop, then the `case` will continue to be successful (return the zero value) and waste time. Instead set the channel's variable to `nil`, this will stop the case from being run because a `nil` channel doesn't return a value.

### How to Time Out Code

Use a `select` statement with a `case` that reads a done channel (work has finished within time), and a `case` that reads a `Time` channel returned by `time.After(time.Duration)`.

Pg 247, **8.7. Multiplexing with `select`** of "The Go Programming Language" has an exercise for the reader to pretty much do the example just mentioned in Pg 220 of "Learning Go an Idiomatic approach...".

There is another timeout strategy on Pg 130 of **5.4.1. Error-Handling Strategies** of "The Go Programming Language".

### Using WaitGroups

Sometimes one goroutine needs to wait for multiple goroutines (or more specifically, the writing channels within the multiple goroutines) to complete their work, and their writing channels closed. If you are waiting for a single goroutine, you can use the done channel pattern that we saw earlier. But if you are waiting on several goroutines, you need to use a
WaitGroup.

Also covered in: <div id="looping-in-parallel"></div> <!--anchor linked to from ../TheGoProgrammingLanguage/ch8_8.5/README.md-->

* Pg 237-239, **8.5. Looping in Parallel** of "The Go Programming Language". The code for this can be found [here](https://github.com/adonovan/gopl.io/blob/1ae3ec64947b7a5331b186f1b1138fc98c0f1c06/ch8/thumbnail/thumbnail_test.go#L117-L146), as well as a copy of it [here](../TheGoProgrammingLanguage/ch8_8.5/thumbnail/thumbnail_test.go#L117-L146). I also used this pattern as a start for the Gitlab interview I did. You can see the pattern in the files: [bytes.go](https://github.com/binarymist/learning-go/blob/main/TheGoProgrammingLanguage/ch8_8.5/gitlab_interview/interview-task-161848681002-my-improvements/internal/dupe/bytes.go), [checksum.go](https://github.com/binarymist/learning-go/blob/main/TheGoProgrammingLanguage/ch8_8.5/gitlab_interview/interview-task-161848681002-my-improvements/internal/dupe/checksum.go), which the [find.go](https://github.com/binarymist/learning-go/blob/main/TheGoProgrammingLanguage/ch8_8.5/gitlab_interview/interview-task-161848681002-my-improvements/internal/dupe/find.go) file consumes
* Pg 250, **8.8. Example: Concurrent Directory Traversal** of "The Go Programming Language" also has an example
* Pg 274, **9.7. Example: Concurrent Non-Blocking Cache** of "The Go Programming Language" also has an example
* "100 Go Mistakes and How to Avoid Them" [Misusing sync.WaitGroup (#71)](https://100go.co/#misusing-syncwaitgroup-71)

"_A `sync.WaitGroup` doesn’t need to be initialized, just declared, as its zero value is_" where we start from.

Basically we just use the following statements in different places:

```go
var wg sync.WaitGroup
wg.Add(3) 
defer wg.Done() // Within a goroutine
defer wg.Done() // Within another goroutine
defer wg.Done() // Within another goroutine
wg.Wait()
```

We don't pass the `sync.WaitGroup` instance to each goroutine, because unless you pass a pointer to it, the copy is modified, which doesn't decrement the original. "_By using a closure to capture the `sync.WaitGroup`, we are assured that every goroutine is referring to the same instance._". 

"_While WaitGroups are handy, they shouldn’t be your first choice when coordinating goroutines. Use them only when you have something to clean up (like closing a channel they all write to) after all of your worker goroutines exit._"

The [golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) type builds on top of `WaitGroup` "_to create a set of goroutines that stop processing when one of them returns an error_". Also covered in "100 Go Mistakes and How to Avoid Them" [Not using errgroup (#73)](https://100go.co/#not-using-errgroup-73).

> "100 Go Mistakes and How to Avoid Them" [Forgetting about sync.Cond (#72)](https://100go.co/#forgetting-about-synccond-72) allows you to "_send repeated notifications to multiple goroutines_"

### Running Code Exactly Once

`sync.Once` allows for initialisation (or any) code to run exactly once. Like `sync.WaitGroup`, we declare a variable, no need to initialise it, because we want it to have its zero value.

"_Also like `sync.WaitGroup`, we must make sure not to make a copy of an instance of `sync.Once`, because each copy has its own state to indicate whether or not it has already been used. Declaring a `sync.Once` instance inside a function is usually the wrong thing to do, as a new instance will be created on every function call and there will be no memory of previous invocations._" It's usually best to declare the `sync.Once` instance at the package level.

Also covered in:

* Pg 268, **9.5. Lazy Initialization: `sync.Once`** of "The Go Programming Language"

### Putting Our Concurrent Tools Together

Covers the example of 3 web services being called, waiting for the result of the first 2 services, then sending them to the third service. The entire process must complete in 50 milliseconds.

## When to Use Mutexes Instead of Channels

A mutual exclusion limits the concurrent execution of some code or access to a shared (_critical section_) piece of data.

[Being puzzled about when to use channels or mutexes (#57)](https://100go.co/#being-puzzled-about-when-to-use-channels-or-mutexes-57) of "100 Go Mistakes and How to Avoid Them" also discusses this. "_Being aware of goroutine interactions can also be helpful when deciding between channels and mutexes. In general, parallel goroutines require synchronization and hence mutexes. Conversely, concurrent goroutines generally require coordination and orchestration and hence channels._"

[Effective Go: Concurrency - Share by communicating](https://go.dev/doc/effective_go#sharing) is along similar lines.

[Go by Example: Mutexes](https://gobyexample.com/mutexes)

I have an [example code](./ch10_ConcurrencyInGo/when_to_use_mutexes_instead_of_channels/) for this. The easiest way to see the differences is to put the files side by side.

## Atomics

"_Like sync.WaitGroup and sync.Once, mutexs must never be copied. If they are passed to a function or accessed as a field on a struct, it must be via a pointer. If a mutex is copied, its lock won’t be shared._"

"_Never try to access a variable from multiple goroutines unless you acquire a mutex for that variable first._"

`sync.Map` is only appropriate in situations:

* When you have a shared map where key/value pairs are inserted once and read many times
* When goroutines share the map, but don’t access each other’s keys and values

"_`sync.Map` uses interface{} as the type for its keys and values; the compiler cannot help you ensure that the right data types are used._"  
"_Given these limitations, in the rare situations where you need to share a map across multiple goroutines, use a built-in map protected by a `sync.RWMutex`._"

# Chapter 12. The Context

The Go wiki [CodeReviewComments#contexts](https://go.dev/wiki/CodeReviewComments#contexts) also touches on this.

[Misunderstanding Go contexts (#60)](https://100go.co/#misunderstanding-go-contexts-60) of "100 Go Mistakes and How to Avoid Them" also discusses this.  
"_In general, a function that users wait for should take a context, as doing so allows upstream callers to decide when calling this function should be aborted._"

# Chapter 13. Writing Tests

## The Basics of Testing

### Reporting Test Failures

The Go wiki [CodeReviewComments#useful-test-failures](https://go.dev/wiki/CodeReviewComments#useful-test-failures) has some thoughts around what test failures should look like.

## Finding Concurrency Problems with the Race Checker

A _data race_ is a type of _race condition_.

Pg 257, **9.1. Race Conditions** of "The Go Programming Language"

"100 Go Mistakes and How to Avoid Them" [Not understanding race problems (data races vs. race conditions and the Go memory model) (#58)](https://100go.co/#not-understanding-race-problems-data-races-vs-race-conditions-and-the-go-memory-model-58) 

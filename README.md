1. Install Go runtime > https://www.golang.com/dl > After installation run go command in Terminal to confirm installation.
2. Install VSCode.
3. Go extension in VSCode.
4. To run properly Go requires additional tools. Close and re-open VSCode > Create a new file and change language to Go > install required tools from prompt.

Questions:
1. How do we run the code in our project?
Go CLI commands
    1. go build = compiles a bunch of go source code files
    2. go run = Compiles and executes one or two files
    3. go fmt = Formats all the code in each file in the current directory
    4. go install = Compiles and installs a package
    5. go get = Downloads the raw source code of someone else's package
    6. go test = runs any tests associated with the current project

2. What does 'package main' mean?
Used to group together code with a similar purpose.
"main" - special name used for a package to tell Go that we want it to be turned into a file that can be executed.
Package == Project == Workspace 

Types of Packages in Go
    1. Executable = Generates a file that we can run
    2. Reusable = Code used as 'helpers'. Good place to put reusable logic

Notes
    1. package main = means we are making a executable package. Must have a func called main. main is special.
    2. package blahblah = defines a package that can be used as a dependency (helper code)

3. What does import "fmt" mean?
import statement used to gain access to other packages outside the current package.
fmt is short for format. fmt is package part of the standard library of Golang. Other standard packages are- math, debug, io, cryto, encoding etc.
Can import packages authored by others (reusable packages) - import "calculator"
www.golang.org/pkg

4. What is func?
This is a Function just like any other programming language.

5. How is the main.go file organized?
    1. Package decalration - package main
    2. Import other packages that we need - import "fmt"
    3. Declare fucntions, tell Golang to do things - func main() {}

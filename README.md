# Go Basics
---
Repository with basics about Go programming.

![Go Logo](go.png)

## Setup

1. Installing Go on Ubuntu
```shell
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt-get update
$ sudo apt-get install golang-go
```

For anothers S.O and distribution check the [main page](https://golang.org/).

2. Set GoPath for your workspace.
```
export GOPATH=/home/myuser/goProjects
```

I recommend you add it to your `.bashrc` or `.zshrc`.

## Run go project

Just type the following command
```
$go run filename.go
```

## Exercises

[00. Hello World](goexercises/00.helloworld.go)
[01. Vars Declarations](goexercises/01.varsdeclarations.go)
[02. Standard Input](goexercises/02.standardinput.go)
[03. More examples about Standard Input](goexercises/03.standardinput.func.go)
[04. Functions](goexercises/04.functions.go)
[05. Calculator](goexercises/05.calculator.go)
[06. Strings](goexercises/06.strings.go)
[07. Switch](goexercises/07.switch.go)
[08. Arrays](goexercises/08.arrays.go)
[09. Packages](goexercises/09.packages.go)
[10. Conditionals](goexercises/10.conditionals.go)
[11. Loops](goexercises/11.loops.go)
[12. Map](goexercises/12.map.go)
[13. Structs](goexercises/13.structs.go)
[14. Methods](goexercises/14.methods.go)
[15. Interfaces](goexercises/15.interfaces.go)
[16. Defer](goexercises/16.defer.go)
[17. Panic](goexercises/17.panic.go)
[18. Pointers](goexercises/18.pointers.go)
[19. Goroutines](goexercises/19.goroutines.go)
[20. Channels](goexercises/20.channels.go)
[21. More Channels](goexercises/21.morechannels.go)
[22. Server](goexercises/22.server.go)
[23. Api](goexercises/23.api.go)
[24. Files](goexercises/24.files.go)
[More examples](goexercises/main.go)



Links && Credits
---

- [Go Packages && Libraries](https://golang.org/pkg/)
- [Platzi Go Course](https://platzi.com/clases/go-basico/)

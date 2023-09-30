# GO 

## Run Command 
**To init the modules required fro GO this creates an mod file to resolve dependencies**

go mod init

**To Run the**

go run main.go

## Help 
go help

## Path
go env GOPATH

This  shows the actual path of the packages installed for GO-lang

## Grammer

; is handled by the lexer (Although it can be placed in all the line as terminator, lexer removes it)

## Type

**Basic Types:**
 String, Boolean, Integer, Float, Complex

**Advance Type:**
 Array, Slice, Maps, Structs, Pointers (etc.)

 ## Go Build
 GOOS="linux" go build

 ## GO-GC
 Its dose have a treshold value in the env variables if it reaches that GC will trigger

 There are two methods in main use case `new()` and ``make()``
 - New Provide a memore of non-zeroed storage with no init
 - make() gives you zeeroed store with init (Can store at allocation itself)

 ## Pointers
 Some times when you pass the variables to the method an copy of the variable would be passed which intends us to pass the exact value menore of it we use pointer so this ensures that the exact memore(Allocated for variable) is used.

Simple Pointer is pointing the memory, 
 Pointer is the reference to direct memory location

 if you need to get the value store in the memory address you need to use ***pointerName**


## Array

Array is not much used in the go lang

But Still you can
```
var fruitList [5]string

fruitList[0]="Orange
fruitList[1]="Orange
fruitList[3]="Orange

fmt.Println(len(fruitrList))

```

Although only 3 value is there the length will be still 5 since the size if 5 while creating the variable

```
var fruitList= [5]string{"Apple","Orange","Banana","Some","More"}

```

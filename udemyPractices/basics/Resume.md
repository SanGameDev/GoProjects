# This is my resume of this part of the course

### [01-1-variables](01-1-variables/main.go)

My scripts need the package at the start of the script
`package` and the name that you want to give it <ins>(**Note:** it doesn't need to 
have the same as the script name)</ins>

After the `package` u will have to use the `import` and the package inside "" to call 
it also for the external packages u have to use the github link like "github.com/userName/repoName" 

For the functions i have to use the word `func` follow by the name of this one and () {} 
<ins>(**Important:** if the name starts with cap its a public function and if its a lower one its a 
private one, ALSO in this language the { it cant go under the line that defines the function)</ins>
  
For the variables y can use the word `var` follow by the name of the variable and then the type of 
this one like this example `var myVariable string`

The types of vars that i can use are `int`, `unit`, `string`, `bool`
  
I can use the `&` in front of my var to see the direction in the memory that has

Go can detect the type of var that im using with `:=` so like in the next example `myVar := 12` 
with this Go knows that im using an `int`

I can use the `const` and this will take the value assigned as his type of var like the example 
`const myVar = "a23d"`, also i can set the const with a type of var from the start like `const myVar int = 65` 

My variables can have different amount of bites like `[8 16 32 64]` and this ones go after the 
type of var like `int16` 

The packages that i used in this class are `fmt`, `strconv`, `unsafe`:
  - I used `fmt` for printing in the terminal 
  - `strconv` for string conversion 
  - and `unsafe` for checking the Size of some vars

---

### [01-2-bytes](01-2-bytes/main.go)

Just some ASCII code things nothing to complicate

---

### [02-operatorsConditionals](02-operatorsConditionals/main.go)

First i saw the always known `>` `<` `<=` `>=` `==` after that the `||` `&&` `!` 
  
Then i learn that i can use `()` for the operators so i can set the order in witch the will be checked like maths

Then the `if` and `if else`(its weird that dont use the () )

For the switch in Go i can use something like this:
  ```Go
    switch {
	case number > 0:
		fmt.Println("positive")
	case number < 0:
		fmt.Println("negative")
	case number == 0:
		fmt.Println("zero")
	}
  ```
  were i use the operators in the case (i also can use it like i know)

---

### [03-arrays](03-arrays/main.go)

I can get the length of an array using `len()`
  
I can declare slices, which are dynamically-sized arrays `var slice1 []int`

I can add elements to the slice using the `append()`
  
I can create a slice from an array using slicing syntax `mySlice := myArrayVar[2:4]`
  
Modifying a slice also modifies the underlying array
  
I can create slices with make(), specifying the length `slice := make([]int, 3)`

---

### [04-maps](04-maps/main.go)

So for what i learn about this ones is that they are like an array of a key and a value and if u want to access a value u need the key and the keys cant repeat

I can create this maps using the `make(map[key type]value type)` like `map1 := make(map[int]string)`

For creating a key/value pair in a map i use the `mapVarName[key value] = value`

For accessing the values i use the `mapVarName[key]`

I can delete a key/value using `delete(mapVarName, key value)`

I can check if a key exist using `value, bool if exist := mapVarName[key value]`

I can initialize a map with values using `mapVarName := map[key type]value type{key value: value, ...}`

---

### [05-for](05-for/main.go)

In Go there is no another loop than for

I can use it like always but without the ()

I can use it like a wile for example `for value comparator value` `for varName < 10`

I can use a range like `for i := range arrayVar` i can use this in arrays and maps

---

### [06-functions](06-functions/main.go)

Wen creating a package the scrips that has the same `package nameOfPackage` the 
private functions can be used in any script that is part of the same package (but not in the script that is using the package) 

Wen creating a function that has to receive and return data the structure has to be 
the next 
```go
func nameOfFunction(nameOfData typeOfData)valueToReturn
``` 
if return more than one value they also go inside some `()` <ins>(**Note:** if the values that im sending or returning i can set first all the names of the vars and then tell him what type they are like `x, y int`)</ins>

Also the functions can handle errors like `func name (x, y int) (int, error)`

Also i can create functions that receive a different amounts of the same type of value 
using the `...` after the name of the value like `valueName ...valueType`

Also y can use only `return` without the name of the values if the function only 
returns one value

### [07-structsInterfaces](07-structsInterfaces/main.go)

The structs are containers that can hold different types of values inside 
and then i can use that struct like a type of value

For how to make one its like the next example:
  ```go
  type nameOfStruct struct{
    value1 typeOfValue 
    value2 typeOfValue
    ...etc
  }
  ```

The type interface are like the structs but they hold methods and dose methods are 
empty so the structs work like the next example:

```go
type Damageable interface {
    TakeDamage(amount int)
}

type Player struct {
    Name string
    Health int
}

func (p Player) TakeDamage(amount int) {
    p.Health -= amount
    fmt.Println(p.Name, "took", amount, "damage. Remaining health:", p.Health)
}

func main() {
    var d Damageable

    d = Player{Name: "Hero", Health: 100}
    d.TakeDamage(10)
}
```

---

### [08-pointers](08-pointers/main.go)

I can set a pointer by using the `*` like the next example `var varName *varType`

I can get the address of a var by using the `&` like `&varName`

I can get a reference of the pointer by using the `*pointerName`

Also i can change the value of a var by using the pointer like `*pointerName = newVar`

Also i can use pointers to methods like `methodName *Struct`

---

### [09-errors](09-errors/main.go)

I can make a new error message using `errors.New()` like the next example 
`errorVarName = errors.New("message of error")` <ins>***Note:***(the firs letter 
cant be in caps)</ins>

I can use a like print but for errors with `fmt.Errorf()`

I can use the `defer func` to make that a functions happens in a case 
of an error (this thing will be always run after the function ends, ALWAYS)

I can recover from a `panic` using a `recover`

---

### [10-interPkgs](10-interPkgs)

The package time is used for working with dates and time

The package log is used for logging messages

The package net is used for networking, including HTTP requests

The package os is used for interacting with the operating system, 
including file operations and environment variables.

The package reflect is used for reflection, which allows a program to 
inspect and manipulate its own structure and behavior at runtime

The package strconv is used for converting strings to other data types and 
vice versa

---

### [11-extrPkgs](11-extrPkgs)

I learn that for getting external packages have to get them by using the 
`go get "github.com/userName/packageName"`

That line will download you the last version of the package, if you want a
specific one you can use the `@` after the line above and then the version u want

---

### [12-generics](12-generics/main.go)

I can make generic functions using type parameters like this example:
```go
func functionName[genericName int | int32 | int64 | float64 | float32](varName genericName)
```

I also can make an interface that saves all the generic data and the i can call that interface
and use it for the generic like:
```go
type genericName interface{
	int | int32 | int64 | float64 | float32
}
```

I can use the `any` type to accept any type of value `func anyType[N any](v1, v2 N)`

I can use the `comparable` type constraint to accept only comparable types 
`func comparableType[N comparable](v1, v2 N)`

I can use the `constraints.Ordered` type constraint to accept only ordered types 
`func orderedValues[N constraints.Ordered](v1, v2 N)`

I can define custom generic `types` `type CustomSlice[V int | string] []V`

I can define generic `structs` `type MyGenericStruct[D any] struct { StrValue string; Data D }`

I can define methods for specific types in a generic struct `func (d MyFirstData) PrintOne()`

---

### [13-roroutines](13-roroutines/main.go)

I can create routines using the `go` before the function call like: `go MyFunction()`

The `time.Sleep()` can be used to make the function wait

I can make channels using the `make`

I can send and receive values to a channel using the `<-` 

The channels can be closed using the `close` and they have to be closed

---

### [14-context](14-context/main.go)

I can create a background context using `context.Background()`

I can add values to a context using `context.WithValue()` like this example: `ctx = context.WithValue(ctx, "my-key", "sarasa")`

I can create a context with a timeout using `context.WithTimeout()` like this example: `ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)`

I can defer the cancellation of a context using `defer cancel()`

I can use the `select` statement to handle multiple channels and context cancellation `select { case <-ctx.Done(): ... }`

I can retrieve values from a context using `ctx.Value("my-key")`

I can handle context cancellation in a goroutine using the `case <-ctx.Done():`

---

### [15-binaryOperators](15-binaryOperators/main.go)

just some binary things
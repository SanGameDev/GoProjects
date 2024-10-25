# This is my resume of this part of the course

## Index:
- 01-1-variables
- 01-2-bytes
- 02-operatorsConditionals
- 03-arrays
- 04-maps
- 05-for
- 06-functions

---

### [01-1-variables](01-1-variables/main.go)

- My scripts need the package at the start of the script
`package` and the name that you want to give it <ins>(**Note:** it doesn't need to have the same as the script name)</ins>

- After the `package` u will have to use the `import` and the package inside "" to call it also for the external packages u have to use the github link like "github.com/userName/repoName" 

- For the functions i have to use the word `func` follow by the name of this one and () {} <ins>(**Important:** if the name starts with cap its a public function and if its a lower one its a private one, ALSO in this language the { it cant go under the line that defines the function)</ins>
  
- For the variables y can use the word `var` follow by the name of the variable and then the type of this one like this example `var myVariable string`

- The types of vars that i can use are `int`, `unit`, `string`, `bool`
  
- I can use the `&` in front of my var to see the direction in the memory that has

- Go can detect the type of var that im using with `:=` so like in the next example `myVar := 12` with this Go knows that im using an `int`

- I can use the `const` and this will take the value assigned as his type of var like the example `const myVar = "a23d"`, also i can set the const with a type of var from the start like `const myVar int = 65` 

- My variables can have different amount of bites like `[8 16 32 64]` and this ones go after the type of var like `int16` 

- The packages that i used in this class are `fmt`, `strconv`, `unsafe`:
  - I used `fmt` for printing in the terminal 
  - `strconv` for string conversion 
  - and `unsafe` for checking the Size of some vars

---

### [01-2-bytes](01-2-bytes/main.go)

- Just some ASCII code things nothing to complicate

---

### [02-operatorsConditionals](02-operatorsConditionals/main.go)

- First i saw the always known `>` `<` `<=` `>=` `==` after that the `||` `&&` `!` 
  
- Then i learn that i can use `()` for the operators so i can set the order in witch the will be checked like maths

- Then the `if` and `if else`(its weird that dont use the () )

- For the switch in Go i can use something like this:
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

- I can get the length of an array using `len()`
  
- I can declare slices, which are dynamically-sized arrays `var slice1 []int`

- I can add elements to the slice using the `append()`
  
- I can create a slice from an array using slicing syntax `mySlice := myArrayVar[2:4]`
  
- Modifying a slice also modifies the underlying array
  
- I can create slices with make(), specifying the length `slice := make([]int, 3)`

---

### [04-maps](04-maps/main.go)

- So for what i learn about this ones is that they are like an array of a key and a value and if u want to access a value u need the key and the keys cant repeat

- I can create this maps using the `make(map[key type]value type)` like `map1 := make(map[int]string)`

- For creating a key/value pair in a map i use the `mapVarName[key value] = value`

- For accessing the values i use the `mapVarName[key]`

- I can delete a key/value using `delete(mapVarName, key value)`

- I can check if a key exist using `value, bool if exist := mapVarName[key value]`

- I can initialize a map with values using `mapVarName := map[key type]value type{key value: value, ...}`

---

### [05-for](05-for/main.go)

- In Go there is no another loop than for

- I can use it like always but without the ()

- I can use it like a wile for example `for value comparator value` `for varName < 10`

- I can use a range like `for i := range arrayVar` i can use this in arrays and maps

---

### [06-functions](06-functions/main.go)


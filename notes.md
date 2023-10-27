### what is module in go ?
    - your project 

### how we can create a module ?
    - go mod init module_name {can be url , project_name}

### structure of go code can be like this 
    - workspace => module => package

### variables scopes can be 
    - global i can not use the short cut var in the global   , functional inside func, block scope inside  {}    

### what is a package ?
    - group of files in the same folder like {utils, middlware}
    varibles , functions in the same package can be used without importing it if them in the same folder 

### how to know if the function is exported 
    - first letter should be capitel like : fmt.Println()
### init() exe first before main() 
    - can be used to perform some initialization tasks, such as setting up global variables, registering handlers, or validating configurations.

### if you have a function that will modify an argument these argument must be pointer ?
``` go
    func birthday(pointerAge *int) {
	fmt.Printf("The Pointer is %v and the value is %v \n", pointerAge, *pointerAge)
	*pointerAge++}

```

### panic() it interupt the code if condition is true  , defer it called at the end of the code 

###  Type Definitions :

- Type aliases: create a new type that is an alias for an existing type
    ```go
    type NewType = ExistingType
    ```
    
     ```go
    type json = map[string]string
    ```
    ```go
    map in golang : map[typeOfKey]typeOfValue
    ```
    - Creating a New Type Based on an Existing Type
 
 we can create a new type based on an existing type using the `type` declaration. This allows you to add additional behavior or constraints to an existing type

```go
type Fahrenheit float64

// Define a method for the new type
func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g°F", f)
}

func main() {
    // Create a variable of the new type
    temperature := Fahrenheit(32.0)

```
## Methods in Go

methods are functions associated with a `specific type`. They allow you to add behaviors or actions to user-defined types (structs) or types from the standard library. Methods are defined with a receiver, which is a parameter that specifies the type on which the method operates.

### Defining a Method

```go
func (receiver ReceiverType) MethodName(parameters) ReturnType {
    // Method implementation
}
```
- `receiver` is the parameter specifying the type on which the method is defined, similar to `this` or `self` in other languages.
- `ReceiverType` is the type the method is associated with.
- `MethodName` is the method's name.
- `parameters` are any additional parameters the method may take.
- `ReturnType` is the type the method returns.

```go
package main

import "fmt"

type Circle struct {
    Radius float64
}

// Method for the Circle type to calculate the area
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    circle := Circle{Radius: 5.0}
    area := circle.Area()
    fmt.Printf("Circle area: %.2f\n", area)
}

```

We define a `Circle` struct with a Radius field.
We define a method, `Area()`, associated with the `Circle type`, which calculates the area of the circle.
In the main function, we create a `Circl`e instance and call the Area method to calculate and print the circle's area.
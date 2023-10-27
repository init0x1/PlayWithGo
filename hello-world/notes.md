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

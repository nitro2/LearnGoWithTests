# LearnGoWithTests
This project captures the learning Go progress by studying the course that https://quii.gitbook.io/learn-go-with-tests

This is useful course that I've recently discovered.

Happy learning!

# Notes

## Array and Slice 
Array and Slice are DIFFERENT
Array is kind of immutable data structure while slice is mutable.

An assignment operation does a copy of every element in array to new location
Eg:
```
a := [3]int{1,2,3} // Define an array
b := a  // Copy array to another array
c := a[:] // Define a slice and reference it to `a`
a[0] = 0
fmt.Println(a)       // Prints "[0 2 3]"
fmt.Println(b)       // Prints "[1 2 3]"
fmt.Println(c)       // Prints "[0 2 3]"

fmt.Printf("Address of &a: %p \n"+              //Address of &a: 0xc000014018 
           "Address of &b: %p \n"+              //Address of &b: 0xc000014030
           "Address of &c: %p \n"+              //Address of &c: 0xc00000c030 
           "Address of c: %p \n", &a,&b,&c,c)   //Address of c: 0xc000014018
```
Reference: https://stackoverflow.com/a/21722697/1177962

# LearnGoWithTests
This project captures the learning Go progress by studying the course that https://quii.gitbook.io/learn-go-with-tests

This is useful course that I've recently discovered.

Happy learning!

# Notes

## Array and Slice 
Array and Slice are DIFFERENT.

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

My current understanding is: `Slice is a container-like while array is a object`

Example: https://play.golang.org/p/bTrRmYfNYCp

## Map
Map is also a container-like, so we can freely change the content without passing the pointer

```
type Dictionary map[string]string

func (d Dictionary) Add(key, value string) error {
	d[key] = value
	return nil
}
...
d := Dictionary{}
d.Add("D","Days")

// d -> "D":"Days"
```

If we use pointer like this:
```
type Dictionary map[string]string

func (d *Dictionary) Add(key, value string) error {
	d[key] = value
	return nil
}
...
d := Dictionary{}
d.Add("D","Days")
```
... then we get error: `./dictionary.go:19: invalid operation: d[key] (type *Dictionary does not support indexing)`

> A map value is a pointer to a runtime.hmap structure.
> 
> So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

Initilize a map:
```
var dictionary = map[string]string{}
// OR
var dictionary = make(map[string]string)
```


## Concurrency
The go routines do not have their own copy of sharing variables. So they have a big chance to fetch the only last values of sharing variables 

```golang
numbers := []int{1, 2, 3, 4}
for _, n := range numbers {
	go func() {
		fmt.Println(n)
	}()
}

// Result:
// 4
// 4
// 4
// 4
```
In above example, each of our go routines have a reference to the `n` variable. Therefore, they all read `n` as `4`.

By passing argument into the go routines, we can make sure the routines use correct values:
```golang
numbers := []int{1, 2, 3, 4}
for _, n := range numbers {
	go func(x int) {
		fmt.Println(x)
	}(n)
}

// Result:
// 4
// 1
// 2
// 3
// Result order is randomly, depend on routines.
```


## Select
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/select

Consider following code:

```
func Racer(a, b string) (winner string) {
    select {
    case <-ping(a):
        return a
    case <-ping(b):
        return b
    }
}

func ping(url string) chan struct{} {
    ch := make(chan struct{})
    go func() {
        http.Get(url)
        close(ch)
    }()
    return ch
}
```

Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?

### Always make channels
Notice how we have to use make when creating a channel; rather than say var ch chan struct{}. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
For channels the zero value is nil and if you try and send to it with <- **it will block forever because you cannot send to nil channels**

AGAIN: NEVER declare `var ch2 chan bool` because `ch2 <- true` will block forever


## Sync

The book advices us use Mutex as a member of struct rather than embedded it

**SHOULD**:
```
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
```

**SHOULD NOT**:
```
type Counter struct {
    sync.Mutex
    value int
}

func (c *Counter) Inc() {
    c.Lock()
    defer c.Unlock()
    c.value++
}
```

It's harmful because the API can be accidently called outside and drives the flow wrongly.
As my understanding, this is the only risk.


- Use channels when passing ownership of data
- Use mutexes for managing state

Use `go vet` to alert some subtle bugs!!!

## Context

https://blog.golang.org/context

> `context.Context` is an `immutable object`, "extending" it with a key-value pair is only possible by making a copy of it and adding the new key-value to the copy

> A Context is safe for simultaneous use by multiple goroutines. Code can pass a single Context to any number of goroutines and cancel that Context to signal all of them.


> At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly.

It's better to manage functions with context.
```
- WithCancel() -> Can use it to cancel go routine in background to prevent memory leak.
- WithDeadline() -> Create a deadline cleanup routine to prevent function from unexpected blocking call. Use absolute time.
- WithTimeout() -> Similiar to WithDeadline(). Use relative time.
- WithValue() -> We can start many contexts with different input value. WithValue() needs a pair of key and val interface{}. 
- Background()
```

### WithCancel()
### WithDeadline() vs WithTimeout()
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

Eg:
```
const shortDuration = 1 * time.Millisecond
d := time.Now().Add(shortDuration)
ctx, cancel := context.WithDeadline(context.Background(), d)
defer cancel()
```

```
const shortDuration = 1 * time.Millisecond
ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
defer cancel()
```

### WithValue()

`func WithValue(parent Context, key, val interface{}) Context`
> WithValue(): The provided key must be comparable and **should not be of type string** or **any other built-in type** to avoid collisions between packages using context. 

- https://pkg.go.dev/context#example-WithValue
- https://stackoverflow.com/questions/40379960/golang-context-withvalue-how-to-add-several-key-value-pairs

> Use context values only for request-scoped data that transits processes and API boundaries, not for passing optional parameters to functions. ??? (Not much understand the statement now)



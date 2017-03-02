package main

import (
  "fmt"

  "github.com/johnsiilver/classes/go/examples/stuff"
)

func main() {
  fmt.Printf("%d\n", stuff.Add(2,2))
}


x := 7
x = 3

x := 7
x := 3


// Declaring variables with the var keyword.
var x int
x = 7

var y string = "Hello World"

var z int = 7
var a = 7

// Declaring the same variables using :=
x := 7
y := "Hello World"
x = 12   // Why not := ???

// Declaring using := for a specific numeric type.
x := int64(7)

// Declaring multiple variables of the same type.
var a, b, c int = 3, 2, 1

// Declaring different variables, but using only one "var" keyword.
var (
  x float64
  y string = "Hello World"
  x uint32
)


var myArray [2]string

myArray[0] = "Hello"
myArray[1] = "World"

fmt.Println(myArray[0])
fmt.Println(myArray[1])

myArray[2] = "Won't work"



Hello
World


mySlice := make([]string, 1, 2)

mySlice[0] = "Hello"

mySlice = append(mySlice, "Again")

mySlice := make([]string, 0)
mySlice = append(mySlice, "apple")
mySlice = append(mySlice, "orange")
mySlice = append(mySlice, "pear")

subSlice = mySlice[0:2]

subSlice[1] = "mango"

fmt.Println(subSlice[2])
fmt.Println(mySlice[2])


cherry
cherry



subSlice[2]

subSlice = append(subSlice, "cherry")


mySlice := make([]string, 2)

mySlice = append(mySlice, "hello", "world")


subSlice = mySlice[0:2]


// Creates a map with keys that are strings, values that are int.
employeeID := make(map[string]int)

// Assign key "john" to value 0.
employeeID["john"] = 0

// Assign key "bill" to value 1.
employeeID["bill"] = 1

// Print out Bill's employee id.
fmt.Println(employeeID["bill"])

// Printing out a non-existing key will provide the value's Zero Value.
fmt.Println(employeeID["kate"])


// Composite literal declaring an empty slice of strings.
mySlice := []string{}

// Composite literal declaring a slice of length 3 and capacity 3 storing
// the value "apple" in index 0, "orange" in index 1, and "pear" in index 2.
mySlice := []string{"apple", "orange", "pear"}

// Composite literal declaring an empty map.
myMap := map[string]string{}

// Composite literal declaring a map with keys and associated values.
myMap := map[string]string{
  "key0": "value0",
  "key1": "value1",
  "key2": "value2",
}

for x := 0; x < 10; x++ {
  fmt.Println("hello")
}

x := 0
for x % 2 {
  fmt.Println("hello")
  x++
}

for {
  fmt.Println("hello")
}

slice := []string{"a", "b", "c"}
for index, value := range slice {
  fmt.Printf("%d: %s\n", index, value)
}

0: a
1: b
2: c

m := map[string]string{
  "a": "apple",
  "b": "banana",
}
for key, value := range m{
  fmt.Printf("%s: %s\n", key, value)
}

b: banana
a: apple


if x < 10 {
  fmt.Println("x is less than 10")
}

if x := y + z; x < 10 {
  fmt.Println("x is less than 10")
}

if x < 10 {
  fmt.Println("x is < 10")
}else if x == 11 {
  fmt.Println("x is 11")
}else {
  fmt.Println("x is > 11")
}

switch [value] {
case [match]:
  [statement]
case [match], [match]:
  [statement]
  [statement]
default:
  [statement]
}

switch x {
case 10:
  fmt.Println("x == 10")
case 20:
  fmt.Println("x == 20")
case 30, 40:
  fmt.Println("x == 30 or 40")
default:
  fmt.Println("don't know what x is")
}

switch [init statement]; [conditional] {
case [match]:
  [statement]
case [match], [match]:
  [statement]
  [statement]
default:
  [statement]
}



switch x := y % 2; x {
case 0:
	fmt.Println("x == 0")
case 1:
	fmt.Println("x == 1")
}

switch {
case [conditional]:
  [statement]
case [conditional]:
  [statement]
default:
  [statement]
}


switch {
case x < 10:
  fmt.Println("x < 10")
case x > 10:
  fmt.Println("x > 10")
}

func <function name>([var name] [var type], [var name] [var type], …) ([return value], [return value], …) {}

func add(x int, y int) int {
  return x + y
}

func addMultiply(x, y int) (int, int) {
  return x+y, x*y
}

func main() {
  a, m := addMultiply(3, 3)
  fmt.Printf("3 + 3 is %d\n", a)
  fmt.Println("3 * 3 is %d\n", m)
}


func addMultiply(x, y int) (add, multiply int) {
	return x+y, x*y
}

package stuff

import "fmt"

func PrintHello() {
  fmt.Println("Hello")
}

func printWorld() {
  fmt.Println("World")
}

func PrintHelloWorld() {
  PrintHello()
  // printWorld is private, but accessible because
  // it is in the same package.
  printWorld()
}


package main

import ".../stuff" // ... just means some path.

func main() {
  // This will work.
  stuff.PrintHello()

  // This will work.
  stuff.PrintHelloWorld()

  // Will get a compile error.
  stuff.printWorld()
}



package blah

var (
  // Multiplier can be accessed by outside packages by referencing
  // blah.Multiplier
  Multiplier = 3

  // divider can only be accessed within the package.
  divider = 2
)

func IncreaseByMultiplier(x int) int {
  return x * Multiplier
}

func DivideByTwo(x int) int {
  return x/divider
}

package main

import ".../blah" // ... stands for a directory path.

func main() {
  blah.Multiplier = 10

  // blah.divider = 2 // Won't work, divider isn't public.

  fmt.Println(blah.IncreaseByMultiplier(4))
}

func someFunc() {
  // str can be accessed by anything inside someFunc().
  // It cannot be accessed outside someFunc().
  str := "Hello World"
  fmt.Println(str)

  // i is accessible only inside the for loop.
  // It cannot be accessed outside the for loop.
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }
  // fmt.Println(i) - uncommenting this will be a compile error.

  // If you need for a loop variable to exist outside the loop, do this:
  var x int
  for x = 0; x < 10; x++ {
    fmt.Println(x)
  }
  // This works because x is declared outside the for loop.
  fmt.Println(x)

  // We create a variable z that is a random integer
  // using the standard library "math/rand".
  if z := rand.Int(); z == 0 {
    fmt.Printf("we randonly got 0 back")
  }else if z < math.MaxInt32 {
    // We can see z here because we are still within the entire if/else if/else.
    fmt.Printf("our random number %d fits in an int32", z)
  }else{
    // Same as above.
    fmt.Printf("our random number %d fits in an int64", z)
  }
  // fmt.Println(z) - uncommenting this will be a compile error.
}


// EmployeeRecord holds information about a company's employee.
type EmployeeRecord struct {
  // FirstName is the first name of the employee.
  FirstName string
  // Lastname is the last name of the employee.
  LastName string
  // ID is the employee's unique employee ID.
  ID int
}

func (e EmployeeRecord) String() string {
  // Sprintf is like Printf except that instead of writing to the screen,
  // it returns a string.
  return fmt.Sprintf("First: %s;Last: %s;ID: %d", e.FirstName, e.LastName, e.ID)
}



rec := EmployeeRecord{
  FirstName: "John",
  LastName: "Doe",
  ID: 0,
}

fmt.Println(rec.String())


var x int32 := 3

func changeInt32(x int32) {
  x = 10
  fmt.Printf("Inside changeInt32: %d\n", x)
}

func main() {
  var x int32 = 100
  changeInt32(x)
  fmt.Printf("After changeInt32: %d\n", x)
}

var x int32 = 100

intptr := new(int)

// This will print 0, the Zero Value stored at the place in memory
// intptr points to.
fmt.Println(*intptr)

// This changes the value stored in intptr to 25.
*intptr = 25

// This will print 25.
fmt.Println(*intptr)

// This will print the memory address that intptr stores.
fmt.Println(intptr)

// intptr  = 10 - Uncommenting this will prevent this from compiling.



func changeInt32(x *int32) {
  *x = 10
  fmt.Printf("Inside changeInt32: %d\n", *x)
}

func main() {
  var x int32 = 100
  changeInt32(&x)
  fmt.Printf("After changeInt32: %d\n", x)
}


func changeSlice(s []string) {
  s[0] = "blueberry"
  s = append(s, "watermelon")
}

func main() {
  s := []string{"apples"}
  changeSlice(s)
  fmt.Printf("%#v\n", s)
}


func changeSlice(s *[]string) {
	// The (*s) says to derference s.
	// (*s)[0] says dereference the pointer, then access element 0.
	// If you did *s[0], it would try to dereference
	// the value stored in s[0].
	(*s)[0] = "blueberry"
	*s = append(*s, "watermelon")
}

func returnSlice(s []string) []string {
	s[0] = "oranges"
	return append(s, "bananas")
}

func main() {
	s := []string{"apples"}

	changeSlice(&s)
	fmt.Printf("%#v\n", s)

	s = returnSlice(s)
	fmt.Printf("%#v\n", s)
}

type someStruct struct{
  field string
}
func (s *someStruct) change() {
  s.field = "hello"
}

func main() {
  s := &someStruct{field: "world"}
  s.change()
  fmt.Println(s.field)
}

<value> := map[<key>]

<value>, <found bool> := map[<key>]

m := map[string]int{"apples": 1}

// Does "apples" exist as a key in "m".  If so,
// return the value stored and boolean true.
if v, ok := m["apples"]; ok {
  fmt.Printf("Found key %q with value %d", "apples", v)
}else{
  fmt.Printf ("Key %q not found", "apples")
}


// foodsIlike is a map of foods like.  I'm very selective.
var foodsILike = map[string]bool{
  "bbq": true,
  "milkshakes": true,
}

// willEat determines if I will eat a meal that contains various foods.
// I'm picky, so if I don't like one food, I don't eat the meal.
func willEat(meal []string) bool {
  for _, f := range meal {
    if !foodILike[f]{
      return false
    }
  }
  return true
}


func Sum(nums ...int) int {
  s := 0
  for _, n := range nums {
    s +=n
  }
  return s
}

func Sum(nums []int) int {
  s := 0
  for _, n := range nums {
    s +=n
  }
  return s
}

x := Sum(1, 5, 7, 9)

x := Sum([]int{1, 5, 7, 9})

fruits := []string{"mango", "pears"}
californiaFruits := []string{"oranges", "apples"}

// Lets append californiaFruits onto fruits.
fruits = append(fruits, californiaFruits...)

// fruits = append(fruits, californiaFruits) - will not work

// PrintFile prints the content of file fn to the screen.
func PrintFile(fn string) error {
  content, err := ioutil.ReadFile(fn)
  if err != nil {
    // We could just return the error, but this allows us to add more
    // context to the error returned by ioutil.ReadFile().
    return fmt.Errorf("could not read file %s: %s", fn, err)
  }

  for _, line := range contents {
    fmt.Println(line)
  }
  return nil
}

func main() {
  files := []string{"exists", "does not exist"}

  for _, fn := range files {
    if err := PrintFile(fn); err != nil {
      fmt.Println("Error: %s", err)
    }
  }
}


func helloWorld() {
  func(str string){
    fmt.Println(str)
  }("hello world")  // Notice the ("hello world") here?  This is where we call the function.
}


printer := func(str string) {
  fmt.Println("str")
}

printer("The sections quotes don't seem right")


var addition = func(x, y int) int{
  return x + y
}

z := addition(2, 2)
fmt.Println(x)


func printing() {
  defer fmt.Println("exiting")
  defer fmt.Println("I must be going")

  fmt.Println("Hello,")
}

// filePrinter prints a file to the screen if it doesn't contain TOP SECRET in
// the file.  This is an example, its is not the most efficient way to do this.
func filePrinter(filePath string) error {
        f, err := os.Open(filePath)
        if err != nil {
                return err
        }
        defer f.Close()

        r := bufio.NewReader(f)
        output := ""
        for {
                out, err := r.ReadString('\n')
                if strings.Contains(out, "TOP SECRET") {
                        return fmt.Errorf("Top Secret file"))
                }
                output += out
                if err != nil {
                        break
                }
        }
        fmt.Println(output)
        return nil
}

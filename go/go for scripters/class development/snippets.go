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
fmt.Println(z)


func printing() {
  defer fmt.Println("exiting")
  defer fmt.Println("I must be going")

  fmt.Println("Hello,")
}

// filePrinter prints a file to the screen if it doesn't
// contain TOP SECRET in the file.  This is an example,
// its is not the most efficient way to do this.
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

func main() {
  if !secureConnection() {
    panic("Anyone could be listening!")
  }
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Printf("Recovered from a panic, panic was: %q\n", r)
    }
  }()
  panic("hello")
}





func filePrinter(filePath string) error {
        f, err := os.Open(filePath)
        if err != nil {
                return err
        }

        r := bufio.NewReader(f)
        output := ""
        for {
                out, err := r.ReadString('\n')
                if strings.Contains(out, "TOP SECRET") {
                        f.Close()  // Now I have to close it here.
                        return fmt.Errorf("Top Secret file"))
                }
                output += out
                if err != nil {
                        break
                }
        }
        fmt.Println(output)
        f.Close()  // And I have to here as well.
        return nil
}


var sample interface{Add(x, y int) int}

type notRight struct{}

func (n notRight)Hello() {
  fmt.Println("hello")
}

// This won't work, because 3 doesn't have any methods on it.
// sample = 3

// This won't work, because notRight does not have the right method.
// sample = notRight{}

type math struct {}

func (m math) Add(x, y int) int {
  return x+y
}

func (m math) Subtract(x, y int) int {
  return x-y
}

// This will work, because it does have the right method.
sample = math{}

// I can call this, because it is defined on the interface.
fmt.Println(sample.Add(3, 2))

// This will not work, because the interface does not define this method.
// fmt.Println(sample.Subtract(3, 2))


type Adder interface {
  Add(x, y int) int
}

var sample Adder


type MathStuff interface {
  Add(x, y int) int
  Multiply(x, y int) int
}

type MyMath1 struct {}

func (m MyMath1) Add(x, y int) int {
  return x+y
}

func (m MyMath1) Multiply(x, y int) int {
  return x * y
}

type MyMath2 struct {}

func (m MyMath2) Add(x, y int) int {
  return x+y
}

var m MathStuff = MyMath1{}

var m MathStuff = MyMath2

// This is Go's built in error type.
type error interface {
  Error() string
}


type CustomError int

func (c CustomError) Error() string {
  return fmt.Errorf("had a CustomError with error code %d", c)
}

func blah() error {
  return CustomError(1)
}

func main() {
  for i := 0; i < 10; i++ {
    go fmt.Println("hello world", i)
  }

  // We haven't talked about select yet.  Here it is used to keep the program
  // from ending.  An empty select blocks forever.  Without this, we might
  // not see anything print, because then main() function might end before
  // any of the goroutines are able to print.
  // If you try this, everything will print, but the program will crash.
  // We will talk about better ways to handle this later.
  select{}
}


ch := make(chan string, 1)

ch <- "hello"

x := <-ch

for item := range ch {
  fmt.Println(item)
}

close(ch)


// printer reads ints off of input and prints it to the screen.  When the
// input channel closes and all data has been read off of input, the exit
// channel will be closed to signal that printer is done processing.
func printer(input chan int, exit chan bool) {
  defer close(exit)
  for i := range input {
    fmt.Println(i)
  }
}

func main() {
  input := make(chan int, 10)
  exit := make(chan boo)

  // Start our printer in a goroutine.
  go printer(input, exit)

  // Spin off 100 goroutines that feed our printer.
  for i := 0; i < 100; i++ {
    go func(i int){
      input <-i
    }(i)
  }

  // Let the print know that we are no longer going to send data.
  close(input)

  // Wait for the printer to exit.
  <-exit
}

func main() {
  wg := &sync.WaitGroup{}

  // Sping off 1000 goroutines, each one printing a number.
  for i := 0; i < 1000; i++ {
    wg.Add(1)  // Add 1 to the counter.  You MUST do this before the func().
    go func(i int) {
      defer wg.Done() // Remove 1 from the counter after the func() is finished.
      fmt.Println(i)
    }(i)
  }

  wg.Wait()  // Wait unil all of the goroutines are finished.
}


var counter = 0

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	counter++
}

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(wg, mu)
	}

	wg.Wait()
	fmt.Println(counter)
}


// printer prints a number divded by 2 if it comes on the half channel.
// It prints a number multiplied by 2 if it comes on the double channel.
// It terminates if it receives anything on exit.
func printer(half, double chan int,
             wg *sync.WaitGroup, exit chan bool) {
  for {
    select {
    case i := <-half:
      fmt.Printf("%d/2 = %d\n", i, i/2)
    case i := <-double:
      fmt.Printf("%d*2 = %d\n", i, i*2)
    case <-exit:
      fmt.Println("Quitting...")
      return
    }
    wg.Done()
  }
}

func main() {
  half := make(chan int, 1)
  double := make(chan int, 1)
  exit := make(chan bool)
  wg := &sync.WaitGroup{}

  for i := 1; i < 11; i++ {
    i := i  // Make i scoped inside the for loop.
    wg.Add(1)
    go func(){
      half <- i
    }()

    wg.Add(1)
    go func() {
      double <- i
    }()
  }
  // This just prevents everything from being in order.  Ignore.
	time.Sleep(2 * time.Second)

  go printer(half, double, wg, exit)

  wg.Wait()
  close(exit)

  // Sleep enough time to let "Quitting" print.
  time.Sleep (2 * time.Second)
}

// Enumeration using constants.
const (
  apple = iota  // apple == 0
  orange // orange == 1
  pear  // pear == 2
)

//////////////////////////
// Extended class
//////////////////////////

func Divide(n int, d int) (float64, error) {
  if d == 0 {
    return 0.0, fmt.Errorf("cannot divide by 0")
  }
  return float64(n)/float64(d), nil
}

func TestDivide(t *testing.T) {
  tests := []struct{
    desc string
    n int
    d int
    err bool
    want float64
  }{
    {"divide by 0 error", 10, 0, true, 0.0},
    {"success", 10, 3, false, float64(10)/float64(3)},
  }

  for _, test := range tests {
    got, err := Divide(test.n, test.d)
    switch {
    case err == nil && test.err:
      t.Errorf("TestDivide(%s): got err == nil, want err != nil", test.desc)
      continue
    case err != nil && !test.err:
      t.Errorf("TestDivide(%s): got err == %s, want err == nil", test.desc, err)
      continue
    case err != nil:
      continue
    }
    if got != test.want {
      t.Errorf("TestDivide(%s): got %v, want %v", test.desc, got, test.want)
    }
  }
}

var RecNotFound = errors.New("not found")

func IsNotFound(e error) bool {
  if e == RecNotFound {
    return true
  }
  return false
}

type ClassID int

type Record struct {
  ID int
  LastName, FirstName string
  Classes []ClassID
}

type Storage interface {
  Record(id string) (Record, error)
  Store(r Record) error
}

const (
  unknown = iota
  read
  write
)

type req struct {
    rType int
    id int
    rec Record
    done chan resp
}

type resp struct {
  rec Record
  err error
}

type Registrar struct {
  store Storage
  ch chan req
}

func New(s Storage) Registrar {
  return Registrar{
    store: s,
    ch: make(chan req, 100),
  }
}

func (r Registrar) Start() {
  go func() {
    for input := range r.ch {
      go r.inputHandler(input)
    }
  }()
}

func (r Registrar) inputHandler(input req) {
  defer close(input.done)

  switch input.rType{
  case read:
    rec, err := r.store.Record(input.id)
    if err != nil {
        input.done <- resp{err: err}
        return
    }
    input.done <- resp{rec: rec}
  case write:
    if err := r.store.Store(input.rec); err != nil {
      input.done <- resp{err: err}
      return
    }
    input.done <- resp{}
  default:
    input.done <- resp{err: errors.New("operation on Registrar not a recognized type: %d", input.rType)}
  }
}

func (r Registrar) Read(id string) (Record, error) {
  q := req{rType: read, id: id, done: make(chan resp, 1)}
  r.ch <- q
  a := <-q.done
  if a.err != nil {
    return Record{}, a.err
  }
  return a.rec, nil
}

func (r Registrar) Write(rec Record) error {
  q := req{rType: write, rec: rec, done: make(chan resp, 1)}
  r.ch <- q
  a := <-q.done
  if a.err != nil {
    return a.err
  }
  return nil
}

// fakeStore impleme3ents Storage
type fakeStore struct {
  mu sync.RWMutex
  data map[string]Record
  rErr, wErr bool
}

func (f *fakeStore) Record(id string) (Record, error) {
  f.mu.RLock()
  defer f.mu.RUnlock()

  if f.rErr {
    return Record{}, errors.New("read error")
  }
  if r, ok := f.data[id]; ok {
    return r, nil
  }
  return Record{}, RecNotFound
}

func (f *fakeStore) Store(r Record) error {
  f.mu.Lock()
  defer f.mu.Unlock()
  if f.wErr {
    return errors.New("write error")
  }
  f.data[r.ID] = r
  return nil
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}

// Data represents data being stored.
type Data struct {
	// FirstName is the person's first name.
	// LastName is the person's last name.
	FirstName, LastName string
	// ID is a unique UUIDv4 identifier for this person.
	ID string
	// cache holds the cached data that can be used in String(). It is
	// private, so it will not be Serialized.
	cache string
}

func (d Data) validate() error {
	switch "" {
	case d.FirstName, d.LastName:
		return fmt.Errorf("cannot have FirstName or LastName as an empty string")
	}
	return nil
}

// String implements fmt.Stringer.
func (d *Data) String() string {
	if d.cache == "" {
		d.cache = fmt.Sprintf("%s,%s", d.LastName, d.FirstName)
	}
	return d.cache
}

// StoreData writes Data to a file in os.Tempdir using the ID as
// the file name. If the file exists, it is overwritten.
func StoreData(d Data) error {
	if err := d.validate(); err != nil {
		return err
	}

	// Bonus points if someone can tell me why you should not do this.
	n := filepath.Join(os.TempDir(), d.ID)
	f, err := os.OpenFile(n, os.O_CREATE+os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("could not open file %q: %s", n, err)
	}

	enc := gob.NewEncoder(f)
	if err := enc.Encode(d); err != nil {
		return fmt.Errorf("problems writing file: %s", err)
	}

	return nil
}

// GetData returns Data for a specific ID.
func GetData(id string) (Data, error) {
	n := filepath.Join(os.TempDir(), id)
	f, err := os.Open(n)
	if err != nil {
		return Data{}, fmt.Errorf("could not open file %q: %s", n, err)
	}
	dec := gob.NewDecoder(f)

	d := Data{}
	if err := dec.Decode(&d); err != nil {
		return d, fmt.Errorf("could not decode file %q: %s", n, err)
	}
	return d, nil
}

type Data struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ID        string `json:"id,omitempty"`
	Cache     string `json:"-"`
}

// Leaving out the String() and validate() methods for space.

func StoreData(d Data) error {
	if err := d.validate(); err != nil {
		return err
	}

	b, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("problem JSON encoding data: %s", err)
	}

	n := filepath.Join(os.TempDir(), d.ID)
	if err := ioutil.WriteFile(n, b, 0666); err != nil {
		return fmt.Errorf("problems writing file %q: %s", n, err)
	}

	return nil
}

func GetData(id string) (Data, error) {
	n := filepath.Join(os.TempDir(), id)

	b, err := ioutil.ReadFile(n)
	if err != nil {
		return Data{}, fmt.Errorf("could not open file %q: %s", n, err)
	}

	d := Data{}

	if err := json.Unmarshal(b, &d); err != nil {
		return d, fmt.Errorf("could not decode file %q: %s", n, err)
	}
	return d, nil
}


type Recs []Rec

func (r Recs) Marshal(w *csv.Writer) error {
	for _, rec := range r {
		w.Write(rec.marshal())
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func (r *Recs) Unmarshal(read *csv.Reader) error {
	for i := 0; true; i++{
		fields, err := read.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		rec := Rec{}
		if err := rec.unmarshal(fields); err != nil {
			if i == 0 {
				continue // First line can be the column headers
			}
			return err
		}

		*r = append(*r, rec)
	}
	return nil
}

type Rec struct {
	LastName, FirstName string
	ID int
}

func (r Rec) String() string {
	return fmt.Sprintf("%d::%s::%s", r.ID, r.LastName, r.FirstName)
}

func (r Rec) marshal() []string {
	return []string{strconv.Itoa(r.ID), r.LastName, r.FirstName}
}

func (r *Rec) unmarshal(f []string) error {
	if len(f) != 3 {
		return fmt.Errorf("expected 3 fields, got %d", len(f))
	}
	i, err := strconv.Atoi(f[0])
	if err != nil {
		return fmt.Errorf("first field could not be "+
			"converted to integer id: %s", err)
	}
	r.ID = i
	r.LastName = f[1]
	r.FirstName = f[2]
	return nil
}

type Recs []Rec

func (r Recs) Marshal(w *csv.Writer) error {
	for _, rec := range r {
		w.Write(rec.marshal())
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func (r *Recs) Unmarshal(read *csv.Reader) error {
	for i := 0; true; i++ {
		fields, err := read.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		rec := Rec{}
		if err := rec.unmarshal(fields); err != nil {
			if i == 0 {
				continue // First line can be the column headers
			}
			return err
		}

		*r = append(*r, rec)
	}
	return nil
}


syntax = "proto3";

package record;

message Rec {
  int64 id = 1;
  string last_name = 2;
  string first_name = 3;
}

func main() {
  rec := &record.Rec{
    Id: "10ba038e-48da-487b-96e8-8d3b99b6d18a",
    LastName: "Doak",
    FirstName: "John",
  }

  b, err := proto.Marshal(rec)
  if err != nil {
    panic(err)
  }

  if err := ioutil.WriteFile(rec.Id, b, 0666); err != nil {
    panic(err)
  }
}

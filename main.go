package main

import (
    "fmt"
    "os"
    "net/http"
)

func fizzbuzz(x int){
    for i := 1; i <= x; i++ {
        if i % 3 == 0 {
            fmt.Print("Fizz");
        } 
        if i % 5 == 0 {
            fmt.Print("Buzz");
        }
        if i % 5 != 0 && i % 3 != 0 {
            fmt.Print(i);
        }
        fmt.Println();
    }
}

func fib_rec(x int)(int){
    if x < 2 {
        return x 
    }
    return (fib_rec(x - 1)) + (fib_rec(x - 2)) 
}

func fib_iter(x int) {
    prev := 0
    next := 1 
    result := 1
    for i := 0; i < x; i++ {
        fmt.Println(result);
        result = prev + next;
        prev = next;
        next = result;
    }
}

func price_of_tile(width int, height int, price float32)(float32){
    return float32(width) * float32(height) * price
}

func handle_error(err error){
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func reverse_string(str string)(string) {
    var result string
    for i, _ := range str {
        result = string(str[i]) + result
    }
    return result
}

func hello(w http.ResponseWriter, req *http.Request){
    http.ServeFile(w, req, req.URL.Path[1:])
}


func api(w http.ResponseWriter, req *http.Request, count *int) {
    fmt.Fprintf(w, "This is an api %d", *count)
    *count += 69
}

type stuff struct {
    x, y int
}

func (s stuff) String() string {
    return fmt.Sprintf("(%d,%d)", s.x, s.y)
}

func do_thing(s *stuff)(int){
    return s.x + s.y
}


func main(){
    thing := stuff{69, 420}
    fmt.Println(do_thing(&thing))
}

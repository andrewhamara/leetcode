import (
    "strconv"
)

func fizzBuzz(n int) []string {
    var fizz []string
    for i := 1; i < n+1; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fizz = append(fizz, "FizzBuzz")
        }else if i % 3 == 0 {
            fizz = append(fizz, "Fizz")
        }else if i % 5 == 0 {
            fizz = append(fizz, "Buzz")
        }else {
            fizz = append(fizz, strconv.Itoa(i))
        }
    }
    return fizz
}

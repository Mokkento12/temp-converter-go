package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
    var input float64
    fmt.Print("Введите температуру в градусах Цельсия: ")
    fmt.Scan(&input)

    c := Celsius(input)
    f := CToF(c)
    fmt.Printf("%v°C = %v°F\n", c, f)

    fmt.Print("Введите температуру в градусах Фаренгейта: ")
    fmt.Scan(&input)

    f = Fahrenheit(input)
    c = FToC(f)
    fmt.Printf("%v°F = %v°C\n", f, c)
}
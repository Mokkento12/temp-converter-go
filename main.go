package main

import (
	"fmt"
	"tempconv-project/tempconv"
)

func main() {
    var input float64

    // Преобразование из Цельсия в Фаренгейт
    fmt.Print("Введите температуру в градусах Цельсия: ")
    fmt.Scan(&input)

    c := tempconv.Celsius(input)
    f := tempconv.CToF(c)
    fmt.Printf("%s = %s\n", c, f)

    // Преобразование из Фаренгейта в Цельсий
    fmt.Print("Введите температуру в градусах Фаренгейта: ")
    fmt.Scan(&input)

    f = tempconv.Fahrenheit(input)
    c = tempconv.FToC(f)
    fmt.Printf("%s = %s\n", f, c)
}
package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv-project/tempconv"
)

func main() {
    // Проверяем, что передан хотя бы один аргумент
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Использование: %s значение\n", os.Args[0])
        os.Exit(1)
    }

    fmt.Println("=== Преобразование температур ===")

    // Обрабатываем каждый аргумент командной строки
    for _, arg := range os.Args[1:] {
        // Преобразуем аргумент в число
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
            os.Exit(1)
        }

        // Преобразуем температуру из Цельсия в Фаренгейт
        c := tempconv.Celsius(t)
        f := tempconv.CToF(c)
        fmt.Printf("Цельсий: %s -> Фаренгейт: %s\n", c, f)

        // Преобразуем температуру из Фаренгейта в Цельсий
        fahrenheit := tempconv.Fahrenheit(t)
        celsius := tempconv.FToC(fahrenheit)
        fmt.Printf("Фаренгейт: %s -> Цельсий: %s\n", fahrenheit, celsius)

        fmt.Println("===============================")
    }
}
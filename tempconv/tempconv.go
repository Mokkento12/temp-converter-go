package tempconv

import "fmt"

// Типы для температуры
type Celsius float64
type Fahrenheit float64

// Преобразование из Цельсия в Фаренгейт
func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

// Преобразование из Фаренгейта в Цельсий
func FToC(f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}

// Строковое представление температуры в Цельсиях
func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}

// Строковое представление температуры в Фаренгейтах
func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g°F", f)
}
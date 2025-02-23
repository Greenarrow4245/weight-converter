package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Welcome to Weight Converter v2!\n")
    time.Sleep(1 * time.Second)

    fmt.Println("Enter a value for conversion: ")
    var value float64
    fmt.Scanln(&value)

    time.Sleep(500 * time.Millisecond)

    fmt.Println("Enter current unit (kg, g, lb, oz, t, mg, st, mcg, it): ")
    var currentUnit string
    fmt.Scanln(&currentUnit)

    time.Sleep(500 * time.Millisecond)

    fmt.Println("Enter unit to convert to (kg, g, lb, oz, t, mg, st, mcg, it): ")
    var targetUnit string
    fmt.Scanln(&targetUnit)

    fmt.Println("\nConverting...")
    time.Sleep(1 * time.Second) // **Added delay for conversion effect**

    convertedValue := ConvertUnits(value, currentUnit, targetUnit)
    fmt.Printf("Converted value: %.5f %s\n", convertedValue, targetUnit)
}

// ConvertUnits converts from any supported unit to another
func ConvertUnits(value float64, fromUnit string, toUnit string) float64 {
    valueInKg := toKilogram(value, fromUnit)
    return fromKilogram(valueInKg, toUnit)
}

// Convert any unit to kilograms
func toKilogram(value float64, unit string) float64 {
    switch unit {
    case "kg":
        return value
    case "g":
        return value / 1000
    case "lb":
        return value * 0.453592
    case "oz":
        return value * 0.0283495
    case "t":
        return value * 1000
    case "mg":
        return value / 1e6
    case "st":
        return value * 6.35029
    case "mcg":
        return value / 1e9
    case "it": // Imperial ton
        return value * 1016.05
    default:
        fmt.Println("Unsupported 'from' unit. Please use one of the specified units.")
        return 0
    }
}

// Convert kilograms to the target unit
func fromKilogram(value float64, unit string) float64 {
    switch unit {
    case "kg":
        return value
    case "g":
        return value * 1000
    case "lb":
        return value / 0.453592
    case "oz":
        return value / 0.0283495
    case "t":
        return value / 1000
    case "mg":
        return value * 1e6
    case "st":
        return value / 6.35029
    case "mcg":
        return value * 1e9
    case "it": // Imperial ton
        return value / 1016.05
    default:
        fmt.Println("Unsupported 'to' unit. Please use one of the specified units.")
        return 0
    }
}

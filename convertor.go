package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type UnitConverter interface {
	ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error)
}

type Mass struct{}

func (m *Mass) ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error) {
	units := map[string]map[string]float64{
		"kilogram":  {"gram": 1000, "pound": 2.20462, "tonne": 0.001, "milligram": 1e6},
		"gram":      {"kilogram": 0.001, "pound": 0.00220462, "tonne": 1e-6, "milligram": 1000},
		"pound":     {"kilogram": 0.453592, "gram": 453.592, "tonne": 0.01638, "milligram": 16380000},
		"tonne":     {"gram": 1e6, "pound": 61.05, "kilogram": 1000, "milligram": 1e9},
		"milligram": {"gram": 1e-3, "pound": 6.11e-8, "tonne": 1e-9, "kilogram": 1e-6},
	}
	if targetMap, ok := units[fromUnit]; ok {
		if factor, ok := targetMap[targetUnit]; ok {
			return value * factor, nil
		}
	}
	return 0, errors.New("Не поддерживается конвертация.")
}

type Length struct{}

func (l *Length) ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error) {
	units := map[string]map[string]float64{
		"meter":      {"kilometer": 1e-3, "centimeter": 100, "decimeter": 10, "millimeter": 1000},
		"kilometer":  {"meter": 1000, "centimeter": 1e5, "decimeter": 1e4, "millimeter": 1e6},
		"centimeter": {"meter": 1e-2, "kilometer": 1e-5, "decimeter": 0.1, "millimeter": 10},
		"decimeter":  {"meter": 0.1, "kilometer": 1e-4, "centimeter": 10, "millimeter": 100},
		"millimeter": {"meter": 1e-3, "kilometer": 1e-6, "centimeter": 0.1, "decimeter": 0.01},
	}
	if targetMap, ok := units[fromUnit]; ok {
		if factor, ok := targetMap[targetUnit]; ok {
			return value * factor, nil
		}
	}
	return 0, errors.New("Не поддерживается конвертация.")
}

type Pressure struct{}

func (p *Pressure) ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error) {
	units := map[string]map[string]float64{
		"pascal":                            {"millimetres_of_the_mercury_column": 0.007501, "metres_of_the_water_column": 0.101972, "atmosphere": 1e-5, "bar": 1e-5},
		"millimetres_of_the_mercury_column": {"pascal": 133.32, "metres_of_the_water_column": 13.6, "atmosphere": 0.001316, "bar": 0.001333},
		"metres_of_the_water_column":        {"pascal": 9.81, "metres_of_the_mercury_column": 0.073556, "atmosphere": 0.000097, "bar": 0.000098},
		"atmosphere":                        {"pascal": 101325, "metres_of_the_water_column": 10332.27, "metres_of_the_mercury_column": 760, "bar": 1.01},
		"bar":                               {"pascal": 100000, "metres_of_the_water_column": 10197.16, "metres_of_the_mercury_column": 750.06, "atmosphere": 0.986923},
	}
	if targetMap, ok := units[fromUnit]; ok {
		if factor, ok := targetMap[targetUnit]; ok {
			return value * factor, nil
		}
	}
	return 0, errors.New("Не поддерживается конвертация.")
}

type Speed struct{}

func (s *Speed) ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error) {
	units := map[string]map[string]float64{
		"meter_per_second":     {"kilometer_per_hour": 3.6, "kilometer_per_second": 1e-3, "sea_speed": 1.94384, "Mach_number": 0.002931},
		"kilometer_per_second": {"kilometer_per_hour": 3600, "meter_per_second": 1e3, "sea_speed": 1943.84, "Mach_number": 2.94},
		"kilometer_per_hour":   {"meter_per_second": 0.277778, "kilometer_per_second": 0.000278, "sea_speed": 0.53996, "Mach_number": 0.00081},
		"sea_speed":            {"meter_per_second": 0.514444, "kilometer_per_second": 0.000514, "kilometer_per_hour": 1.852, "Mach_number": 0.0015},
		"Mach_number":          {"meter_per_second": 343, "kilometer_per_second": 0.343, "kilometer_per_hour": 1235.52, "sea_speed": 666.739},
	}
	if targetMap, ok := units[fromUnit]; ok {
		if factor, ok := targetMap[targetUnit]; ok {
			return value * factor, nil
		}
	}
	return 0, errors.New("Не поддерживается конвертация.")
}

type Volume struct{}

func (v *Volume) ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error) {
	units := map[string]map[string]float64{
		"volume_meter":      {"volume_decimeter": 1000, "volume_centimeter": 1e6, "volume_millimeter": 1e9, "volume_kilometer": 0.000001, "liter": 1000},
		"volume_decimeter":  {"volume_meter": 0.001, "volume_centimeter": 1000, "volume_millimeter": 1e6, "volume_kilometer": 1e-12, "liter": 1},
		"volume_centimeter": {"volume_meter": 1e-6, "volume_decimeter": 1e-3, "volume_millimeter": 1e3, "volume_kilometer": 1e-15, "liter": 1e-3},
		"volume_millimeter": {"volume_meter": 1e-9, "volume_decimeter": 1e-6, "volume_centimeter": 1e-3, "volume_kilometer": 1e-18, "liter": 1e-6},
		"volume_kilometer":  {"volume_meter": 1e9, "volume_decimeter": 1e12, "volume_centimeter": 1e15, "volume_millimeter": 1e18, "liter": 1e12},
		"liter":             {"volume_meter": 1e-3, "volume_decimeter": 1, "volume_centimeter": 1e3, "volume_millimeter": 1e6, "volume_kilometer": 1e-12},
	}
	if targetMap, ok := units[fromUnit]; ok {
		if factor, ok := targetMap[targetUnit]; ok {
			return value * factor, nil
		}
	}
	return 0, errors.New("Не поддерживается конвертация.")
}

type Angle struct{}

func (a *Angle) ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error) {
	units := map[string]map[string]float64{
		"degrees":   {"radians": 0.0174533, "grads": 1.11111, "turnovers": 0.00277778, "sec": 3600, "min": 60, "hour": 0.0666667},
		"radians":   {"degrees": 57.2958, "grads": 63.662, "turnovers": 0.159155, "sec": 206265, "min": 3437.85, "hour": 3.81972},
		"grads":     {"degrees": 0.9, "radians": 0.015708, "turnovers": 0.0025, "sec": 3240, "min": 54, "hour": 0.06},
		"turnovers": {"degrees": 360, "radians": 6.28319, "grads": 400, "sec": 1296000, "min": 21600, "hour": 24},
		"sec":       {"degrees": 0.000277778, "radians": 4.84814e-6, "grads": 0.000308642, "turnovers": 1.2861e-6, "min": 0.0166667, "hour": 0.0000185185},
		"min":       {"degrees": 0.0166667, "radians": 0.000290888, "grads": 0.0185185, "turnovers": 0.0000463, "sec": 60, "hour": 0.00111111},
		"hour":      {"degrees": 15, "radians": 0.261799, "grads": 16.6667, "turnovers": 0.0416667, "sec": 54000, "min": 900},
	}
	if targetMap, ok := units[fromUnit]; ok {
		if factor, ok := targetMap[targetUnit]; ok {
			return value * factor, nil
		}
	}
	return 0, errors.New("Не поддерживается конвертация.")
}

type Temperature struct{}

func (t *Temperature) ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error) {
	if fromUnit == "Celsius" && targetUnit == "Fahrenheit" {
		return value*9/5 + 32, nil
	} else if fromUnit == "Fahrenheit" && targetUnit == "Celsius" {
		return (value - 32) * 5 / 9, nil
	} else if fromUnit == "Celsius" && targetUnit == "Kelvin" {
		return value + 273.15, nil
	} else if fromUnit == "Kelvin" && targetUnit == "Celsius" {
		return value - 273.15, nil
	} else if fromUnit == "Fahrenheit" && targetUnit == "Kelvin" {
		return (value-32)*5/9 + 273.15, nil
	} else if fromUnit == "Kelvin" && targetUnit == "Fahrenheit" {
		return (value-273.15)*9/5 + 32, nil
	}
	return 0, errors.New("Не поддерживается конвертация.")
}

type Area struct{}

func (a *Area) ConvertTo(fromUnit string, targetUnit string, value float64) (float64, error) {
	units := map[string]map[string]float64{
		"square_meter":      {"square_kilometer": 1e-6, "square_centimeter": 1e4, "square_decimeter": 1e2, "quare_millimeter": 1e6},
		"square_kilometer":  {"square_meter": 1e6, "square_centimeter": 1e10, "square_decimeter": 1e8, "quare_millimeter": 1e12},
		"square_centimeter": {"square_meter": 1e-4, "square_kilometer": 1e-10, "square_decimeter": 1e-2, "quare_millimeter": 1e2},
		"square_decimeter":  {"square_meter": 1e-2, "square_kilometer": 1e-8, "square_centimeter": 1e2, "square_millimeter": 1e4},
		"square_millimeter": {"square_meter": 1e-6, "square_kilometer": 1e-12, "square_centimeter": 1e-2, "square_decimeter": 1e-4},
	}
	if targetMap, ok := units[fromUnit]; ok {
		if factor, ok := targetMap[targetUnit]; ok {
			return value * factor, nil
		}
	}
	return 0, errors.New("Не поддерживается конвертация.")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getUnitChoice(units []string) string {
	for i, unit := range units {
		fmt.Printf("%d. %s\n", i+1, unit)
	}
	choice := getUserInput("\nВыберите единицу (введите номер): ")
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(units) {
		fmt.Println("Ошибка: неверный выбор.")
		return getUnitChoice(units)
	}
	return units[index-1]
}

func main() {
	for {
		fmt.Println("Выберите тип конверсии:")
		fmt.Println("1. Масса")
		fmt.Println("2. Длина")
		fmt.Println("3. Площадь")
		fmt.Println("4. Объём")
		fmt.Println("5. Скорость")
		fmt.Println("6. Давление")
		fmt.Println("7. Углы")
		fmt.Println("8. Температура")
		fmt.Println("0. Выход\n")

		choice := getUserInput("Введите номер опции: ")
		var converter UnitConverter
		var units []string

		switch choice {
		case "0":
			fmt.Println("Выход из программы.")
			return
		case "1":
			converter = &Mass{}
			units = []string{"kilogram", "gram", "pound", "tonne", "milligram"}
		case "2":
			converter = &Length{}
			units = []string{"meter", "kilometer", "centimeter", "decimeter", "meter"}
		case "3":
			converter = &Area{}
			units = []string{"square_meter", "square_kilometer", "square_centimeter", "square_decimeter", "square_millimeter"}
		case "4":
			converter = &Volume{}
			units = []string{"volume_meter", "volume_kilometer", "volume_centimeter", "volume_decimeter", "volume_millimeter", "liter"}
		case "5":
			converter = &Speed{}
			units = []string{"meter_per_second", "kilometer_per_second", "kilometer_per_hour", "sea_speed", "Mach_number"}
		case "6":
			converter = &Pressure{}
			units = []string{"pascal", "atmosphere", "bar", "metres_of_the_water_column", "millimetres_of_the_mercury_column"}
		case "7":
			converter = &Angle{}
			units = []string{"degrees", "radians", "grads", "turnovers", "sec", "min", "hour"}
		case "8":
			converter = &Temperature{}
			units = []string{"Celsius", "Fahrenheit", "Kelvin"}
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
			continue
		}

		fromUnit := getUnitChoice(units)
		toUnit := getUnitChoice(units)

		valueStr := getUserInput("Введите значение: ")
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println("Ошибка: неверное значение.")
			continue
		}

		result, err := converter.ConvertTo(fromUnit, toUnit, value)
		if err != nil {
			fmt.Println("Ошибка конверсии:", err)
		} else {
			fmt.Printf("%.15f %s = %.15f %s\n", value, fromUnit, result, toUnit)
		}
	}
}

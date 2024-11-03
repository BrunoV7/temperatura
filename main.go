package main

import (
	"fmt"
	"os"
)

func main() {
	var option int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Fahrenheit para Celsius")
		fmt.Println("2. Celsius para Fahrenheit")
		fmt.Println("3. Kelvin para Celsius")
		fmt.Println("4. Celsius para Kelvin")
		fmt.Println("5. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fahrenheitToCelsius()
		case 2:
			celsiusToFahrenheit()
		case 3:
			kelvinToCelsius()
		case 4:
			celsiusToKelvin()
		case 5:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func fahrenheitToCelsius() {
	var fahrenheit float64
	fmt.Print("Digite a temperatura em Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f Fahrenheit é igual a %.2f Celsius\n", fahrenheit, celsius)
}

func celsiusToFahrenheit() {
	var celsius float64
	fmt.Print("Digite a temperatura em Celsius: ")
	fmt.Scanln(&celsius)
	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f Celsius é igual a %.2f Fahrenheit\n", celsius, fahrenheit)
}

func kelvinToCelsius() {
	var kelvin float64
	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scanln(&kelvin)
	celsius := kelvin - 273
	fmt.Printf("%.2f Kelvin é igual a %.2f Celsius\n", kelvin, celsius)
}

func celsiusToKelvin() {
	var celsius float64
	fmt.Print("Digite a temperatura em Celsius: ")
	fmt.Scanln(&celsius)
	kelvin := celsius + 273
	fmt.Printf("%.2f Celsius é igual a %.2f Kelvin\n", celsius, kelvin)
}

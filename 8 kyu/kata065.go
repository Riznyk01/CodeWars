package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(GetPlanetName(num))
}

func GetPlanetName(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	case 9:
		return "Pluto"
	default:
		return ""
	}
}

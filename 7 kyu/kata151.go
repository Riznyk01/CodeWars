package main

import "fmt"

func main() {
	fmt.Printf("%c", Triangle("RBRGBRBGGRRRBGBBBGG"))
}

func Triangle(row string) rune {
	res := row
	for len(res) > 1 {
		func() {
			out := ""
			for i := 0; i < len(res)-1; i++ {
				switch res[0+i : 2+i] {
				case "BB", "RG", "GR":
					out += "B"
				case "RR", "BG", "GB":
					out += "R"
				case "GG", "BR", "RB":
					out += "G"
				}
			}
			res = out
		}()
	}
	return rune(res[0])
}

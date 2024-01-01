package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func main() {
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))
}
func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	var winner string
	for {
		if fighter1.Name == firstAttacker {
			fighter2.Health -= fighter1.DamagePerAttack
			if fighter2.Health <= 0 {
				winner = fighter1.Name
				break
			}
			fighter1.Health -= fighter2.DamagePerAttack
			if fighter1.Health <= 0 {
				winner = fighter2.Name
				break
			}
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			if fighter1.Health <= 0 {
				winner = fighter2.Name
				break
			}
			fighter2.Health -= fighter1.DamagePerAttack
			if fighter2.Health <= 0 {
				winner = fighter1.Name
				break
			}
		}
	}
	return winner
}

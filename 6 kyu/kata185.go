package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(DeadAntCount("...ant...ant..nat.ant.t..ant...ant..ant..ant.anant..t"))
}
func DeadAntCount(ants string) int {
	counts := make([]int, 3)
	parts := strings.NewReplacer("ant", "").Replace(ants)

	counts[0] = strings.Count(parts, "a")
	counts[1] = strings.Count(parts, "n")
	counts[2] = strings.Count(parts, "t")

	return slices.Max(counts)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(XMasTree(2))
}
func XMasTree(height int) []string {
	if height < 1 {
		return []string{}
	}
	a := make([]string, height+2)
	h, ll := 2*height-1, 0
	t := strings.Repeat("_", height-1) + "#" + strings.Repeat("_", height-1)
	a[len(a)-1], a[len(a)-2] = t, t
	for i := height - 1; i >= 0; i-- {
		a[i] = strings.Repeat("_", ll) + strings.Repeat("#", h) + strings.Repeat("_", ll)
		h -= 2
		ll++
	}
	return a
}

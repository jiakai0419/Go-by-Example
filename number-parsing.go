package main

import "fmt"
import "strconv"

func main() {
	P := fmt.Println

	f, _ := strconv.ParseFloat("1.234", 64)
	P(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	P(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	P(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	P(u)

	k, _ := strconv.Atoi("135")
	P(k)

	_, e := strconv.Atoi("wat")
	P(e)
}

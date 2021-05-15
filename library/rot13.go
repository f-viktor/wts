// its actually all of the 'rot's
// need one that shifts through the entire space and prints everything that's fully printable

package wts

import (
  "fmt"
)


func RotMain(message string) {
  sayHi("Rot/Caesar Decode", message)

  for i := 0; i < 27; i++ {
	 r := rotN(message,i)
   fmt.Printf("Rot%d => "+r+"\n", i)
  }
  r := rotN(message,47)
  fmt.Printf("Rot%d => "+r+"\n", 47)
}


func rotNgine(b rune,n int) rune {
	var a, z rune
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+rune(n))%(z-a+1) + a
}

func rotN(message string, n int) string {
  p := []rune(message)
	for i := 0; i < len(p); i++ {
		p[i] = rotNgine(p[i],n)
	}
	return string(p)
}

// its actually all of the 'rot's
// need one that shifts through the entire space and prints everything that's fully printable

package wts

import (
  "fmt"
  "strings"
)


func RotMain(s string) {

  fmt.Println("\n[+]_Rot/Caesar_Decode_"+strings.Repeat("_",len(s)))
  for i := 0; i < 27; i++ {
	 r := rotN(s,i)
   fmt.Printf("Rot%d => "+r+"\n", i)
  }
  r := rotN(s,47)
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
	for i := 0; i < len(message); i++ {
		p[i] = rotNgine(p[i],n)
	}
	return string(p)
}

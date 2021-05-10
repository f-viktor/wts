// replaces everything that isnt a number, with a whitespace, then splits on whitespaces
// also tries to print it as an offset from the letter A
// if its just a line of numbers, it will try to split it after 1/2/3 digits and print it

package wts

import (
  "fmt"
  "strings"
  "regexp"
  "strconv"
)


func DecimalMain(s string) {

  offsetarray := []int{0,'a','A',32,48}

  fmt.Println("\n[+]_Decimal_Decode_"+strings.Repeat("_",len(s)))
  // match antyting that isnt a number
  reNaN := regexp.MustCompile(`[^0-9]+`)
  dec := reNaN.ReplaceAllString(s, " ")
  decarray := strings.Split(dec," ")

  var intarray []int
  for _, item := range decarray {
    num, _ := strconv.Atoi(item)
    intarray = append(intarray, num)
  }

  printDecodingArray(intarray)
  callOffsets(intarray, offsetarray)

  intarray = splitAfterN(s,1)
  printDecodingArray(intarray)
  callOffsets(intarray, offsetarray)

  intarray = splitAfterN(s,2)
  printDecodingArray(intarray)
  callOffsets(intarray, offsetarray)

  intarray = splitAfterN(s,3)
  printDecodingArray(intarray)
  callOffsets(intarray, offsetarray)
}




func callOffsets(intarray []int, offsetarray []int){
  for _, item := range offsetarray{
    abcOffset(intarray,item)
  }
}

// use array as offset from base to create string
func abcOffset(intarray []int, base int) {
  fmt.Printf("Offset from '"+string(base)+"' Decode => ")
  for _, item := range intarray {
    char := string(item+base)
    fmt.Printf(char)
  }
  fmt.Printf("\n")
}

// split into an integer after every N-th rune
func splitAfterN(s string, N int) []int {
  a:= []rune(s)
  var res string
  var intarray []int
  for i, r := range a {
      res = res + string(r)
      if i > 0 && (i+1)%N == 0 {
          num, _ :=strconv.Atoi(res)
          intarray = append(intarray, num)
          res = ""
      }
    }
  return intarray
}

func printDecodingArray(intarray []int){
  fmt.Printf("[+] Decoding ")
  for _, item := range intarray {
    fmt.Printf("%d ", item)
  }
  fmt.Printf("\n")
}
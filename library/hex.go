package wts

import(
  "fmt"
  "strings"
  "strconv"
)


func HexMain(message string){
  sayHi("0xHex Decode",message)
  //fmt.Println("\n[>]_Binary_Decode_"+strings.Repeat("_",len(message)))
  //fmt.Println("[?] Decoding Target "+message)


  //fmt.Println("[+] Reduced to most frequent characters ",message)

  hexDecodeMessage(message)
  message = removeInvalid(message,"0123456789abcdefABCDEF")

  for _, i := range []int{2,4,8,16} {
    split := forceDelimiter(message, i," ")
    hexDecodeMessage(split) //split to 8 chars
  }

  fmt.Println("[+] Switching enidianness")

  split := switchEnidianness(message, 999999) // reverse the whole thing
  split = forceDelimiter(split, 2," ") // print each byte as reverse
  hexDecodeMessage(split)

  for _, i := range []int{2,4,8,16} {
    split := switchEnidianness(message, i)
    hexDecodeMessage(split)
  }
}



func hexDecodeMessage(dec string) {

  fmt.Println("[?] Decoding "+dec)
  binArray := strings.Split(dec, " ")

  // Try hex decoding
  error := false
  var decodedChars string

  for _, val := range binArray{
    binaryCode, err := strconv.ParseInt(val, 16, 64)
    if err != nil {error = true}
    decodedChars = decodedChars + " " + string(binaryCode)
  }
  decodedChars = decodedChars[1:] //remove first space

  if error {
		  safePrintln("[x] Hex => " +decodedChars)
	} else {  safePrintln("[√] Hex => " +decodedChars)}
}

// reverse byte order in groups of <split>
func switchEnidianness(dec string, split int) string {
    bytesString := forceDelimiter(dec, 2," ") //split to bytes
    bytesArray := strings.Split(bytesString, " ")

    var transformed []string

    for i:=0; i<len(bytesArray);i=i+split {
      chunk := split
      //min(split,len(array)-i)
      if (i+split >= len(bytesArray)) {chunk = len(bytesArray)-i}
      transformed = append(transformed,reverseStrSlice(bytesArray[i:i+chunk])...)
    }

    bytesString = forceDelimiter(strings.Join(transformed,""), split*2, " ") //split to bytes

    return bytesString
}

func reverseStrSlice(arr []string) []string{
   for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
      arr[i], arr[j] = arr[j], arr[i]
   }
   return arr
}

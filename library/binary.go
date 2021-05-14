package wts

import(
  "fmt"
  "strings"
  "regexp"
  "strconv"
)


func BinaryMain(message string){
  sayHi("Binary Decode",message)
  //fmt.Println("\n[>]_Binary_Decode_"+strings.Repeat("_",len(message)))
  //fmt.Println("[?] Decoding Target "+message)

  freqMap := CountRunes(message)

  dec := message
  //has at least three types of characters but no more than 4
  regxp := ""
  if len(freqMap) > 2 {
    for _, item := range freqMap[:3]{
      regxp = regxp+string(item.Key)
    }

    // remove everything but the most frequent 3 chars [1;0;delimiter]
    reNaN := regexp.MustCompile(`[^`+regxp+`]+`)
    dec = reNaN.ReplaceAllString(dec, "")
    fmt.Println("[+] Reduced to most frequent characters ",dec)
  }
  if len(freqMap) >= 2 {
    layout := switcharoo(dec,freqMap, []rune{'0','1',' '})
    binaryDecodeMessage(layout)

    layout = strings.Replace(layout, " ","",-1)
    binaryDecodeMessage(forceDelimiter(layout,8," ")) //split to 8 chars

    layout = strings.Replace(layout, " ","",-1)
    binaryDecodeMessage(forceDelimiter(layout,4," ")) //split to 4 chars (will not be printable)

    layout = switcharoo(dec,freqMap, []rune{'1','0',' '})
    binaryDecodeMessage(layout)

    layout = strings.Replace(layout, " ","",-1)
    binaryDecodeMessage(forceDelimiter(layout,8," ")) //split to 8 chars

    layout = strings.Replace(layout, " ","",-1)
    binaryDecodeMessage(forceDelimiter(layout,4," ")) //split to 4 chars (will not be printable)




  } else { fmt.Println("[!] Morse decode requires 2 different type of characters 1;0 (optional:delimiter)")}
}



func binaryDecodeMessage(dec string) {

  fmt.Println("[?] Decoding "+dec)
  binArray := strings.Split(dec, " ")

  // Try binary decoding
  error := false
  var decodedChars string
  var decodedNums string

  for _, val := range binArray{
    binaryCode, err := strconv.ParseInt(val, 2, 64)
    if err != nil {error = true}
    decodedChars = decodedChars + " " + string(binaryCode)
    decodedNums = decodedNums + " " + strconv.FormatInt(binaryCode,10)
  }
  decodedChars = decodedChars[1:] //remove first space
  decodedNums = decodedNums[1:] //remove first space

  if error {
		  safePrintln("[x] Binary => [ "+decodedNums+" ] => " +decodedChars)
	} else {  safePrintln("[x] Binary => [ "+decodedNums+" ] => " +decodedChars)}
}

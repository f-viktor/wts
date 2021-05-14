package wts

import(
  "fmt"
  "strings"
  "regexp"
  "github.com/alwindoss/morse"
)


func MorseMain(message string){
  sayHi("Morse Decode", message)

  freqMap := CountRunes(message)

  //has at least three types of characters but no more than 4
  regxp := ""
  if len(freqMap) >= 3 {
    for _, item := range freqMap[:4]{
      regxp = regxp+string(item.Key)
    }
    // remove everything but the most frequent 4 chars
    reNaN := regexp.MustCompile(`[^`+regxp+`]+`)
    dec := reNaN.ReplaceAllString(message, "")
    fmt.Println("[+] Reduced to most frequent characters ",dec)

    morseDecodeMessage(dec,freqMap, []string{".","-"," "," / "})
    morseDecodeMessage(dec,freqMap, []string{"-","."," "," / "})
    morseDecodeMessage(dec,freqMap, []string{".","-"," / "," "})
    morseDecodeMessage(dec,freqMap, []string{"-","."," / "," "})

  } else {
    fmt.Println("[!] Morse decode requires 3 different type of characters long;short;delimiter (optional:newline)")
  }
}


func morseDecodeMessage(dec string, freqMap []frequency, layout []string) {

  //change this to switcharoo
  // remap chars to .- / for th ee lib
  transformed := strings.Replace(dec, string(freqMap[0].Key), layout[0],-1)
  transformed = strings.Replace(transformed, string(freqMap[1].Key), layout[1],-1)
  transformed = strings.Replace(transformed, string(freqMap[2].Key), layout[2],-1)
  if len(freqMap) > 3 {
      transformed = strings.Replace(transformed, string(freqMap[3].Key), layout[3],-1)
  }
  fmt.Println("[?] Decoding "+transformed)

  // Try standard url decoding first
  h := morse.NewHacker()
  morseCode, err := h.Decode(strings.NewReader(transformed))
  if err != nil {
		  safePrintln("[x] URL => "+string(morseCode))
	} else {  safePrintln("[√] URL => "+string(morseCode))}
}

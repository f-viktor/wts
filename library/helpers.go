package wts

import(
  "fmt"
  "strings"
  //"regexp"
)


func sayHi(name string, message string){
  padding := 80
  if len(message) < 80 {padding = len(message)}

  name = strings.Replace(name," ","_",-1)
  name = name + strings.Repeat("_",40-len(name))

  fmt.Println("\n[>]_"+name+"_"+strings.Repeat("_",padding))
  fmt.Println("[+] Investigating " + message)
}

func freqCheck(){
  fmt.Println("[!] Morse decode requires 2 different type of characters 1;0 (optional:delimiter)")

}

// remap characters based on their frequency to a new layout
func switcharoo (dec string, freqMap []frequency, layout []rune) string {

  var transformed []rune
  for _, original := range []rune(dec){
    for freqIdx, freqChar := range freqMap{
      if original == freqChar.Key{
        transformed = append(transformed, layout[freqIdx])
      }
    }
  }
  return string(transformed)
}

// insert % chars after every second character
// if any of the resulting chars are invalid, the whole decode will failed
// could add something that decodes per character
func forceDelimiter(s string, frequency int, delimiter string) string{
  for i := 0; i < len(s); i += frequency+len(delimiter) {
    s = s[:i] + delimiter + s[i:]
  }
  s=s[1:]

  //fmt.Println("[?] Modified to " + string(s))
  return string(s)
}

// set wether safeprint does anything
var unsafeP = false
func SetUnsafePrint(unsafePrint bool) {
  unsafeP = unsafePrint
}

// replace linebreak and control chars with ¶ to not have linebreaks
func safePrintln(message string) {
  if !unsafeP {
    var printonly []rune
    for _,v := range []rune(message){
      if (v > 31 ) {
        printonly=append(printonly,v)
      } else {
        printonly=append(printonly,'¶')
      }
    }
    message = string(printonly)
  //  space := regexp.MustCompile(`\s+`)
  //  message = space.ReplaceAllString(message, " ")
  }
  fmt.Println(message)
}

package wts

import(
  "fmt"
  "encoding/base32"
  "strings"
)


func Base32Main(message string){
  sayHi("Base32 Decode",message)

  b32DecodeMessage(message)
  message = b32stripUndesired(message)
  if message != "" { b32DecodeMessage(message) }
}

func b32DecodeMessage(message string) {
  fmt.Println("[?] Decoding "+message)

  // Try standard b32 decoding first
  decoded, err := base32.StdEncoding.DecodeString(message)
  if err != nil {
    safePrintln("[x] Base32 (failed/partial) => "+ string(decoded))
  } else {  safePrintln("[√] Base32 => "+string(decoded))}
}

// remove chars that are outside of the characterset
func b32stripUndesired(s string) string{
  var legalChars = `ABCDEFGHIJKLMNOPQRSTUVWXYZ234567`

  stripped := ""
  s = strings.ToUpper(s)

  // for each element of the original string
  for _, orig := range s {
    // for each element of the legal character array
    for _, leg := range legalChars{
      if orig == leg { // check if it is legal
        stripped = stripped + string(orig)
      }
    }
  }

  if s == stripped {
    fmt.Println("Original string contained no illegal characters")
    return ""
  }
  fmt.Println("[?] Stripped to " + string(stripped))
  return string(stripped)
}

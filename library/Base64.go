package wts

import(
  "fmt"
  "encoding/base64"

)


func Base64Main(message string){
  sayHi("Base64 Decode", message)

  b64DecodeMessage(message)
  message = b64stripUndesired(message)
  if message != "" { b64DecodeMessage(message) }
}

func b64DecodeMessage(message string) {
  fmt.Println("[?] Decoding "+message)

  // Try standard b64 decoding first
  decoded, err := base64.StdEncoding.DecodeString(message)
  if err != nil {
    safePrintln("[x] Base64 => " + string(decoded))
  } else {  safePrintln("[√] Base64 => " + string(decoded))}

  //same but without padding
  decoded, err = base64.RawStdEncoding.DecodeString(message)
  if err != nil {
    safePrintln("[x] Unpadded Base64 => "+string(decoded))
  } else {  safePrintln("[√] Unpadded Base64 => "+string(decoded))}

  // Try URL b64 decoding
  decoded, err = base64.URLEncoding.DecodeString(message)
  if err != nil {
    safePrintln("[x] URL Base64 => "+string(decoded))
  } else {  safePrintln("[√] URL Base64 => "+string(decoded))}

  //same but without padding
  decoded, err = base64.RawURLEncoding.DecodeString(message)
  if err != nil {
    safePrintln("[x] Unpadded URL Base64 => "+string(decoded))
  } else { safePrintln("[√] Unpadded URL Base64 => "+string(decoded))}

}

// remove chars that are outside of the characterset
func b64stripUndesired(s string) string{
  // intentionally missing /+= as they can make the string invalid if placed incorrectly
  var legalChars = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`

  stripped := ""

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

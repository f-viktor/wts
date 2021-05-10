package wts

import(
  "fmt"
  "encoding/base64"
  "strings"
)


func Base64Main(message string){
  fmt.Println("\n[+]_Base64_Decode_"+strings.Repeat("_",len(message)))

  b64DecodeMessage(message)
  message = b64stripUndesired(message)
  if message != "" { b64DecodeMessage(message) }
}

func b64DecodeMessage(message string) {
  fmt.Println("[?] Decoding "+message)

  // Try standard b64 decoding first
  decoded, err := base64.StdEncoding.DecodeString(message)
  if err != nil {
    fmt.Printf("[x] Base64 (failed/partial) => %s\n", decoded)
  } else {  fmt.Println("[√] Base64 => "+string(decoded))}

  //same but without padding
  decoded, err = base64.RawStdEncoding.DecodeString(message)
  if err != nil {
    fmt.Printf("[x] Unpadded Base64 (failed/partial) => %s\n", decoded)
  } else {  fmt.Println("[√] Unpadded Base64 => "+string(decoded))}

  // Try URL b64 decoding
  decoded, err = base64.URLEncoding.DecodeString(message)
  if err != nil {
    fmt.Printf("[x] URL Base64 (failed/partial) => %s\n", decoded)
  } else {  fmt.Println("[√] URL Base64 => "+string(decoded))}

  //same but without padding
  decoded, err = base64.RawURLEncoding.DecodeString(message)
  if err != nil {
    fmt.Printf("[x] Unpadded URL Base64 (failed/partial) => %s\n", decoded)
  } else {  fmt.Println("[√] Unpadded URL Base64 => "+string(decoded))}

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

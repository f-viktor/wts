package wts

import(
  "fmt"
  "github.com/akamensky/base58"
  "strings"
)

// USm3fpXnKG5EUBx2ndxBDMPVciP5hGey2Jh4NDv6gmeo1LkMeiKrLJUUBk6Z

func Base58Main(message string){
  fmt.Println("\n[>]_Base58_Decode_"+strings.Repeat("_",len(message)))

  b58DecodeMessage(message)
  message = b58stripUndesired(message)
  if message != "" { b58DecodeMessage(message) }
}

func b58DecodeMessage(message string) {
  fmt.Println("[?] Decoding "+message)

  // there are other abcs eg for bitcoin

  // Try  b58 decode
  decoded, err := base58.Decode(message)
  if err != nil {
    fmt.Printf("[x] Base58 Flickr (failed/partial) => %s\n", string(decoded))
  } else {  fmt.Println("[√] Base58 => "+string(decoded))}

}

// remove chars that are outside of the characterset
func b58stripUndesired(s string) string{
  // intentionally missing /+= as they can make the string invalid if placed incorrectly
  var legalChars = `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`

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

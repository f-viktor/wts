package wts

import(
  "fmt"
	"net/url"
  "strings"
)


func URLmain(message string){
  fmt.Println("\n[+]_URL_Decode_"+strings.Repeat("_",len(message)))

  urlDecodeMessage(message)
  message = urlstripUndesired(message)
  if message != "" { urlDecodeMessage(message) }
}

func urlDecodeMessage(message string) {
  fmt.Println("[?] Decoding "+message)

  // Try standard url decoding first
  decoded, err := url.PathUnescape(message)
  if err != nil {
    fmt.Printf("[x] URL (failed/partial) => %s\n", decoded)
  } else {  fmt.Println("[√] URL => "+decoded)}

}

// insert % chars after every second character
func urlstripUndesired(s string) string{

  fmt.Println("[?] Stripped to " + string(s))
  return string(s)
}

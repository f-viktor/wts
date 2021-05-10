package wts

import(
  "fmt"
	"net/url"
  "strings"
)


func URLMain(message string){
  fmt.Println("\n[+]_URL_Decode_"+strings.Repeat("_",len(message)))

  // two quick ones
  message = urlDecodeMessage(message)
  // the double
  urlDecodeMessage(message)
  message = forceDelimiter(message)
  urlDecodeMessage(message)
}

func urlDecodeMessage(message string) string {
  fmt.Println("[?] Decoding "+message)

  // Try standard url decoding first
  decoded, err := url.PathUnescape(message)
  if err != nil {
    fmt.Printf("[x] URL (failed/partial) => %s\n", decoded)
  } else {  fmt.Println("[√] URL => "+decoded)}

  return decoded
}

// insert % chars after every second character
// if any of the resulting chars are invalid, the whole decode will failed
// could add something that decodes per character
func forceDelimiter(s string) string{
  for i := 0; i < len(s); i += 3 {
    s = s[:i] + "%" + s[i:]
  }

  fmt.Println("[?] Modified to " + string(s))
  return string(s)
}

package wts

import(
  "fmt"
  "strings"
  "html"
)


func HTMLMain(message string){
  fmt.Println("\n[>]_HTML_Decode_"+strings.Repeat("_",len(message)))

  // two quick ones
  message = htmlDecodeMessage(message)
  // the double
  htmlDecodeMessage(message)
  message = htmlForceDelimiter(message)
  htmlDecodeMessage(message)
}

func htmlDecodeMessage(message string) string {
  fmt.Println("[?] Decoding "+message)

  // Try standard url decoding first
  decoded := html.UnescapeString(message)
  fmt.Println("[√] URL => "+decoded)

  return decoded
}

// insert % chars after every second character
// if any of the resulting chars are invalid, the whole decode will failed
// could add something that decodes per character
func htmlForceDelimiter(s string) string{
  for i := 0; i < len(s); i += 6 {
    s = s[:i] + ";&#" + s[i:]
  }
  s=s[1:]

  fmt.Println("[?] Modified to " + string(s))
  return string(s)
}

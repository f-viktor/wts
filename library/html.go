package wts

import(
  "fmt"
  "html"
)


func HTMLMain(message string){
  sayHi("HTML Decode", message)

  // two quick ones
  message = htmlDecodeMessage(message)
  // the double
  htmlDecodeMessage(message)
  message = forceDelimiter(message,3,";&#")
  htmlDecodeMessage(message)
}

func htmlDecodeMessage(message string) string {
  fmt.Println("[?] Decoding "+message)

  // Try standard url decoding first
  decoded := html.UnescapeString(message)
  safePrintln("[√] URL => "+decoded)

  return decoded
}

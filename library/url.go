package wts

import(
  "fmt"
	"net/url"
)


func URLMain(message string){
  sayHi("URL Decode", message)

  // two quick ones
  message = urlDecodeMessage(message)
  // the double
  urlDecodeMessage(message)
  message = forceDelimiter(message,2,"%")
  urlDecodeMessage(message)
}

func urlDecodeMessage(message string) string {
  fmt.Println("[?] Decoding "+message)

  // Try standard url decoding first
  decoded, err := url.PathUnescape(message)
  if err != nil {
    safePrintln("[x] URL (failed/partial) => "+ decoded)
  } else {  safePrintln("[√] URL => "+decoded)}

  return decoded
}

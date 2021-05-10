package wts

//https://github.com/Equim-chan/leb128

import(
  "fmt"
  "strings"
  "errors"
  "ekyu.moe/leb128"
)


func Base128Main(message string){
  fmt.Println("\n[+]_Base128_Decode"+strings.Repeat("_",len(message)))

  b128DecodeMessage(message)
  leb128DecodeMessage(message)
}

func b128DecodeMessage(message string) {
  fmt.Println("[?] Decoding "+message)

  // Try standard b128 decoding first
  decoded, err := b128Decode(message)
  if err != nil {
    fmt.Printf("[x] Base128 (failed/partial) => %s\n", decoded)
  } else {  fmt.Println("[√] Base128 => "+string(decoded))}
}

func leb128DecodeMessage(message string) {

  // Try unsigned LEB128 decoding (actually returns a number)
  udecoded, n := leb128.DecodeUleb128([]byte(message))
  if int(n) < len(message) {
    fmt.Printf("[x] ULEB128 (failed/partial) => %s => %i \n", string(udecoded), udecoded)
  } else {  fmt.Println("[√] ULEB128 => "+string(udecoded))}

  // Try signed LEB128 decoding (actually returns a number)
  sdecoded, n := leb128.DecodeSleb128([]byte(message))
  if int(n) < len(message) {
    fmt.Printf("[x] SLEB128 (failed/partial) => %s => %i \n", string(sdecoded), sdecoded)
  } else {  fmt.Println("[√] SLEB128 => "+string(sdecoded))}
}


// yoinked from: https://pkg.go.dev/go.chromium.org/luci/common/data/base128#Decode
// would be nice if you just made this an importable package

// Decode decodes src into DecodedLen(len(src)) bytes, returning the actual
// number of bytes written to dst.
//
// If Decode encounters invalid input, it returns an error describing the
// failure.

func b128Decode( message string) (string, error) {
  src := []byte(message)
	var dst []byte
	whichByte := uint(1)
	bufByte := byte(0)
	for _, v := range src {
		if (v & 0x80) != 0 {
			return string(dst), errors.New("base128: high bit set in base128 string")
		}
		if whichByte > 1 {
			dst = append(dst, bufByte|(v>>(8-whichByte)))
		}
		bufByte = v << whichByte
		if whichByte == 8 {
			whichByte = 0
		}
		whichByte++
	}
	return string(dst), nil
}

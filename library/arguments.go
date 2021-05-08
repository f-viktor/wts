package wts

import (
   "flag"
   "os"
)
type CLAConfig struct {
	ToBDecoded		string
}

// parse command line arguments
func ParseArgs() CLAConfig {

   toBDecoded  := flag.String("s", "d2hhdCBhbSBpPw==", "jsm2 sessionId of the member")

   help := flag.Bool("h", false, "Display this help text")

   flag.Parse()

   args := CLAConfig{*toBDecoded}

   // if help is requested print help menu
   if *help {
      flag.Usage()
      os.Exit(2)
   }

   return args
}

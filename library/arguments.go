package wts

import (
   "flag"
   "os"
)
type CLAConfig struct {
	ToBDecoded		string
  UnsafePrint   bool
}

// parse command line arguments
func ParseArgs() CLAConfig {

   toBDecoded  := flag.String("s", "d2hhdCBhbSBpPw==", "string to be identified")
   unsafePrint  := flag.Bool("unsafeP", false, "Do not replace linebreaks and unprintables with ¶ (messy)")
   help := flag.Bool("h", false, "Display this help text")

   flag.Parse()

   args := CLAConfig{*toBDecoded, *unsafePrint}

   // if help is requested print help menu
   if *help {
      flag.Usage()
      os.Exit(2)
   }

   return args
}

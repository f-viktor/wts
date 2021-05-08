package main

import (
	"wts/library"
)

//decoderFuncts := make([]interface{}, 3, 3)  could allow multiple types of functs
// Add your custom decoders to this list, to get them invoked
var decoderFuncts = [...]func(string){wts.Base64main, wts.Base58main, wts.Base32main, wts.Base128main, wts.URLmain}
//var hashIdentifierFuncts = [...]func(string){wts.Base64main}
//var statisticFunctions = [...]func(string){wts.Base64main}

func main(){
  args := wts.ParseArgs()


  // run all decoder functions
  for _, fp := range decoderFuncts {
    fp(args.ToBDecoded)
  }
/*
  // run all hash identifiers
  for _, fp := range decoderFuncts {
    fp(args.ToBDecoded)
  }

  // run all statistic functions
  for _, fp := range decoderFuncts {
    fp(args.ToBDecoded)
  }*/
}

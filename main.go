package main

import (
	"wts/library"
)

//decoderFuncts := make([]interface{}, 3, 3)  could allow multiple types of functs
// Add your custom decoders to this list, to get them invoked
var decoderFuncts = [...]func(string){
	wts.RotMain,
	wts.MorseMain,
	wts.BinaryMain,
	wts.DecimalMain,
	wts.HTMLMain,
	wts.URLMain,
	wts.Base32Main,
	wts.Base58Main,
	wts.Base64Main,
	wts.Base128Main,
	wts.FrequencyMain}
//var hashIdentifierFuncts = [...]func(string){wts.Base64Main}
//var statisticFunctions = [...]func(string){wts.Base64Main}

func main(){
  args := wts.ParseArgs()
	wts.SetUnsafePrint(args.UnsafePrint) // replace unprintable with ¶


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

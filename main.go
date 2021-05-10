package main

import (
	"wts/library"
)

//decoderFuncts := make([]interface{}, 3, 3)  could allow multiple types of functs
// Add your custom decoders to this list, to get them invoked
var decoderFuncts = [...]func(string){
	wts.Base64Main,
	wts.Base58Main,
	wts.Base32Main,
	wts.Base128Main,
	wts.URLMain,
	wts.RotMain,
	wts.DecimalMain,
	wts.HTMLMain,
	wts.MorseMain,
	wts.FrequencyMain}
//var hashIdentifierFuncts = [...]func(string){wts.Base64Main}
//var statisticFunctions = [...]func(string){wts.Base64Main}

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

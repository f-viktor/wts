package wts

import(
  //"fmt"
  "strings"
  "strconv"
  "regexp"
  )

// md4,md5, sha1, sha2, sha3, NTLM, LANMMAN, BLAKE2
func HashMain(message string){
  sayHi("Hash Identifier",message)
  message = strings.TrimSpace(message)
  legalCount, legals := hashCountLegal(message)
  printLegalHash(legalCount, legals, message)
  lengthHexMatch(legalCount, len(message))
  cryptMatch(message)
}


// remove chars that are outside of the characterset
func hashCountLegal(s string) (int, []rune) {
  var legalChars = `0123456789abcdefABCDEF`
  legalCount := 0
  var legals []rune

  // for each element of the original string
  for _, orig := range []rune(s) {
    // for each element of the legal character array
    for _, leg := range []rune(legalChars) {
      if orig == leg { // check if it is legal
        legalCount++
        legals = append(legals, orig)
      }
    }
  }

  return legalCount, legals
}

// print the amount of legal characters
func printLegalHash(legalCount int, legals []rune, message string) {
  // typecasting hell
  validPercent := strconv.Itoa(int(float64(legalCount)/float64(len(message))*100))
  outp := "[?] Valid raw hex characters "
  outp += strconv.Itoa(legalCount)+"/"+strconv.Itoa(len(message))
  outp += " (" + validPercent +"%)"
  outp += " | Legal hash = [ "

  for _, c := range legals {
    outp += string(c)
  }
  outp +=  " ]"

  safePrintln(outp)
}


func lengthHexMatch(legals int, fullength int) {
  hashLengths := make(map[int]string)
  hashLengths[16]  = "LM[LanMan]"
  hashLengths[32]  = "MD2|MD4|MD5|NTLM"
  hashLengths[40]  = "SHA0|SHA1"
  hashLengths[46]  = "bcypt(raw)"
  hashLengths[56]  = "BLAKE-224|SHA2-224|SHA2-512/224|SHA3-224"
  hashLengths[64]  = "BLAKE-256|SHA2-256|SHA2-512/256|SHA3-256"
  hashLengths[96]  = "BLAKE-384|SHA2-384|SHA3-384"
  hashLengths[128] = "BLAKE-512|SHA2-512|SHA3-512"
  hashLengths[0]   = "SHA3|Argon2|PBKDF2" //arbitrary length, still should be multiple of 2

  //MD6 Variable, 0<d≤512 bits
  //BLAKE2 up to 64 bytes (BLAKE2b); up to 32 bytes (BLAKE2s); arbitrary (BLAKE2X)
  //BLAKE3 256 bits, arbitrarily extensible
  //NTLMv1 u4-netntlm::kNS:338d08f8e26de93300000000000000000000000000000000:9526fb8c23a90751cdd619b6cea564742e1e4bf33006ba41:cb8086049ec4736c
  //NTMLv2 admin::N46iSNekpT:08ca45b7d7ea58ee:88dcbe4446168966a153a0064958dac6:5c7830315c7830310000000000000b45c67103d07d7b95acd12ffa11230e0000000052920b85f78d013c31cdb3b92f5d765c783030


}




//try to parse https://passlib.readthedocs.io/en/stable/modular_crypt_format.html
func cryptMatch(message string) {
  //map keyed with strings, valued with function pointers
   hashFuncts := map[string]func([]string){
     "1"    : md5_cryptMatch,
     "md5"  : sun_md5_cryptMarch,
     "sha1" : sha1_cryptMatch,
     "2"    : bcyptMatch,
     "3"    : bsd_nthashMatch,
     "5"    : sha256_cryptMatch,
     "6"    : sha512_cryptMatch,
     "P"    : phpassMatch,
     "H"    : phpassMatch,
     "p5k2" : dlitzCta_pbkdf2_sha1Match,
     "scram"  : scramMatch,
     "argon2" : argon2Match,
     "scrypt" : scryptMatch,
     "pbkdf2" : pbkdf2_sha1Match,
     "pbkdf2-sha256" : pbkdf2_sha256Match,
     "pbkdf2-sha512" : pbkdf2_sha512Match,
     "bcrypt-sha256" : bcrypt_sha256Match,
   }

  //(\$[0-z,/,+,.,=]+){2,4} at least two $ signs with something between
  rePW := regexp.MustCompile(`(\$[0-z,/,+,.,=]+){2,4}`)
  if (rePW.Match([]byte(message))) {
    cryptPart := rePW.Find([]byte(message))
    safePrintln("[?] CRYPT-like hash found "+string(cryptPart))
    cryptArr := strings.Split(string(cryptPart),"$")

    // if a key matches call the corresponding function from decoderFuncts
    for key, matchingFunct := range hashFuncts{
      if (strings.Contains(strings.ToUpper(cryptArr[0]),strings.ToUpper(key))) {
        matchingFunct(cryptArr)
      }
    }
  //edge cases
  // des_crypt https://passlib.readthedocs.io/en/stable/lib/passlib.hash.des_crypt.html#format
  } else if len(message) == 13{
    des_cryptMatch(message)
  // bsdi_crypt https://passlib.readthedocs.io/en/stable/lib/passlib.hash.bsdi_crypt.html#format
  } else if len(message) == 13 && strings.HasPrefix(message, "_"){
    bsdi_cryptMatch(message)
  }
}

package wts

import(
//  "fmt"
  )



//$1$01234567$b5lh2mHyD2PdJjFfALlEz1
func md5_cryptMatch(message []string) {
  // bcrypt $2b$[cost]$[22 character salt][31 character hash]

}

func sun_md5_cryptMarch(message []string) { }

func sha1_cryptMatch(message []string) { }

func bcyptMatch(message []string) {
  // bcrypt $2b$[cost]$[22 character salt][31 character hash]

  /*
  $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
  \__/\/ \____________________/\_____________________________/
   Alg Cost      Salt                        Hash


The original specification did not define how to handle non-ASCII character, nor how to handle a null terminator. The specification was revised to specify that when hashing strings:
the string must be UTF-8 encoded
the null terminator must be included
With this change, the version was changed to $2a$[15]

$2x$, $2y$ (June 2011)

In June 2011, a bug was discovered in crypt_blowfish, a PHP implementation of bcrypt. It was mis-handling characters with the 8th bit set.[16] They suggested that system administrators update their existing password database, replacing $2a$ with $2x$, to indicate that those hashes are bad (and need to use the old broken algorithm). They also suggested the idea of having crypt_blowfish emit $2y$ for hashes generated by the fixed algorithm.

Nobody else, including canonical OpenBSD, adopted the idea of 2x/2y. This version marker change was limited to crypt_blowfish.

$2b$ (February 2014)

A bug was discovered in the OpenBSD implementation of bcrypt. They were storing the length of their strings in an unsigned char (i.e. 8-bit Byte).[15] If a password was longer than 255 characters, it would overflow and wrap at 255.[17]

bcrypt was created for OpenBSD. When they had a bug in their library, they decided to bump the version number.
   */
}

func bsd_nthashMatch(message []string) { }
func sha256_cryptMatch(message []string) { }
func sha512_cryptMatch(message []string) { }
func phpassMatch(message []string) { }
func dlitzCta_pbkdf2_sha1Match(message []string) { }
func scramMatch(message []string) { }
func argon2Match(message []string) { }
func scryptMatch(message []string) { }
func pbkdf2_sha1Match(message []string) { }
func pbkdf2_sha256Match(message []string) { }
func pbkdf2_sha512Match(message []string) { }
func bcrypt_sha256Match(message []string) { }

func des_cryptMatch(message string) { }
func bsdi_cryptMatch(message string) { }
func ntlmMatch(message string) { }

package crypticmysql

import (
	"crypto/aes"
	"crypto/block"
	"bytes"
	"io"
)


/* MySQL (and ruby) both default to padding the 
   plain string up to the next block size with the
   byte value of the number of bytes to be padded
   -- what kind of crap is that??
   so if a block is 16 bytes
   and the string is 11 bytes
   there are 5 bytes of padding, which will be filled
   with \x05
*/
func Aes128EbcEncrypt(in, key []byte) []byte {

	blocks := len(in) / 16

	if len(in)%16 > 0 {
		blocks += 1
	}

	var inbytes = make([]byte, (16 * blocks))

	padLength := (16 * blocks) - len(in)

	for x := 0; x < len(inbytes); x++ {
		inbytes[x] = byte(padLength)
	}
	copy(inbytes, in)

	var crypted bytes.Buffer
	var r io.Reader = bytes.NewBuffer(inbytes)
	cipher, _ := aes.NewCipher(key)
	w := block.NewECBEncrypter(cipher, &crypted)
	io.Copy(w, r)

	cb := crypted.Bytes()

	return cb
}

func Aes128EbcDecrypt(in, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)

	var r io.Reader = block.NewECBDecrypter(cipher, bytes.NewBuffer(in))
	var w io.Writer

	var plain bytes.Buffer
	w = &plain
	io.Copy(w, r)
	return bytes.TrimSpace(plain.Bytes())

}
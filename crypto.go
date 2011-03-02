package crypticmysql

import (
	"crypto/aes"
	"crypto/block"
	"bytes"

)

const aesKeyLen = 128

func aesKey(key []byte) []byte {
	const keyLen = aesKeyLen / 8

	if len(key) == keyLen {
		return key
	}

	k := make([]byte, keyLen)
	copy(k, key)
	for i := keyLen; i < len(key); {
		for j := 0; j < keyLen && i < len(key); j, i = j+1, i+1 {
			k[j] ^= key[i]
		}
	}
	return k
}

// Encrypt using MySQL AES Encrypt
func AESEncrypt(in, key []byte) []byte {
	cipher, err := aes.NewCipher(aesKey(key))
	if err != nil {
		return nil
	}

	pBlocks := (len(in) + aes.BlockSize) / aes.BlockSize
	pText := make([]byte, pBlocks*aes.BlockSize)
	copy(pText, in)
	padLen := byte(len(pText) - len(in))
	for i := len(in); i < len(pText); i++ {
		pText[i] = padLen
	}
	pBuf := bytes.NewBuffer(pText)

	cBuf := bytes.NewBuffer(make([]byte, 0, len(pText)))
	cWtr := block.NewECBEncrypter(cipher, cBuf)
	_, err = pBuf.WriteTo(cWtr)
	if err != nil {
		return nil
	}

	return cBuf.Bytes()
}

// Decrypt using MySQL AES Decrypt
func AESDecrypt(in, key []byte) []byte {
	cipher, err := aes.NewCipher(aesKey(key))
	if err != nil {
		return nil
	}

	pBuf := bytes.NewBuffer(make([]byte, 0, len(in)))

	cRdr := block.NewECBDecrypter(cipher, bytes.NewBuffer(in))
	_, err = pBuf.ReadFrom(cRdr)
	if err != nil {
		return nil
	}

	pText := pBuf.Bytes()
	if len(pText) < 1 {
		return nil
	}
	lenpt := len(pText) - int(pText[len(pText)-1])
	if lenpt < 0 {
		return nil
	}
	return pText[:lenpt]
}

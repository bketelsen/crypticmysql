CrypticMysql is a library that implements AES_DECRYPT and AES_ENCRYPT functions from MySQL in Go


## Prerequisites

* Go Development Environment


## Database

Mysql's AES_DECRYPT() AND AES_ENCRYPT() functions accept a string and a key. AES only needs a 16 byte key, so 
anything longer than 16 bytes is "wrapped" back into the 16 byte key array using an XOR function.
 

	> select AES_ENCRYPT("brian","abcdefghijklmnop");
	> y??doC?T?T.?#r?

 
## Replicating this in Go

	package main
	import (
		github.com/bketelsen/crypticmysql
	)
	func main(){
		cryptedText := crypticmysql.AESEncrypt([]byte("brian"),[]byte("abcdefghijklmnop"))	
		decryptedText := crypticmysql.AESDecrypt(cryptedText, []byte("abcdefghijklmnop"))
	}


## TODO

* Base64 function to encode & decode base64 values if needed
* <DONE> That goofy wrap-around key function that mysql uses. <PeterGo>
* Learn Markdown so this doc won't look so bad

## Contributors:

* Brian Ketelsen
* peterGo <go.peter.90@gmail.com>
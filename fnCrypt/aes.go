package fnCrypt

import (
	"crypto/aes"
	"crypto/cipher"
	rand "crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/d3v-friends/go-pure/fnParams"
	"io"
	mRand "math/rand"
	"time"
)

func EncryptAES256(secret, str string) (res string, err error) {
	// 암호화
	var byteSecret = []byte(secret)
	if len(byteSecret) != 32 {
		err = fmt.Errorf("invalid secret key: func=EncryptAes256, secret=%s", byteSecret)
	}

	var block cipher.Block
	if block, err = aes.NewCipher(byteSecret); err != nil {
		return
	}

	var gcm cipher.AEAD
	if gcm, err = cipher.NewGCM(block); err != nil {
		return
	}

	var nonce = make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	var byteStr = []byte(str)
	var ciphertext = gcm.Seal(nonce, nonce, byteStr, nil)
	res = base64.StdEncoding.EncodeToString(ciphertext)
	return
}

func DecryptAES256(secret, str string) (res string, err error) {
	var byteSecret = []byte(secret)
	if len(byteSecret) != 32 {
		err = fmt.Errorf("invalid secret key: func=EncryptAes256, secret=%s", byteSecret)
	}

	var block cipher.Block
	if block, err = aes.NewCipher(byteSecret); err != nil {
		return
	}

	var gcm cipher.AEAD
	if gcm, err = cipher.NewGCM(block); err != nil {
		return
	}

	var byteStr []byte
	if byteStr, err = base64.StdEncoding.DecodeString(str); err != nil {
		return
	}

	var nonceSize = gcm.NonceSize()
	var nonce, pureCiphertext = byteStr[:nonceSize], byteStr[nonceSize:]

	if byteStr, err = gcm.Open(nil, nonce, pureCiphertext, nil); err != nil {
		return
	}

	res = string(byteStr)
	return
}

func NewSecret(lengths ...int) (res string) {
	var r = mRand.New(mRand.NewSource(time.Now().UnixNano()))
	//var st = 33
	//var ed = 126

	var length = fnParams.Get(lengths)
	if length == 0 {
		length = 32
	}

	res = ""
	for i := 0; i < length; i++ {
		var keyCode = byte(r.Intn(93) + 33)
		res = fmt.Sprintf("%s%c", res, keyCode)
	}

	return
}

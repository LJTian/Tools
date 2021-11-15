package des

import (
	"bytes"
	"errors"
	"github.com/forgoer/openssl"
)

/*
	对key做了预处理，key 为 8位，16位，24位，都可以处理
	默认情况下试一下对应关系
	8位： key1 = key2 = key3
	16位： key1 = key3 = key[:8] key2 = [9:]
	24: key1 = key[:8] key2 = key[9:16] key3 = key[17:]
*/

// 校验并补填key的值
func checkKey(key []byte) (outKey []byte, err error) {

	var des3Key bytes.Buffer

	switch len(key) {
	case 8:
		des3Key.Write(key[:])
		des3Key.Write(key[:])
		des3Key.Write(key[:])

	case 16:
		des3Key.Write(key[:])
		des3Key.Write(key[:8])

	case 24:
		des3Key.Write(key[:])

	default:
		outKey = des3Key.Bytes()
		err = errors.New("The length of the key must be 8, 16, or 24.")
	}
	outKey = des3Key.Bytes()

	return
}

// 3des 解密
func Des3ECBEncrypt(srcMsg []byte, key []byte) (outMsg []byte, err error) {

	// 校验key的值
	var des3Key []byte
	if des3Key, err = checkKey(key); err != nil {
		return
	}

	outMsg, err = openssl.Des3ECBEncrypt(srcMsg, des3Key, openssl.ZEROS_PADDING)
	return
}

// 3des 解密
func Des3ECBDecrypt(srcMsg []byte, key []byte) (outMsg []byte, err error) {

	var des3Key []byte
	if des3Key, err = checkKey(key); err != nil {
		return
	}

	// 校验key的值
	outMsg, err = openssl.Des3ECBDecrypt(srcMsg, des3Key, openssl.ZEROS_PADDING)
	return
}

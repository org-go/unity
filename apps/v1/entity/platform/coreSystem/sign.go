package coreSystem

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)





/**
 * $(aesEncryptECB)
 * @Description:
 * @receiver this
 * @param origDataStr
 * @param keyStr
 * @return hexEncrypted
 */
func (this core) aesEncryptECB(origDataStr, keyStr string) (hexEncrypted string) {
	origData := []byte(origDataStr)
	key := []byte(keyStr)
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return hex.EncodeToString(encrypted)
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}



/**
 * $(_encode)
 * @Description:
 * @receiver this
 * @param secretKey
 * @return string
 */
func (this core) _encode(secretKey string) string {

	h := md5.New()
	h.Write([]byte(secretKey))
	result := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("%s\n", result)
	return result
}

package signx

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	"github.com/buger/jsonparser"
	"golang.org/x/crypto/pkcs12"
)

/**
 * $(GetSign)
 * @Description: 获取签名
 * @param req
 * @param key
 * @return string
 */
func GetSign(req interface{}, keys string) string {

	m := make(map[string]interface{})
	tmp := make(map[string]interface{})
	jsonBody, _ := json.Marshal(req)
	_ = json.Unmarshal(jsonBody, &m)
	for k, val := range m {
		if len(keys) > 20 && k == `BizCode`{
			tmp[`bizCode`] = val
		} else {
			tmp[k] = val
		}
	}

	jsonBody, _ = json.Marshal(tmp)
	var signStr string
	_ = jsonparser.ObjectEach(jsonBody, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if string(key) == "sign" {
			return nil
		}
		ks := string(key)
		/*if len(value) < 1 {
			return nil
		}*/
		if len(keys) > 20 && ks == `bizCode`{
			ks = `BizCode`
		}
		signStr += fmt.Sprintf("%s=%v&", ks, string(value))
		return nil
	})
	// {amount=5&BizCode=test1&originRequestId=JHSTESTTOOL202211151642493239432&reqTime=20221115164422&storeCode=shop1B2BJHSTESTTOOL202211151644223880591}
	// {amount=5&BizCode=test1&originRequestId=JHSTESTTOOL202211151642493239432&receiptNo=&reqTime=20221115164422&storeCode=shop1&termCode=B2BJHSTESTTOOL202211151644223880591}
	//signStr = signStr + "key=" + key
	//signStr = signStr + "B2B" + key
	signStr = fmt.Sprintf(`{%s}`, strings.TrimRight(signStr, `&`) + keys)
	// 1. pkg.SIGN: consumerId=8823936075282&requestId=100000B2B100000		= A215C34B53D2B55A946117A1E33EEA96
	s := md5Encode(signStr)
	//s := ComputeHmacSha256(signStr, key)
	sign := strings.ToUpper(s)
	//fmt.Println("pkg.SIGN:", sign)
	return sign
}



func AesEncryptECB(origDataStr, keyStr string) (hexEncrypted string) {
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

// HMACSHA1 HMACSHA1
//  keyStr 密钥
//  value  消息内容
func HMACSHA1(keyStr, value string) string {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(value))
	sha1 := fmt.Sprintf("%x", mac.Sum(nil))
	//进行url编码
	res := url.QueryEscape(sha1)
	return res
}

// ComputeHmacSha256 sha256加密
func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func GetPrivateKey(privateKeyName, privatePassword string) (*rsa.PrivateKey, error) {
	f, err := os.Open(privateKeyName)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	// 因为pfx证书公钥和密钥是成对的，所以要先转成pem.Block
	blocks, err := pkcs12.ToPEM(bytes, privatePassword)
	if err != nil {
		return nil, err
	}
	if len(blocks) != 2 {
		return nil, errors.New("解密错误")
	}
	// 拿到第一个block，用x509解析出私钥（当然公钥也是可以的）
	privateKey, err := x509.ParsePKCS1PrivateKey(blocks[0].Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

/**
 * $(todo)
 * @Description:
 * @param secretKey
 * @return string
 */
func md5Encode(secretKey string) string {

	h := md5.New()
	h.Write([]byte(secretKey))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}

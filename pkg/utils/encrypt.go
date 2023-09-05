package utils

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

var Encrypt *Encryption 

// AES 加密
type Encryption struct {
	key string
}

func init() {
	Encrypt = newEncryption()
}

func newEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd 填充密码长度
// 这个函数用于填充密码的长度，以满足特定的块大小要求。
// 输入参数：
//   - srcByte []byte：要填充的密码字节切片。
//   - blockSize int：目标块大小，通常是加密算法的分组大小。
//
// 返回值：
//   - []byte：填充后的密码字节切片。
func PadPwd(srcByte []byte, blockSize int) []byte {
	// 计算需要填充的字节数，即目标块大小减去密码长度与目标块大小取模的余数
	padNum := blockSize - len(srcByte)%blockSize

	// 创建一个长度为 padNum，每个字节都是 padNum 的切片 ret
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)

	// 将填充切片 ret 追加到密码切片 srcByte 的末尾
	srcByte = append(srcByte, ret...)

	// 返回填充后的密码切片
	return srcByte
}

// AesEncoding 加密
// 这个函数用于使用AES算法对输入字符串进行加密。
// 输入参数：
//   - k *Encryption：包含加密密钥的 Encryption 结构体。
//   - src string：要加密的原始字符串。
//
// 返回值：
//   - string：加密后的字符串。
func (k *Encryption) AesEncoding(src string) string {
	// 将输入字符串转换为字节切片
	srcByte := []byte(src)

	// 创建 AES 加密块，使用 k.key 作为密钥
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		// 处理错误情况，例如密钥无效时返回空字符串
		return ""
	}

	// 使用密码填充函数 PadPwd 填充原始字节切片
	NewSrcByte := PadPwd(srcByte, block.BlockSize())

	// 创建一个目标字节切片，用于存储加密后的数据
	dst := make([]byte, len(NewSrcByte))

	// 使用 AES 加密块对填充后的字节进行加密，并将结果存储在 dst 中
	block.Encrypt(dst, NewSrcByte)

	// 对加密后的字节进行 base64 编码
	pwd := base64.StdEncoding.EncodeToString(dst)

	// 返回加密后的字符串
	return pwd
}

// UnPadPwd 去掉填充的部分
// 这个函数用于从加密后的数据中去掉填充的字节，以还原原始数据。
// 输入参数：
//   - dst []byte：包含填充数据的字节切片。
// 返回值：
//   - []byte：去除填充后的原始数据字节切片。
//   - error：如果输入数据长度不正确，则返回错误。

func UnPadPwd(dst []byte) ([]byte, error) {
	// 检查输入数据长度是否小于等于0，如果是，返回错误
	if len(dst) <= 0 {
		return nil, errors.New("长度有误")
	}

	// 获取填充的长度
	unpadNum := int(dst[len(dst)-1])

	// 如果填充长度大于数据长度，返回错误
	if len(dst) < unpadNum {
		return nil, errors.New("数据损坏")
	}

	// 去除填充部分，得到原始数据
	str := dst[:(len(dst) - unpadNum)]

	// 返回去除填充后的原始数据
	return str, nil
}

// AesDecoding 解密
// 这个函数用于将经过 AES 加密和 Base64 编码的字符串解密，还原为原始数据。
// 输入参数：
//   - k *Encryption：包含解密密钥的 Encryption 结构体。
//   - pwd string：经过加密和编码的字符串。
//
// 返回值：
//   - string：解密后的原始数据字符串。
func (k *Encryption) AesDecoding(pwd string) string {
	// 将 Base64 编码的字符串转换为字节切片
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		// 处理解码错误，例如字符串不合法时返回空字符串
		return ""
	}

	// 创建 AES 解密块，使用 k.key 作为密钥
	block, errBlock := aes.NewCipher([]byte(k.key))
	if errBlock != nil {
		// 处理创建解密块错误，例如密钥无效时返回空字符串
		return ""
	}

	// 创建一个目标字节切片，用于存储解密后的数据
	dst := make([]byte, len(pwdByte))

	// 使用 AES 解密块对编码后的字节进行解密，并将结果存储在 dst 中
	block.Decrypt(dst, pwdByte)

	// 使用 UnPadPwd 函数去除填充部分
	dst, err = UnPadPwd(dst)
	if err != nil {
		// 处理解除填充错误，例如数据损坏时返回空字符串
		return ""
	}

	// 将解密后的字节转换为字符串，并返回原始数据
	return string(dst)
}

// SetKey 设置密钥
func (k *Encryption) SetKey(key string) {
	k.key = key
}

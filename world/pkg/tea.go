package pkg

import (
	"encoding/hex"
	"strings"
)

// 加密
func EncryptStr(plain, charset, key string) string {
	if plain == "" || charset == "" || key == "" {
		return ""
	}
	arr := encryptInt(toIntArray([]byte(plain), true), toIntArray([]byte(key), false))
	bytes := toByteArray(arr, false)
	return strings.ToUpper(hex.EncodeToString(bytes))
}

// 解密
func DecryptStr(hexStr, charset, key string) string {
	if hexStr == "" || charset == "" || key == "" {
		return ""
	}
	decodedByteArray, _ := hex.DecodeString(hexStr)
	arr := decryptInt(toIntArray(decodedByteArray, false), toIntArray([]byte(key), false))
	bytes := toByteArray(arr, true)
	return string(bytes)
}

func encryptInt(v []int32, k []int32) []int32 {
	n := len(v) - 1
	if n < 1 {
		return v
	}
	if len(k) < 4 {
		key := make([]int32, 4)
		copy(key, k[0:])
		k = key
	}
	var sum int32 = 0
	var e int32
	var z = v[n]
	var y = v[0]
	var delta = 0x9E3779B9
	var q = 6 + 52/(n+1)
	var p int

	for q > 0 {
		q--
		sum = sum + int32(delta)
		e = int32(uint32(sum) >> 2 & 3)
		for p = 0; p < n; p++ {
			y = v[p+1]
			temp := (((int32(uint32(z)>>5) ^ (y << 2)) + (int32(uint32(y)>>3) ^ z<<4)) ^ ((int32(sum) ^ y) + (k[int32(p)&3^e] ^ z)))
			v[p] = v[p] + temp
			z = v[p]
		}
		y = v[0]
		temp := (((int32(uint32(z)>>5) ^ (y << 2)) + (int32(uint32(y)>>3) ^ z<<4)) ^ ((int32(sum) ^ y) + (k[int32(p)&3^e] ^ z)))
		v[n] = v[n] + temp
		z = v[n]
	}
	return v
}

func decryptInt(v []int32, k []int32) []int32 {
	n := len(v) - 1
	if n < 1 {
		return v
	}
	if len(k) < 4 {
		key := make([]int32, 4)
		copy(key, k[0:])
		k = key
	}
	var z = v[n]
	var y = v[0]
	var delta = 0x9E3779B9
	var sum = 0
	var e int32
	var q = 6 + 52/(n+1)
	var p int

	sum = q * delta
	for sum != 0 {
		e = int32(uint32(sum) >> 2 & 3)
		for p = n; p > 0; p-- {
			z = v[p-1]
			temp := (((int32(uint32(z)>>5) ^ (y << 2)) + (int32(uint32(y)>>3) ^ z<<4)) ^ ((int32(sum) ^ y) + (k[int32(p)&3^e] ^ z)))
			v[p] = v[p] - temp
			y = v[p]
		}
		z = v[n]
		temp := (((int32(uint32(z)>>5) ^ (y << 2)) + (int32(uint32(y)>>3) ^ z<<4)) ^ ((int32(sum) ^ y) + (k[int32(p)&3^e] ^ z)))
		v[0] = v[0] - temp
		y = v[0]
		sum = sum - delta
	}
	return v

}

func toByteArray(data []int32, includeLength bool) []byte {
	var n int
	if includeLength {
		n = int(data[len(data)-1])
	} else {
		n = len(data) << 2
	}

	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = (byte)(uint(data[uint(i)>>2]) >> ((i & 3) << 3))
	}
	return result
}

func toIntArray(data []byte, includeLength bool) []int32 {
	length := len(data)
	var n = (uint32(length) >> 2) + 1
	if (length & 3) == 0 {
		n = uint32(length) >> 2
	}

	result := make([]int32, n)
	if includeLength {
		result = make([]int32, n+1)
		result[n] = int32(length)
	}

	for i := 0; i < length; i++ {
		index := uint32(i) >> 2
		result[index] = (result[index] | (int32(0x000000ff&data[i]) << ((i & 3) << 3)))
	}

	return result
}

package main

import (
	"fmt"
	"strings"
	"world/pkg"
)

func main() {
	// str := "FEF0"
	// decodedByteArray, err := hex.DecodeString(str)
	// fmt.Println("byte:", decodedByteArray, "err:", err)
	// encodedString := hex.EncodeToString(decodedByteArray)
	// fmt.Println(encodedString)

	// var arr = []int{0, 1, 2, 3, 4, 20}
	// result := toByteArray(arr, true)
	// fmt.Println(len(result), result)

	// result2 := toIntArray(result, true)
	// fmt.Println(len(result2), result2)

	// k := []int{0, 1, 2}
	// fmt.Println(k)
	// key := make([]int, 4)
	// copy(key, k[0:len(k)])
	// k = key
	// fmt.Println(k)

	// var a int = 100
	// var b int = 200
	// a = a ^ b
	// b = b ^ a
	// a = a ^ b
	// fmt.Println(a, b)
	// var val3 int32 = 2147483647
	// fmt.Println(val3 + 1 + 1)

	// val := uint(7) >> 5
	// val2 := (2654435665 << 4)
	// fmt.Println(val)
	// fmt.Println(val2)

	// k := []int32{0, 1, 2}
	// v := []int32{5, 6, 7}
	// result1 := decryptInt(v, k)
	// fmt.Println(result1)

	str0 := "EDF57ECAE2E1CE51F182285FA074C17AAABAF1B41BB797D5098A2E6824CA915FA1E40096"

	jsonStr := "{\"name\":\"Pankaj Kumar\",\"age\":32}"
	jsonStr = "ddxs"
	str := pkg.EncryptStr(jsonStr, "utf-8", "key")
	str1 := strings.ToUpper(str)
	fmt.Println(str)
	fmt.Println(str0 == str1)
	strx := pkg.DecryptStr(str1, "utf-8", "key")
	fmt.Println(strx)

}

func encrypt(buf []byte, key string) []byte {
	return nil
}

func decrypt(buf []byte, key string) []byte {
	return nil
}

func encryptInt(v []int32, k []int32) []int32 {
	n := len(v) - 1
	if n < 1 {
		return v
	}
	if len(k) < 4 {
		key := make([]int32, 4)
		copy(key, k[0:len(k)])
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
		// fmt.Printf("sum:%d, e:%d, k:%v\n", sum, e, k)
		for p = 0; p < n; p++ {
			// fmt.Printf("v[p]:%d, v[p+1]:%d,y:%d, z:%d\n", v[p], v[p+1], y, z)
			y = v[p+1]
			// temp1 := (int32(uint32(z)>>5) ^ (y << 2))
			// temp2_1 := int32(uint32(y) >> 3)
			// temp2_2 := int32(z << 4)
			// temp2_0 := int32(temp2_1 ^ temp2_2)
			// fmt.Printf("temp2_0:%d, y:%d, z:%d, temp2_1:%d,temp2_2:%d \n", temp2_0, y, z, temp2_1, temp2_2)
			// temp2 := int32(int32(uint(y)>>3) ^ z<<4)
			// temp3 := int32(sum) ^ y
			// temp4 := int32(k[int32(p)&3^e] ^ z)
			// temp0 := int32(temp1+temp2) ^ (temp3 + temp4)
			// fmt.Printf("temp0:%d, temp1:%d,temp2:%d, temp3:%d, temp4:%d\n", temp0, temp1, temp2, temp3, temp4)

			// temp101 := (int32(uint32(z)>>5) ^ (y << 2)) + (int32(uint32(y)>>3) ^ z<<4)
			// temp102 := (int32(sum) ^ y) + (k[int32(p)&3^e] ^ z)
			// temp100 := temp101 ^ temp102

			temp := (((int32(uint32(z)>>5) ^ (y << 2)) + (int32(uint32(y)>>3) ^ z<<4)) ^ ((int32(sum) ^ y) + (k[int32(p)&3^e] ^ z)))
			// fmt.Printf("temp:%d, temp0:%d, temp100:%d\n", temp, temp0, temp100)
			// fmt.Printf("v[p]:%d, temp:%d\n", v[p], temp)
			v[p] = v[p] + temp
			z = v[p]
			// fmt.Printf("q:%d, p:%d,sum:%d, e:%d, y:%d , z:%d, v[p]:%d\n", q, p, sum, e, y, z, v[p])
			// break
		}
		// fmt.Printf("[内部] q:%d,p:%d, sum:%d, e:%d, y:%d , z:%d, v:%d,\n", q, p, sum, e, y, z, v)
		y = v[0]
		temp := (((int32(uint32(z)>>5) ^ (y << 2)) + (int32(uint32(y)>>3) ^ z<<4)) ^ ((int32(sum) ^ y) + (k[int32(p)&3^e] ^ z)))
		v[n] = v[n] + temp
		z = v[n]
	}
	// fmt.Printf("[外部] q:%d, sum:%d,  y:%d , z:%d, v:%d,\n", q, sum, y, z, v)
	return v
}

func decryptInt(v []int32, k []int32) []int32 {
	n := len(v) - 1
	if n < 1 {
		return v
	}
	if len(k) < 4 {
		key := make([]int32, 4)
		copy(key, k[0:len(k)])
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

func toByteArray(data []int, includeLength bool) []byte {
	var n int
	if includeLength {
		n = data[len(data)-1]
	} else {
		n = len(data) << 2
	}
	fmt.Println(n)

	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = (byte)(uint(data[uint(i)>>2]) >> ((i & 3) << 3))
	}
	return result
}

func toIntArray(data []byte, includeLength bool) []int {
	length := len(data)
	var n = (uint(length) >> 2) + 1
	if (length & 3) == 0 {
		n = uint(length) >> 2
	}

	result := make([]int, n)
	if includeLength {
		result = make([]int, n+1)
		result[n] = length
	}

	for i := 0; i < length; i++ {
		index := uint(i) >> 2
		result[index] = (result[index] | (int(0x000000ff&data[i]) << ((i & 3) << 3)))
	}

	return result
}

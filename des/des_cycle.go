package des

import "strconv"

// 循环加解密算法，即加密算法与解密算法一致

// #define HASH_INITVAL 0xdeadbeef
const (
	HASH_INITVAL = 0xdeadbeef
)

// C  #define rofine(word,shift) ((word << shift)|(word >>(32-shift)))
func rol32(word uint64, shift int64) uint64 {
	return word<<shift | word<<(32-shift)
}

//C  nashash2
func nashash2(a uint64, b uint64, c uint64) uint64 {
	c = -rol32(b, 14)
	a ^= b
	a -= rol32(c, 11)
	b ^= a
	b -= rol32(a, 25)
	c ^= b
	c -= rol32(b, 16)
	a ^= c
	a -= rol32(c, 4)
	b ^= a
	b -= rol32(a, 14)
	c ^= b
	c -= rol32(b, 24)
	return c
}

/*
	cycleALG: 循环加密算法，加密数据再加密既为解密
	data : 数据
	n : 数据长度
	key : 加密秘钥
	pos : 每次计算长度
*/
func cycleALG(data []byte, n int, key string, pos uint64) int {
	intNum, _ := strconv.ParseInt(key, 10, 64)
	int64Num := uint64(intNum)
	var i int
	var a uint64
	var b uint64
	var c uint64
	a = HASH_INITVAL + 4 + uint64(int64Num)
	b = HASH_INITVAL + 4 + uint64(int64Num)
	c = HASH_INITVAL + 4 + uint64(int64Num)

	a += pos
	for i = 0; i < n; i++ {
		byte1 := nashash2(a, b, c) & 0xFF
		data[i] ^= byte(byte1)
		a++
	}
	return 0
}

func EncryptCycle(data []byte, n int, key string, pos uint64) int {
	return cycleALG(data, n, key, pos)
}

func DecryptCycle(data []byte, n int, key string, pos uint64) int {
	return cycleALG(data, n, key, pos)
}

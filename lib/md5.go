package lib

import (
	"encoding/binary"
	"math/bits"
)

func MD5(data []byte) [16]byte {
	s := [64]uint32{
		7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
	}

	K := [64]uint32{
		0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
		0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
		0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
		0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
		0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
		0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
		0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
		0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
		0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
		0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
		0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
		0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
		0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
		0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
		0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
		0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
	}

	var a0 uint32 = 0x67452301 // A
	var b0 uint32 = 0xefcdab89 // B
	var c0 uint32 = 0x98badcfe // C
	var d0 uint32 = 0x10325476 // D

	// In implementations that only work with complete bytes:
	// append 0x80
	// pad with 0x00 bytes so that the message length in bytes ≡ 56 (mod 64).

	// append "1" bit to data
	tmp := [1 + 63 + 8]byte{0x80}
	// calculate number of padding bytes
	pad := (55 - len(data)) % 64
	// append length in bits
	binary.LittleEndian.PutUint64(tmp[1+pad:], uint64(len(data))<<3)

	// udpate data
	data = append(data, tmp[:1+pad+8]...)

	// Process the message in successive 512-bit chunks:
	for i := 0; i < len(data); i += 64 {
		// break chunk into sixteen 32-bit words M[j], 0 ≤ j ≤ 15
		chunk := data[i : i+64]
		var words [16]uint32
		for j := 0; j < 16; j++ {
			words[j] = binary.LittleEndian.Uint32(chunk[j*4 : (j+1)*4])
		}

		// Initialize hash value for this chunk:
		A := a0
		B := b0
		C := c0
		D := d0

		// Main loop:
		var j uint32
		for j = 0; j < 64; j++ {
			var F, g uint32
			if j < 16 {
				F = (B & C) | ((^B) & D)
				g = j
			} else if j < 32 {
				F = (D & B) | ((^D) & C)
				g = (5*j + 1) % 16
			} else if j < 48 {
				F = B ^ C ^ D
				g = (3*j + 5) % 16
			} else {
				F = C ^ (B | (^D))
				g = (7 * j) % 16
			}
			F = F + A + K[j] + words[g]
			A = D
			D = C
			C = B
			B = B + bits.RotateLeft32(F, int(s[j]))
		}

		// Add this chunk's hash to result so far:
		a0 = a0 + A
		b0 = b0 + B
		c0 = c0 + C
		d0 = d0 + D
	}

	var digest [16]byte
	binary.LittleEndian.PutUint32(digest[0:], a0)
	binary.LittleEndian.PutUint32(digest[4:], b0)
	binary.LittleEndian.PutUint32(digest[8:], c0)
	binary.LittleEndian.PutUint32(digest[12:], d0)

	return digest
}

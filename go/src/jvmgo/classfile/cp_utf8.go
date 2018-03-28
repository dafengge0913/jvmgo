package classfile

import (
	"fmt"
	"unicode/utf16"
)

// CONSTANT_Utf8_info {
//     u1 tag;
//     u2 length;
//     u1 bytes[length];
// }

type ConstantUtf8Info struct {
	str string
}

func (info *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	info.str = decodeMUTF8(bytes)
}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(byteArr []byte) string {
	utfLen := len(byteArr)
	charArr := make([]uint16, utfLen)

	var c, char2, char3 uint16
	count := 0
	charArrCount := 0

	// 如果UTF-16编码的字符 其值小于等于0x7F(127)的话 则MUTF-8直接用一个字节对其编码 这时 MUTF-8编码是完全和ASCII码兼容的
	for count < utfLen {
		c = uint16(byteArr[count])
		if c > 127 {
			break
		}
		count++
		charArr[charArrCount] = c
		charArrCount++
	}

	for count < utfLen {
		c = uint16(byteArr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			charArr[charArrCount] = c
			charArrCount++
		case 12, 13:
			// UTF-16编码字符的数值范围在0x0080~0x07FF之间的情况 当然还要包括0x0这种情况。
			// 在这些情况下 MUTF-8编码将使用两个字节 对于第一个字节 前三个比特位是110 后面的5个比特位用来存放UTF-16编码字符数值的高5位
			// 对于第二个字节 前两个比特位是10 后面6个比特位用来存放UTF-16编码字符数值的低6位 对于数值为0x0的这种特殊情况 其MUTF-8编码后的值为0xC0和0x80
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(byteArr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			charArr[charArrCount] = c&0x1F<<6 | char2&0x3F
			charArrCount++

		case 14:
			// 如果UTF-16编码字符的数值范围在0x0800~0xFFFF之间的话 MUTF-8将使用三个字节对其进行编码
			// 对于第一个字节 前四个比特位是1110 后面的4个比特位用来存放UTF-16编码字符数值的高4位
			// 对于第二个字节 前两个比特位是10 后面6个比特位用来存放UTF-16编码字符数值的中间6位
			// 对于第三个字节 前两个比特位仍然是10 后面6个比特位用来存放UTF-16编码字符数值的最低6位
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(byteArr[count-2])
			char3 = uint16(byteArr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count-1))
			}
			charArr[charArrCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charArrCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utfLen
	charArr = charArr[0:charArrCount]
	runes := utf16.Decode(charArr)
	return string(runes)
}

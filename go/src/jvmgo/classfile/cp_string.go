package classfile

// CONSTANT_String_info常量表示java.lang.String字面量
// 本身并不存放字符串数据 只存了常量池索引 这个索引指向一个CONSTANT_Utf8_info常量
//
// CONSTANT_String_info {
//     u1 tag;
//     u2 string_index;
// }

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (info *ConstantStringInfo) readInfo(reader *ClassReader) {
	info.stringIndex = reader.readUint16()
}

func (info *ConstantStringInfo) String() string {
	return info.cp.getUtf8(info.stringIndex)
}

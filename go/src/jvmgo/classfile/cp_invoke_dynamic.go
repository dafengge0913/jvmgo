package classfile

// CONSTANT_MethodHandle_info {
//     u1 tag;
//     u1 reference_kind;
//     u2 reference_index;
// }
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (info *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	info.referenceKind = reader.readUint8()
	info.referenceIndex = reader.readUint16()
}



package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (info *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	info.nameIndex = reader.readUint16()
	info.descriptorIndex = reader.readUint16()
}

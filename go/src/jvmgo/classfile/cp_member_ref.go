package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (info *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	info.classIndex = reader.readUint16()
	info.nameAndTypeIndex = reader.readUint16()
}

func (info *ConstantMemberRefInfo) ClassName() string {
	return info.cp.getClassName(info.classIndex)
}

func (info *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return info.cp.getNameAndType(info.nameAndTypeIndex)
}

// CONSTANT_Fieldref_info {
//     u1 tag;
//     u2 class_index;
//     u2 name_and_type_index;
// }
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

// CONSTANT_Methodref_info {
//     u1 tag;
//     u2 class_index;
//     u2 name_and_type_index;
// }
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

// CONSTANT_InterfaceMethodref_info {
//     u1 tag;
//     u2 class_index;
//     u2 name_and_type_index;
// }
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

package classfile

// CONSTANT_Class_info {
//     u1 tag;
//     u2 name_index;
// }

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (info *ConstantClassInfo) readInfo(reader *ClassReader) {
	info.nameIndex = reader.readUint16()
}

func (info *ConstantClassInfo) Name() string {
	return info.cp.getUtf8(info.nameIndex)
}

// CONSTANT_MethodType_info {
//     u1 tag;
//     u2 descriptor_index;
// }

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (info *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	info.descriptorIndex = reader.readUint16()
}

// CONSTANT_InvokeDynamic_info {
//     u1 tag;
//     u2 bootstrap_method_attr_index;
//     u2 name_and_type_index;
// }

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (info *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	info.bootstrapMethodAttrIndex = reader.readUint16()
	info.nameAndTypeIndex = reader.readUint16()
}

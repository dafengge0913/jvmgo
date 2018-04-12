package classfile

// 字段结构定义
// field_info {
//     u2 				access_flags;
//     u2 				name_index;
//     u2 				descriptor_index;
//     u2 				attributes_count;
//     attribute_info 	attributes[attributes_count];
// }

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlags
}

func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

// (1) 基本类型 boolean byte short char int long float和double的描述符是单个字母
//     分别对应 Z B S C I J F和D 注意 long的描述符是J而不是L
// (2) 引用类型的描述符是L＋类的完全限定名＋分号
// (3) 数组类型的描述符是[＋数组元素类型描述符
// (4) 字段描述符就是字段类型的描述符
// (5) 方法描述符是 参数类型的域描述符按照申明顺序放入一对括号中后+返回值类型描述符
//     其中void返回值由单个字母V表示。
func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}

func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (mi *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

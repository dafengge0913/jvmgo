package classfile

// SourceFile_attribute {
//     u2 attribute_name_index;
//     u4 attribute_length;
//     u2 sourcefile_index;
// }

type SourceFileAttribute struct {
	cp              ConstantPool
	SourceFileIndex uint16
}

func (attr *SourceFileAttribute) readInfo(reader *ClassReader) {
	attr.SourceFileIndex = reader.readUint16()
}

func (attr *SourceFileAttribute) FileName() string {
	return attr.cp.getUtf8(attr.SourceFileIndex)
}

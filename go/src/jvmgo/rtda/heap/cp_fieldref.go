package heap

import "jvmgo/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (ref *FieldRef) ResolvedField() *Field {
	if ref.field == nil {
		ref.resolveFieldRef()
	}
	return ref.field
}

// 类D想通过字段符号引用访问类C的某个字段 首先要解
// 析符号引用得到类C 然后根据字段名和描述符查找字段
func (ref *FieldRef) resolveFieldRef() {
	d := ref.cp.class
	c := ref.ResolvedClass()
	field := lookupField(c, ref.name, ref.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	ref.field = field
}

func lookupField(class *Class, name string, descriptor string) *Field {
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, inter := range class.interfaces {
		if field := lookupField(inter, name, descriptor); field != nil {
			return field
		}
	}
	for class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}
	return nil
}

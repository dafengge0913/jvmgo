package heap

import "jvmgo/classfile"

// 接口方法符号引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (ref *InterfaceMethodRef) ResolvedInterfaceMethod()*Method  {
	if ref.method == nil {
		ref.resolveInterfaceMethodRef()
	}
	return ref.method
}

func (ref *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := ref.cp.class
	c := ref.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	ref.method = method
}


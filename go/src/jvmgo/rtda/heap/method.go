package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (method *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		method.maxStack = uint(codeAttr.MaxStack())
		method.maxLocals = uint(codeAttr.MaxLocals())
		method.code = codeAttr.Code()
	}
}

func (method *Method) IsSynchronized() bool {
	return 0 != method.accessFlags&ACC_SYNCHRONIZED
}

func (method *Method) IsBridge() bool {
	return 0 != method.accessFlags&ACC_BRIDGE
}

func (method *Method) IsVarargs() bool {
	return 0 != method.accessFlags&ACC_VARARGS
}

func (method *Method) IsNative() bool {
	return 0 != method.accessFlags&ACC_NATIVE
}

func (method *Method) IsAbstract() bool {
	return 0 != method.accessFlags&ACC_ABSTRACT
}

func (method *Method) IsStrict() bool {
	return 0 != method.accessFlags&ACC_STRICT
}

func (method *Method) MaxStack() uint {
	return method.maxStack
}

func (method *Method) MaxLocals() uint {
	return method.maxLocals
}

func (method *Method) Code() []byte {
	return method.code
}

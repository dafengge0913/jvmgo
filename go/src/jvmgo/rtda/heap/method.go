package heap

import (
	"jvmgo/classfile"
)

type Method struct {
	ClassMember
	maxStack        uint
	maxLocals       uint
	code            []byte
	argSlotCount    uint
	exceptionTable  ExceptionTable
	lineNumberTable *classfile.LineNumberTableAttribute
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}
func (method *Method) injectCodeAttribute(returnType string) {
	method.maxStack = 4
	method.maxLocals = method.argSlotCount
	// 虚拟机规范并没有规定如何实现和调用本地方法
	// 使用 0xfe 保留指令 实现本地方法调用
	switch returnType[0] {
	case 'V':
		method.code = []byte{0xfe, 0xb1} // return
	case 'D':
		method.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		method.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		method.code = []byte{0xfe, 0xad} // lreturn
	case 'L', '[':
		method.code = []byte{0xfe, 0xb0} // areturn
	default:
		method.code = []byte{0xfe, 0xac} // ireturn
	}
}

func (method *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		method.maxStack = uint(codeAttr.MaxStack())
		method.maxLocals = uint(codeAttr.MaxLocals())
		method.code = codeAttr.Code()
		method.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), method.class.constantPool)
		method.lineNumberTable = codeAttr.LineNumberTableAttribute()
	}
}

func (method *Method) calArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		method.argSlotCount++
		if paramType == "J" || paramType == "D" {
			method.argSlotCount++
		}
	}
	if !method.IsStatic() {
		// 实例方法 编译器会在参数列表的前面添加一个参数 this引用
		method.argSlotCount++
	}
}

func (method *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := method.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
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

func (method *Method) ArgSlotCount() uint {
	return method.argSlotCount
}

func (method *Method) GetLineNumber(pc int) int {
	if method.IsNative() {
		return -2
	}
	if method.lineNumberTable == nil {
		return -1
	}
	return method.lineNumberTable.GetLineNumber(pc)
}

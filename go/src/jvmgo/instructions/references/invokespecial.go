package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

// 在创建类实例时 编译器会在new指令的后面加入
// invokespecial指令来调用构造函数初始化对象

// 私有方法和构造函数不需要动态绑定 所以invokespecial指令可以加快方法调用速度
// 使用super关键字调用超类中的方法不能使用invokevirtual指令 否则会陷入无限循环
type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (ins *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(ins.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 从操作数栈中弹出this引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoke := resolvedMethod
	if currentClass.IsSuper() && resolvedClass.IsSuperClassOf(currentClass) && resolvedMethod.Name() != "<init>" {
		// 查找最近的超类方法
		methodToBeInvoke = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}
	if methodToBeInvoke == nil || methodToBeInvoke.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoke)
}

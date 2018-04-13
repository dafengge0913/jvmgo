package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

// Check whether object is of given type
type CHECK_CAST struct {
	base.Index16Instruction
}

func (ins *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	// 弹出对象引用 再推回去 这样就不会改变操作数栈的状态
	stack.PushRef(ref)
	if ref == nil {
		// null 引用可以转换成任何类型
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(ins.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}

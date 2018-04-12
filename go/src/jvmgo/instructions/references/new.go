package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

// Create new object
type NEW struct {
	base.Index16Instruction
}

func (ins *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(ins.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

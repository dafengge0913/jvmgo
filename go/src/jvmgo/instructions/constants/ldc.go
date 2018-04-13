package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// ldc系列指令从运行时常量池中加载常量值 并把它推入操作数栈

type LDC struct {
	base.Index8Instruction
}

func (ins *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, ins.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (ins *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, ins.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	default:
		panic("todo: ldc: ")
	}
}

type LDC2_W struct {
	base.Index16Instruction
}

func (ins *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(ins.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

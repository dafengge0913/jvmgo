package conversions

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// Convert long to double
type L2D struct {
	base.NoOperandsInstruction
}

func (ins *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// Convert long to float
type L2F struct {
	base.NoOperandsInstruction
}

func (ins *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// Convert long to int
type L2I struct {
	base.NoOperandsInstruction
}

func (ins *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}

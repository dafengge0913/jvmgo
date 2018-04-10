package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// Negate double
type DNEG struct {
	base.NoOperandsInstruction
}

func (neg *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct {
	base.NoOperandsInstruction
}

func (neg *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct {
	base.NoOperandsInstruction
}

func (neg *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct {
	base.NoOperandsInstruction
}

func (neg *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}

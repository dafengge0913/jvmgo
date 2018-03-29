package loads

import (
	"jvmgo/rtda"
	"jvmgo/instructions/base"
)

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

type DLOAD struct {
	base.Index8Instruction
}

func (ins *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(ins.Index))
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (ins *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (ins *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (ins *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (ins *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

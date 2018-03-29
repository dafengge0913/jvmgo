package loads

import (
	"jvmgo/rtda"
	"jvmgo/instructions/base"
)

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

type ALOAD struct {
	base.Index8Instruction
}

func (ins *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(ins.Index))
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (ins *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (ins *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (ins *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (ins *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

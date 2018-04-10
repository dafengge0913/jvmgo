package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

type FSTORE struct {
	base.Index8Instruction
}

func (ins *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(ins.Index))
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (ins *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (ins *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (ins *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (ins *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

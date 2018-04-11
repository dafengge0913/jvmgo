package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (ins *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, ins.Offset)
}

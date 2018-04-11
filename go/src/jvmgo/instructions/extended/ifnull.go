package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// 根据引用是否是null进行跳转
type IFNULL struct {
	base.BranchInstruction
}

func (ins *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, ins.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (ins *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, ins.Offset)
	}
}

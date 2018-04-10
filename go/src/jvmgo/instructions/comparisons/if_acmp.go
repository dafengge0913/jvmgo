package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// 根据栈顶的两个引用是否相同进行跳转

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2
}

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (ins *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, ins.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (ins *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, ins.Offset)
	}
}

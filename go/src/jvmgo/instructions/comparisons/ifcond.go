package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// if<ins> 指令把操作数栈顶的int变量弹出 然后跟0进行比较 满足条件则跳转
type IFEQ struct {
	base.BranchInstruction
}

func (ins *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, ins.Offset)
	}
}

type IFNE struct {
	base.BranchInstruction
}

func (ins *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, ins.Offset)
	}
}

type IFLT struct {
	base.BranchInstruction
}

func (ins *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, ins.Offset)
	}
}

type IFLE struct {
	base.BranchInstruction
}

func (ins *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, ins.Offset)
	}
}

type IFGT struct {
	base.BranchInstruction
}

func (ins *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, ins.Offset)
	}
}

type IFGE struct {
	base.BranchInstruction
}

func (ins *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, ins.Offset)
	}
}

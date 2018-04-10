package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// 弹出int float等占用一个操作数栈位置的变量
type POP struct {
	base.NoOperandsInstruction
}

func (ins *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// 弹出double long等占用两个操作数栈位置的变量
type POP2 struct {
	base.NoOperandsInstruction
}

func (ins *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

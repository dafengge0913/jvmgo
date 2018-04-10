package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// 交换栈顶变量
type SWAP struct {
	base.NoOperandsInstruction
}

func (ins *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

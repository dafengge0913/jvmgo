package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// 在创建类实例时 编译器会在new指令的后面加入
// invokespecial指令来调用构造函数初始化对象

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (ins *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}

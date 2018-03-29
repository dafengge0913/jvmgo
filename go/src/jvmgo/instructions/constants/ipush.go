package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// 从操作数中获取一个byte型整数 扩展成int型 然后推入栈顶
type BIPUSH struct {
	val int8
}

func (ins *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	ins.val = reader.ReadInt8()
}

func (ins *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(ins.val)
	frame.OperandStack().PushInt(i)
}

// 从操作数中获取一个short型整数 扩展成int型 然后推入栈顶
type SIPUSH struct {
	val int16
}

func (ins *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	ins.val = reader.ReadInt16()
}

func (ins *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(ins.val)
	frame.OperandStack().PushInt(i)
}

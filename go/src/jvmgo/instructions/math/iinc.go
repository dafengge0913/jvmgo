package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// iinc指令给局部变量表中的int变量增加常量值
type IINC struct {
	Index uint  // 局部变量表索引
	Const int32 // 常量值
}

func (iinc *IINC) FetchOperands(reader *base.BytecodeReader) {
	iinc.Index = uint(reader.ReadUint8())
	iinc.Const = int32(reader.ReadInt8())
}

func (iinc *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(iinc.Index)
	val += iinc.Const
	localVars.SetInt(iinc.Index, val)
}

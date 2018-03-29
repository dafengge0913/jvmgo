package base

import "jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (ins *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {}

// 跳转指令
type BranchInstruction struct {
	Offset int
}

func (ins *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	ins.Offset = int(reader.ReadInt16())
}

// 局部变量表索引
type Index8Instruction struct {
	Index uint
}

func (ins *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	ins.Index = uint(reader.ReadUint8())
}

// 常量池索引
type Index16Instruction struct {
	Index uint
}

func (ins *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	ins.Index = uint(reader.ReadUint16())
}

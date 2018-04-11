package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (ins *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ins.defaultOffset = reader.ReadInt32()
	ins.low = reader.ReadInt32()
	ins.high = reader.ReadInt32()
	jumpOffsetsCount := ins.high - ins.low + 1
	ins.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (ins *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= ins.low && index <= ins.high {
		offset = int(ins.jumpOffsets[index-ins.low])
	} else {
		offset = int(ins.defaultOffset)
	}
	base.Branch(frame, offset)
}

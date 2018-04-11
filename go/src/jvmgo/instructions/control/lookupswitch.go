package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (ins *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ins.defaultOffset = reader.ReadInt32()
	ins.npairs = reader.ReadInt32()
	ins.matchOffsets = reader.ReadInt32s(ins.npairs * 2)
}

func (ins *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < ins.npairs*2; i += 2 {
		if ins.matchOffsets[i] == key {
			offset := ins.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(ins.defaultOffset))
}

package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// goto_w指令和goto指令的唯一区别就是索引从2字节变成4字节
type GOTO_W struct {
	offset int
}

func (ins *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	ins.offset = int(reader.ReadInt32())
}

func (ins *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, ins.offset)
}

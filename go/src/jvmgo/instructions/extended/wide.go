package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/instructions/stores"
	"jvmgo/rtda"
	"jvmgo/instructions/loads"
	"jvmgo/instructions/math"
)

// 加载类指令 存储类指令 ret指令和iinc指令需要
// 按索引访问局部变量表 索引以uint8的形式存在字节码中
// 对于大部分方法 局部变量表大小都不会超过256
// 所以用一字节来表示索引就够了
// 如果方法的局部变量表超过限制
// WIDE指令扩展前述指令
type WIDE struct {
	modifiedInstruction base.Instruction
}

func (wide *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x36:
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x37:
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x38:
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x39:
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x3a:
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		wide.modifiedInstruction = inst
	case 0xa9: // ret
		// 暂时没有实现ret指令
		panic("Unsupported opcode: 0xa9!")
	}
}

func (wide *WIDE) Execute(frame *rtda.Frame) {
	wide.modifiedInstruction.Execute(frame)
}

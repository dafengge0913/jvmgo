package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// int 左位移
type ISHL struct {
	base.NoOperandsInstruction
}

func (sh *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt() // 位移数
	v1 := stack.PopInt() // 要进行位移的变量
	// int变量只有32位 v2的前5个比特就足够
	// Go位移操作符右侧必须是无符号整数
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// int算术右位移
type ISHR struct {
	base.NoOperandsInstruction
}

func (sh *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// int逻辑右位移
type IUSHR struct {
	base.NoOperandsInstruction
}

func (sh *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	// Go语言没有>>>运算符 转成无符号整数 位移后 转回有符号整数
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// long左位移
type LSHL struct {
	base.NoOperandsInstruction
}

func (sh *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// long算术右位移
type LSHR struct {
	base.NoOperandsInstruction
}

func (sh *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	// int变量有64位 取v2的前6个比特
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// long逻辑右位移
type LUSHR struct {
	base.NoOperandsInstruction
}

func (sh *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}

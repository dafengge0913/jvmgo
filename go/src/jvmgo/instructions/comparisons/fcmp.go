package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// Compare float

// 浮点数计算可能产生NaN值 无法比较 fcmpg和fcmpl结果不同
func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

type FCMPG struct {
	base.NoOperandsInstruction
}

func (cmp *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (cmp *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

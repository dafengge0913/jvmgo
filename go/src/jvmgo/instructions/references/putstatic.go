package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

// Set static field in class
type PUT_STATIC struct {
	base.Index16Instruction
}

func (ins *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(ins.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		// 如果是final字段则实际操作的是静态常量 只能在类初始化方法中给它赋值
		// 类初始化方法由编译器生成 名字是<clinit>
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	// 根据字段类型从操作数栈中弹出相应的值 然后赋给静态变量
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}

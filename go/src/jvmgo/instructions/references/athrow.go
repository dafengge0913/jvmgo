package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	"reflect"
)

// Throw exception or error
type ATHROW struct {
	base.NoOperandsInstruction
}

func (ins *ATHROW) Execute(frame *rtda.Frame) {
	// 从操作数栈中弹出异常对象引用
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}
	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handlerUncaughtException(thread, ex)
	}
}

func handlerUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	thread.ClearStack()
	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ":" + goMsg)
	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1
		handlerPc := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPc > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPc)
			return true
		}
		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

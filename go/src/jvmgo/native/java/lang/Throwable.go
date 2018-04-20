package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

// private native Throwable fillInStackTrace(int dummy);
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(thisObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	// 跳过栈顶两帧 fillInStackTrace(int) 和 fillInStackTrace()
	skip := distanceToObject(thisObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

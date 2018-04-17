package reserved

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/native"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (ins *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}

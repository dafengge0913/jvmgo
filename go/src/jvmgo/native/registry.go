package native

import "jvmgo/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := getKey(className, methodName, methodDescriptor)
	registry[key] = method
}

func getKey(className, methodName, methodDescriptor string) string {
	return className + "~" + methodName + "~" + methodDescriptor
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := getKey(className, methodName, methodDescriptor)
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {}

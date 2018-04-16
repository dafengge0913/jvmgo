package main

import (
	"jvmgo/rtda"
	"fmt"
	"jvmgo/instructions/base"
	"jvmgo/instructions"
	"jvmgo/rtda/heap"
)

func interpret(method *heap.Method, logIns bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(thread)
	loop(thread, logIns)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtda.Thread, logIns bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		ins := instructions.NewInstruction(opcode)
		ins.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logIns {
			logInstruction(frame, ins)
		}
		// execute
		ins.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func logInstruction(frame *rtda.Frame, ins base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, ins, ins)
}

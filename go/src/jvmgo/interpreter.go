package main

import (
	"jvmgo/rtda"
	"fmt"
	"jvmgo/instructions/base"
	"jvmgo/instructions"
	"jvmgo/rtda/heap"
)

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		ins := instructions.NewInstruction(opcode)
		ins.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// execute
		fmt.Printf("pc:%2d ins:%T %v\n", pc, ins, ins)
		ins.Execute(frame)
	}
}

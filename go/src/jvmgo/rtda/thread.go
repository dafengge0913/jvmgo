package rtda

type Thread struct {
	pc int
	statck *Stack
}

func NewThread() *Thread  {
	return &Thread{
		statck:newStack(1024),
	}
}

func (t *Thread) PushFrame(frame *Frame)  {
	t.statck.push(frame)
}

func (t *Thread) PopFrame()*Frame  {
	return t.statck.pop()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.statck.top()
}
package heap

type Object struct {
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (obj *Object) Class() *Class {
	return obj.class
}

func (obj *Object) Fields() Slots {
	return obj.fields
}

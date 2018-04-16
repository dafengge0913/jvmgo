package heap

type Object struct {
	class *Class
	data  interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (obj *Object) Class() *Class {
	return obj.class
}

func (obj *Object) Fields() Slots {
	return obj.data.(Slots)
}

func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(obj.class)
}

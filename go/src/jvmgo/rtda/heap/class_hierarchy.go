package heap

func (class *Class) IsSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (class *Class) IsImplements(inter *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == inter || i.IsSubInterfaceOf(inter) {
				return true
			}
		}
	}
	return false
}

func (class *Class) IsSubInterfaceOf(inter *Class) bool {
	for _, superInter := range class.interfaces {
		if superInter == inter || superInter.IsSubInterfaceOf(inter) {
			return true
		}
	}
	return false
}

func (class *Class) IsAssignableFrom(other *Class) bool {
	if other == class {
		return true
	}
	if !other.IsArray() {
		if !other.IsInterface() {
			if !class.IsInterface() {
				return other.IsSubClassOf(class)
			} else {
				return other.IsImplements(class)
			}
		} else {
			if !class.IsInterface() {
				return class.isJlObject()
			} else {
				return class.IsSubInterfaceOf(other)
			}
		}
	} else {
		if !class.IsArray() {
			if !class.IsInterface() {
				return class.isJlObject()
			} else {
				return class.isJlCloneable() || class.isJioSerializable()
			}
		} else {
			otherC := other.ComponentClass()
			classC := class.ComponentClass()
			return otherC == classC || classC.IsAssignableFrom(otherC)
		}
	}
	return false
}

func (class *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(class)
}

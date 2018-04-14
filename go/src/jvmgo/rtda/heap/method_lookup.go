package heap

func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func lookupMethodInInterfaces(inters []*Class, name, descriptor string) *Method {
	for _, inter := range inters {
		for _, method := range inter.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterfaces(inter.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}

func lookupInterfaceMethod(inter *Class, name, descriptor string) *Method {
	for _, method := range inter.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(inter.interfaces, name, descriptor)
}

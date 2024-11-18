package app

import "sync"

type Container struct {
	bindings   map[string]func() interface{}
	singletons map[string]func() interface{}
	instances  map[string]interface{}
	mutex      sync.Mutex
}

var App = &Container{
	bindings:   make(map[string]func() interface{}),
	singletons: make(map[string]func() interface{}),
	instances:  make(map[string]interface{}),
}

func (c *Container) Bind(name string, resolver func() interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.bindings[name] = resolver
}

func (c *Container) Singleton(name string, resolver func() interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.singletons[name] = resolver
}

func (c *Container) Resolve(name string) interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the singleton instance already exists
	if instance, ok := c.instances[name]; ok {
		return instance
	}

	// Check if the singleton resolver exists and create instance
	if resolver, ok := c.singletons[name]; ok {
		instance := resolver()
		c.instances[name] = instance
		return instance
	}

	// Check if the binding resolver exists and resolve transient
	if resolver, ok := c.bindings[name]; ok {
		return resolver()
	}

	return nil
}

// ResolveWithoutLock method without locking, mainly for internal use if needed
func (c *Container) ResolveWithoutLock(name string) interface{} {
	if instance, ok := c.instances[name]; ok {
		return instance
	}

	if resolver, ok := c.singletons[name]; ok {
		instance := resolver()
		c.instances[name] = instance
		return instance
	}

	if resolver, ok := c.bindings[name]; ok {
		return resolver()
	}

	return nil
}

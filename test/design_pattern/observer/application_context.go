/**
 * @description observer
 * @author zhangbingbing@baidu.com
 * @date 2019/3/13
 */
package observer

import (
	"gin-app/test/design_pattern/observer/listener"
	"gin-app/test/design_pattern/observer/event"
)

var (
	listenerMap    = map[int]listener.ApplicationListener{}
	DefaultContext = &DefaultApplicationContext{}
)

type ApplicationContext interface {
	Publish(ent event.ApplicationEvent)
}

type DefaultApplicationContext struct {
	ApplicationContext
	listeners []listener.ApplicationListener
}

func init() {
	register(&listener.AListener{})
	register(&listener.BListener{})
}
func (context *DefaultApplicationContext) GetListener(tp int) listener.ApplicationListener {
	for _, listener := range context.listeners {
		if listener.GetType() == tp {
			return listener
		}
	}
	return nil
}

func NewDefaultApplicationContext() *DefaultApplicationContext {
	context := &DefaultApplicationContext{}
	context.listeners = []listener.ApplicationListener{
		&listener.AListener{},
		&listener.BListener{},
	}
	return context
}

func (context *DefaultApplicationContext) Publish(ent event.ApplicationEvent) {
	/*listener := context.GetListener(ent.GetType())
	if listener != nil {
		go listener.OnApplicationEvent(ent)
	}*/
	if ltn, ok := listenerMap[ent.GetType()]; ok {
		go ltn.OnApplicationEvent(ent)
	}
}

func register(lnt listener.ApplicationListener) {
	listenerMap[lnt.GetType()] = lnt
}

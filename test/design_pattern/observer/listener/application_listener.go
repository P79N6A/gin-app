/**
 * @description observer
 * @author zhangbingbing@baidu.com
 * @date 2019/3/13
 */
package listener

import "gin-app/test/design_pattern/observer/event"

type ApplicationListener interface {
	OnApplicationEvent(ent event.ApplicationEvent)
	GetType() int
}

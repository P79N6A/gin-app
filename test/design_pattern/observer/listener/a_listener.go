/**
 * @description listener
 * @author zhangbingbing@baidu.com
 * @date 2019/3/13
 */
package listener

import (
	"log"
	"gin-app/test/design_pattern/observer/event"
)

type AListener struct {
	ApplicationListener
}

func (aListener *AListener) OnApplicationEvent(ent event.ApplicationEvent) {
	log.Println("run a listener")
}
func (aListener *AListener) GetType() int {
	return 1
}

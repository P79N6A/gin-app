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

type BListener struct {
	ApplicationListener
}

func (bListener *BListener) OnApplicationEvent(ent event.ApplicationEvent) {
	log.Println("run b listener")
}
func (bListener *BListener) GetType() int {
	return 2
}

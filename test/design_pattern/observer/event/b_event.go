/**
 * @description event
 * @author zhangbingbing@baidu.com
 * @date 2019/3/13
 */
package event

type BEvent struct {
	ApplicationEvent
}

func (bEvent *BEvent) GetSource() interface{} {
	return nil
}

func (bEvent *BEvent) GetType() int {
	return 2
}

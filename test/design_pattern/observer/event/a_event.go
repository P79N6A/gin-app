/**
 * @description event
 * @author zhangbingbing@baidu.com
 * @date 2019/3/13
 */
package event

type AEvent struct {
	ApplicationEvent
	data map[string]interface{}
}

func (aEvent *AEvent) GetSource() interface{} {
	return nil
}

func (aEvent *AEvent) GetType() int {
	return 1
}

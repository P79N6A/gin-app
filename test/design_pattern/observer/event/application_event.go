/**
 * @description observer
 * @author zhangbingbing@baidu.com
 * @date 2019/3/13
 */
package event

type ApplicationEvent interface {
	GetSource() interface{}
	GetType() int
}

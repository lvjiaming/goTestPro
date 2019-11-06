/**
 利用空接口实现字典
 */
package MyDictionary

type Dictionary struct {
	data map[interface{}]interface{} // 创建一个map,key 和value都为空接口类型(即任意类型)
}

func NewDictionary () *Dictionary {
	data := &Dictionary{}
	data.data = make(map[interface{}]interface{})
	return data
}

func (d *Dictionary) Get (key interface{}) interface{} {
	data := d.data[key]
	return data
}

func (d *Dictionary) Set (key, value interface{})  {
	d.data[key] = value
}

func (d *Dictionary) Visit (cb func(i, j interface{}) bool)  {
	if cb == nil {
		return
	}
	for key, value := range d.data{
		if !cb(key, value) {
			return
		}
	}
}

func (d *Dictionary) Clear ()  {
	d.data = make(map[interface{}]interface{})
}






/**\
客户管理
 */
package Server

import "myPro/CustomerSystem/Model"

type CustomerManager struct {
	customerList []Model.Customer
	customerNum int
}

/**
 获取所有的客户
 */
func (c *CustomerManager) GetCustomerList () []Model.Customer {
	return c.customerList
}

/**
 添加
 */
func (c *CustomerManager) Add (customer Model.Customer) bool {
	c.customerNum++
	customer.Id = c.customerNum
	c.customerList = append(c.customerList, customer)
	return true
}

/**
 通过id删除
 */
func (c *CustomerManager) Del (id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	} else {
		list1 := c.customerList[:index]
		list2 := c.customerList[index+1:]
		c.customerList = append(list1, list2...)
		return true
	}
}

/**
 通过id查找
 */
func (c *CustomerManager) FindById (id int) int {
	index := -1
	for key, val := range c.customerList{
		if val.Id == id {
			index = key
		}
	}
	return index
}

/**
 返回客户管理实例
 */
func NewCustomerManager() *CustomerManager {
	return &CustomerManager{
		customerNum: 0,
	}
}


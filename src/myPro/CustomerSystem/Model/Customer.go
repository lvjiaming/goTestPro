/**
客户
 */
package Model

import "fmt"

type Customer struct {
	Id int
	Name string
	Sex string
	Age int
	Phone string
	Email string
}

/**
 返回一个客户（包含id）
 */
func NewCustomer(name string, sex string, age int, phone string, email string) *Customer {
	return &Customer{
		Name: name,
		Sex: sex,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

/**
 获取客户信息
 */
func (c *Customer)GetCustomerInfo() string {
	info := fmt.Sprintf("%v\t %v\t %v\t %v\t %v\t %v\t ",
		c.Id, c.Name, c.Sex, c.Age, c.Phone, c.Email)
	return info
}

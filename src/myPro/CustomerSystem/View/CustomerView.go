/**
 操作
 */
package View

import (
	"fmt"
	"myPro/CustomerSystem/Model"
	"myPro/CustomerSystem/Server"
)

type customerView struct {
	key string
	loop bool
	manager *Server.CustomerManager
}

func (c *customerView) CustomerSystem ()  {
	c.manager = Server.NewCustomerManager()
	c.loop = true
	for  {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&c.key)

		switch c.key {
		case "1":
			c.add()
		case "2":
			fmt.Print("修改客户\n")
		case "3":
			c.del()
		case "4":
			c.show()
		case "5":
			c.exit()
		}
		if !c.loop {
			break
		}
	}
	fmt.Println("已退出了客户关系管理系统...")
}

func (c *customerView) add ()  {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := Model.NewCustomer(name, gender, age, phone, email)
	if c.manager.Add(*customer) {
		fmt.Println("---------------------添加成功---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

func (c *customerView) del ()  {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Print("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		c.exit()
	}
	if c.manager.Del(id) {
		fmt.Println("---------------------删除成功---------------------")
	} else {
		fmt.Println("---------------------删除失败---------------------")
	}
}

func (c *customerView) show ()  {
	list := c.manager.GetCustomerList()
	if len(list) == 0 {
		fmt.Println("---------------------没有客户信息---------------------")
	} else {
		fmt.Println("---------------------------客户列表---------------------------")
		fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
		for i := 0; i < len(list); i++ {
			//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
			fmt.Println(list[i].GetCustomerInfo())
		}
		fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
	}
}

func (c *customerView) exit ()  {
	fmt.Print("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&c.key)
		if c.key == "Y" || c.key == "y" || c.key == "N" || c.key == "n" {
			break
		}
		fmt.Print("你的输入有误，确认是否退出(Y/N)：")
	}
	if c.key == "Y" || c.key == "y" {
		c.loop = false
	}
}

func GetView() customerView {
	return customerView{}
}
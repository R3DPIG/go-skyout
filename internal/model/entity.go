package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// BaseModel 公共字段，对应苍穹外卖各表的 id / create_time / update_time
type BaseModel struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

// Employee 员工表
type Employee struct {
	BaseModel
	Name     string `gorm:"column:name;type:varchar(32)" json:"name"`
	Username string `gorm:"column:username;type:varchar(32);uniqueIndex" json:"username"`
	Password string `gorm:"column:password;type:varchar(64)" json:"-"`
	Phone    string `gorm:"column:phone;type:varchar(11)" json:"phone"`
	Sex      string `gorm:"column:sex;type:varchar(2)" json:"sex"`
	IDNumber string `gorm:"column:id_number;type:varchar(18)" json:"idNumber"`
	Status   int    `gorm:"column:status;default:1" json:"status"` // 1 启用 0 禁用
}

func (Employee) TableName() string { return "employee" }

// Category 分类表（菜品分类 / 套餐分类）
type Category struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(32);uniqueIndex" json:"name"`
	Sort int    `gorm:"column:sort" json:"sort"`
}

func (Category) TableName() string { return "category" }

// Dish 菜品表
type Dish struct {
	BaseModel
	Name        string          `gorm:"column:name;type:varchar(64)" json:"name"`
	CategoryID  uint64          `gorm:"column:category_id" json:"categoryId"`
	Price       decimal.Decimal `gorm:"column:price;type:decimal(10,2)" json:"price"`
	Image       string          `gorm:"column:image;type:varchar(200)" json:"image"`
	Description string          `gorm:"column:description;type:varchar(400)" json:"description"`
	Status      int             `gorm:"column:status;default:1" json:"status"` // 1 起售 0 停售
}

func (Dish) TableName() string { return "dish" }

// DishFlavor 菜品口味表
type DishFlavor struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	DishID uint64 `gorm:"column:dish_id" json:"dishId"`
	Name   string `gorm:"column:name;type:varchar(32)" json:"name"`
	Value  string `gorm:"column:value;type:varchar(255)" json:"value"` // JSON 数组，如 ["不辣","微辣"]
}

func (DishFlavor) TableName() string { return "dish_flavor" }

// Setmeal 套餐表
type Setmeal struct {
	BaseModel
	Name       string          `gorm:"column:name;type:varchar(64)" json:"name"`
	CategoryID uint64          `gorm:"column:category_id" json:"categoryId"`
	Price      decimal.Decimal `gorm:"column:price;type:decimal(10,2)" json:"price"`
	Image      string          `gorm:"column:image;type:varchar(200)" json:"image"`
	Status     int             `gorm:"column:status;default:1" json:"status"` // 1 起售 0 停售
}

func (Setmeal) TableName() string { return "setmeal" }

// SetmealDish 套餐-菜品关联表
type SetmealDish struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	SetmealID uint64          `gorm:"column:setmeal_id" json:"setmealId"`
	DishID    uint64          `gorm:"column:dish_id" json:"dishId"`
	Name      string          `gorm:"column:name;type:varchar(32)" json:"name"`
	Price     decimal.Decimal `gorm:"column:price;type:decimal(10,2)" json:"price"`
	Copies    int             `gorm:"column:copies" json:"copies"`
}

func (SetmealDish) TableName() string { return "setmeal_dish" }

// User C 端用户表（微信小程序）
type User struct {
	BaseModel
	OpenID   string `gorm:"column:open_id;type:varchar(64);uniqueIndex" json:"openId"`
	NickName string `gorm:"column:nick_name;type:varchar(32)" json:"nickName"`
	Phone    string `gorm:"column:phone;type:varchar(11)" json:"phone"`
	Sex      string `gorm:"column:sex;type:varchar(2)" json:"sex"`
	Avatar   string `gorm:"column:avatar;type:varchar(200)" json:"avatar"`
}

func (User) TableName() string { return "user" }

// AddressBook 地址簿
type AddressBook struct {
	BaseModel
	UserID    uint64 `gorm:"column:user_id" json:"userId"`
	Consignee string `gorm:"column:consignee;type:varchar(50)" json:"consignee"`
	Phone     string `gorm:"column:phone;type:varchar(11)" json:"phone"`
	Detail    string `gorm:"column:detail;type:varchar(200)" json:"detail"`
	IsDefault int    `gorm:"column:is_default;default:0" json:"isDefault"`
}

func (AddressBook) TableName() string { return "address_book" }

// ShoppingCart 购物车
type ShoppingCart struct {
	BaseModel
	UserID    uint64          `gorm:"column:user_id" json:"userId"`
	DishID    *uint64         `gorm:"column:dish_id" json:"dishId"`
	SetmealID *uint64         `gorm:"column:setmeal_id" json:"setmealId"`
	Name      string          `gorm:"column:name;type:varchar(64)" json:"name"`
	Image     string          `gorm:"column:image;type:varchar(200)" json:"image"`
	Price     decimal.Decimal `gorm:"column:price;type:decimal(10,2)" json:"price"`
	Number    int             `gorm:"column:number" json:"number"`
}

func (ShoppingCart) TableName() string { return "shopping_cart" }

// Orders 订单表；Status: 1 待付款 2 待接单 3 已接单 4 派送中 5 已完成 6 已取消
type Orders struct {
	BaseModel
	Number       string          `gorm:"column:number;type:varchar(50);uniqueIndex" json:"number"`
	UserID       uint64          `gorm:"column:user_id" json:"userId"`
	Amount       decimal.Decimal `gorm:"column:amount;type:decimal(10,2)" json:"amount"`
	Status       int             `gorm:"column:status" json:"status"`
	Phone        string          `gorm:"column:phone;type:varchar(11)" json:"phone"`
	Address      string          `gorm:"column:address;type:varchar(255)" json:"address"`
	Consignee    string          `gorm:"column:consignee;type:varchar(32)" json:"consignee"`
	CheckoutTime time.Time       `gorm:"column:checkout_time" json:"checkoutTime"`
}

func (Orders) TableName() string { return "orders" }

// OrderDetail 订单明细
type OrderDetail struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   uint64          `gorm:"column:order_id" json:"orderId"`
	DishID    *uint64         `gorm:"column:dish_id" json:"dishId"`
	SetmealID *uint64         `gorm:"column:setmeal_id" json:"setmealId"`
	Name      string          `gorm:"column:name;type:varchar(64)" json:"name"`
	Price     decimal.Decimal `gorm:"column:price;type:decimal(10,2)" json:"price"`
	Copies    int             `gorm:"column:copies" json:"copies"`
}

func (OrderDetail) TableName() string { return "order_detail" }

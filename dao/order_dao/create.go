package order_dao

import (
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/module/db"
	"github.com/foxiswho/echo-go/util"
)

type CreateOrder struct {
	Order               *models.Order
	OrderExt            *models.OrderExt
	Consignee           *models.OrderConsignee
	OrderGoods          []*models.OrderGoods
	OrderGoodsStructure []*models.OrderGoodsStructure
}

func NewCreateOrder() *CreateOrder {
	return new(CreateOrder)
}

func (s *CreateOrder) Process() (*models.Order, error) {
	engine := db.DB().Engine
	id, err := engine.InsertOne(s.Order)
	if err != nil {
		return nil, util.NewError("订单保存错误")
	}
	s.Order.Id = int(id)
	s.Consignee.Id = s.Order.Id
	_, err = engine.InsertOne(s.Consignee)
	if err != nil {
		return nil, util.NewError("订单收货人保存错误")
	}
	s.OrderExt.Id = s.Order.Id
	_, err = engine.InsertOne(s.OrderExt)
	if err != nil {
		return nil, util.NewError("订单扩展信息保存错误")
	}
	for _, goods := range s.OrderGoods {
		goods.OrderId = s.Order.Id
		_, err = engine.Insert(goods)
		if err != nil {
			return nil, util.NewError("订单商品保存错误")
		}
	}
	for _, goods := range s.OrderGoodsStructure {
		goods.OrderId = s.Order.Id
		_, err = engine.Insert(goods)
		if err != nil {
			return nil, util.NewError("订单商品保存错误")
		}
	}

	return s.Order, nil
}
package order_service

import (
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/util"
	"github.com/foxiswho/echo-go/module/db"
	"fmt"
	"github.com/foxiswho/echo-go/service/base"
)

//创建订单，数据来源于购物车
type CreateOrderFormatFromCart struct {
	Carts []models.Cart
}

func NewCreateOrderFormatFromCart() *CreateOrderFormatFromCart {
	return new(CreateOrderFormatFromCart)
}

func (s *CreateOrderFormatFromCart) SetCart(carts []models.Cart) {
	s.Carts = carts
}

func (s *CreateOrderFormatFromCart) Process() ([]*models.OrderGoodsData, error) {
	if s.Carts == nil {
		return nil, util.NewError("购物车商品数据不能为空")
	}
	count := len(s.Carts)
	if count == 0 {
		return nil, util.NewError("购物车商品数据不能为空")
	}
	order_goods_data_all := make([]*models.OrderGoodsData, count)
	for key, val := range s.Carts {
		order_goods_data := models.NewOrderGoodsData()
		goods, err := base.NewGoodsService().GetById(val.GoodsId)
		if err != nil {
			return nil, err
		}
		goods_price, err := base.NewGoodsPriceService().GetById(val.GoodsId)
		if err != nil {
			return nil, err
		}
		order_goods_data.Goods = models.NewGoods()
		order_goods_data.Goods = goods
		order_goods_data.GoodsPrice = models.NewGoodsPrice()
		order_goods_data.GoodsPrice = goods_price
		order_goods_data.Num = val.Num
		order_goods_data_all[key] = order_goods_data
	}
	return order_goods_data_all, nil
}
//删除
func (s *CreateOrderFormatFromCart) DeleteCarts() {
	for _, val := range s.Carts {
		affected, err := db.DB().Engine.ID(val.Id).Delete(models.NewCart())
		if err != nil {
			fmt.Println("err=>", err)
		}
		fmt.Println("affected=>", affected)
	}
}

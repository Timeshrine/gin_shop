package dao

import (
	"context"
	"gin_shop/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&OrderDao{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderById(id uint, userId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", id, userId).First(&order).Error
	return
}

func (dao *OrderDao) ListOrderByUserId(uId uint) (orderes []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id=?", uId).Find(&orderes).Error
	return
}

func (dao *OrderDao) UpdateOrderByUserId(aId uint, order *model.Order) error {
	return dao.DB.Model(&OrderDao{}).Where("id=?", aId).Updates(&order).Error
}

func (dao *OrderDao) DeleteOrderByOrderId(aId, uid uint) error {
	return dao.DB.Model(&OrderDao{}).Where("id=? AND user_id=?", aId, uid).Delete(&OrderDao{}).Error
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (order []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).Count(&total).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).
		Find(&order).Error
	return
}
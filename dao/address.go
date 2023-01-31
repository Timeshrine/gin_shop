package dao

import (
	"context"
	"gin_shop/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func (dao *AddressDao) CreateAddress(in *model.Address) error {
	return dao.DB.Model(&AddressDao{}).Create(&in).Error
}

func (dao *AddressDao) GetAddressByAid(aid uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aid).First(&address).Error
	return
}

func (dao *AddressDao) ListAddressByUserId(uId uint) (addresses []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id=?", uId).Find(&addresses).Error
	return
}

func (dao *AddressDao) UpdateAddressByUserId(aId uint, address *model.Address) error {
	return dao.DB.Model(&AddressDao{}).Where("id=?", aId).Updates(&address).Error
}

func (dao *AddressDao) DeleteAddressByAddressId(aId, uid uint) error {
	return dao.DB.Model(&AddressDao{}).Where("id=? AND user_id=?", aId, uid).Delete(&AddressDao{}).Error
}

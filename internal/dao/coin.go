package dao

import (
	"altair-backend/internal/model"
	"altair-backend/internal/request"
)

const tableCoin = "coin"

//新增币种列表
func (d *Dao) CreateCoin(coin *request.CoinCreate) error {
	return d.dao.Table(tableCoin).Create(coin).Error
}

//获取币种列表
func (d *Dao) GetCoinList(name, code string, offset, size uint64) (list []*model.Coin, err error) {
	db := d.dao
	db = db.Table(tableCoin)
	if name != "" {
		db = db.Where(" name = ?", name)
	}
	if code != "" {
		db = db.Where("code = ?", code)
	}
	err = db.Where("is_del = ?", 0).Offset(int(offset)).Limit(int(size)).Find(&list).Error
	return

}

func (d *Dao) CountCoinList(name, code string) (count int64, err error) {
	db := d.dao
	db = db.Table(tableCoin)

	if name != "" {
		db = db.Where(" name = ?", name)
	}
	if name != "" {
		db = db.Where("code = ?", code)
	}
	err = db.Where("is_del = ?", 0).Count(&count).Error
	return
}

func (d *Dao) GetCoin(code string) (coin model.Coin, err error) {
	err = d.dao.Table(coin.TableName()).Where("code = ?", code).Find(&coin).Error
	return
}

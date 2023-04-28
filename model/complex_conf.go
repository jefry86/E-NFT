package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

type NftComplexConf struct {
	ID        int    `json:"id" gorm:"column:id"`
	Title     string `json:"title" gorm:"column:title"`           // 合成标题
	ComplexID uint   `json:"complex_id" gorm:"column:complex_id"` // 合成藏品
	GoodsID   uint   `json:"goods_id" gorm:"column:goods_id"`     // 碎片藏品
	Status    uint8  `json:"status" gorm:"column:status"`         // 是否可用
	DtCreate  uint   `json:"dt_create" gorm:"column:dt_create"`   // 创建时间
	DtUpdate  uint   `json:"dt_update" gorm:"column:dt_update"`   // 更新时间
}

func (m *NftComplexConf) TableName() string {
	return "nft_complex_conf"
}

type NftComplexConfDao struct {
	sourceDB  *gorm.DB
	replicaDB []*gorm.DB
	m         *NftComplexConf
}

func NewNftComplexConfDao(ctx context.Context, dbs ...*gorm.DB) *NftComplexConfDao {
	dao := new(NftComplexConfDao)
	switch len(dbs) {
	case 0:
		panic("database connection required")
	case 1:
		dao.sourceDB = dbs[0]
		dao.replicaDB = []*gorm.DB{dbs[0]}
	default:
		dao.sourceDB = dbs[0]
		dao.replicaDB = dbs[1:]
	}
	return dao
}

func (d *NftComplexConfDao) Create(ctx context.Context, obj *NftComplexConf) error {
	err := d.sourceDB.Model(d.m).Create(&obj).Error
	if err != nil {
		return fmt.Errorf("NftComplexConfDao: %w", err)
	}
	return nil
}

func (d *NftComplexConfDao) Get(ctx context.Context, fields, where string) (*NftComplexConf, error) {
	items, err := d.List(ctx, fields, where, 0, 1)
	if err != nil {
		return nil, fmt.Errorf("NftComplexConfDao: Get where=%s: %w", where, err)
	}
	if len(items) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &items[0], nil
}

func (d *NftComplexConfDao) List(ctx context.Context, fields, where string, offset, limit int) ([]NftComplexConf, error) {
	var results []NftComplexConf
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where).Offset(offset).Limit(limit).Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("NftComplexConfDao: List where=%s: %w", where, err)
	}
	return results, nil
}

func (d *NftComplexConfDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) error {
	err := d.sourceDB.Model(d.m).Where(where, args...).
		Updates(update).Error
	if err != nil {
		return fmt.Errorf("NftComplexConfDao:Update where=%s: %w", where, err)
	}
	return nil
}

func (d *NftComplexConfDao) Delete(ctx context.Context, where string, args ...interface{}) error {
	if len(where) == 0 {
		return gorm.ErrInvalidData
	}
	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
		return fmt.Errorf("NftComplexConfDao: Delete where=%s: %w", where, err)
	}
	return nil
}

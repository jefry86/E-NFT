package service

import (
	"fmt"
	"nft_platform/global"
	"nft_platform/utils"
	"nft_platform/vm/response"
)

type MallPlatform struct {
}

func (m *MallPlatform) List(pageNo, pageSize int) (*response.MallPlatformListRes, error) {
	offset, pageSize := utils.PageOffset(pageNo, pageSize)

	count, err := mallPlatformMode.CountByStatus(1)
	if err != nil {
		global.SLogger.Errorf("count platform err:%s", err.Error())
		return nil, fmt.Errorf("系统错误！")
	}

	if count == 0 {
		return nil, nil
	}

	list, err := mallPlatformMode.ListByStatus(1, offset, pageSize)
	if err != nil {
		global.SLogger.Errorf("list platform err:%s", err.Error())
		return nil, fmt.Errorf("系统错误！")
	}
	res := make([]response.MallPlatformList, 0)
	for _, platform := range *list {
		res = append(res, response.MallPlatformList{
			Name:      platform.Name,
			Logo:      platform.Logo,
			Site:      platform.Site,
			HasOnline: platform.HasOnline,
		})
	}

	return &response.MallPlatformListRes{
		List: res,
		Page: response.Page{
			Total: count,
			Size:  pageSize,
		},
	}, nil
}

package mapers

import (
	"go_server/dto"
	"go_server/storage/dao"
)

func DataToDao(src *dao.Data) *dto.Data {
	if src == nil {
		return nil
	}
	return &dto.Data{
		ID:   src.ID,
		Data: src.Data,
	}
}

func DataToDto(src *dao.Data) *dto.Data {
	if src == nil {
		return nil
	}
	return &dto.Data{
		ID:   src.ID,
		Data: src.Data,
	}
}

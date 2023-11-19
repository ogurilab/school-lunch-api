package domain

import (
	"context"
	"time"
)

type Menu struct {
	ID                       string
	OfferedAt                time.Time // YYYY-MM-DD
	RegionID                 int32
	PhotoUrl                 string
	WikimediaCommonsUrl      string
	ElementarySchoolCalories int32
	JuniorHighSchoolCalories int32
}

type MenuRepository interface {
	Create(ctx context.Context, menu *Menu) error
	FindByID(ctx context.Context, id string) (*Menu, error)
}

type MenuUsecase interface {
	Create(ctx context.Context, menu *Menu) error
	FindByID(ctx context.Context, id string) (*Menu, error)
}

func NewMenu(
	id string,
	offeredAt time.Time,
	regionID int32,
	photoUrl string,
	wikimediaCommonsUrl string,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
) *Menu {
	return &Menu{
		ID:                       id,
		OfferedAt:                offeredAt,
		RegionID:                 regionID,
		PhotoUrl:                 photoUrl,
		WikimediaCommonsUrl:      wikimediaCommonsUrl,
		ElementarySchoolCalories: elementarySchoolCalories,
		JuniorHighSchoolCalories: juniorHighSchoolCalories,
	}
}

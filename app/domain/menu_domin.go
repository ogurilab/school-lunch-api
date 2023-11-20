package domain

import (
	"context"
	"time"

	"github.com/ogurilab/school-lunch-api/util"
)

type Menu struct {
	ID                       string
	OfferedAt                time.Time // YYYY-MM-DD
	RegionID                 int32
	PhotoUrl                 string
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

func newMenu(
	id string,
	offeredAt time.Time,
	regionID int32,
	photoUrl string,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
) (*Menu, error) {

	if _, err := util.ParseUlid(id); err != nil {
		return nil, err
	}

	return &Menu{
		ID:                       id,
		OfferedAt:                offeredAt,
		RegionID:                 regionID,
		PhotoUrl:                 photoUrl,
		ElementarySchoolCalories: elementarySchoolCalories,
		JuniorHighSchoolCalories: juniorHighSchoolCalories,
	}, nil
}

func ReNewMenu(
	id string,
	offeredAt time.Time,
	regionID int32,
	photoUrl string,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
) (*Menu, error) {
	return newMenu(
		id,
		offeredAt,
		regionID,
		photoUrl,
		elementarySchoolCalories,
		juniorHighSchoolCalories,
	)
}

func NewMenu(
	offeredAt time.Time,
	regionID int32,
	photoUrl string,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
) (*Menu, error) {
	id := util.NewUlid()
	return newMenu(
		id,
		offeredAt,
		regionID,
		photoUrl,
		elementarySchoolCalories,
		juniorHighSchoolCalories,
	)
}

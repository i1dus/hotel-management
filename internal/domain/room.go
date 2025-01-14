package domain

type Room struct {
	ID          int
	Number      string
	Type        RoomCategory
	Price       int
	Cleaned     bool
	Description string
}

// RoomCategory категория номера
type RoomCategory int

const (
	RoomCategoryUnknown  RoomCategory = 0
	RoomCategoryStandard RoomCategory = 1
	RoomCategoryComfort  RoomCategory = 2
	RoomCategoryLuxe     RoomCategory = 3
)

// RoomCategoryName название категории номера
type RoomCategoryName string

const (
	RoomCategoryNameUnknown  RoomCategoryName = "неизвестный"
	RoomCategoryNameStandard RoomCategoryName = "стандарт"
	RoomCategoryNameComfort  RoomCategoryName = "комфорт"
	RoomTypeNameLuxe         RoomCategoryName = "люкс"
)

func (r RoomCategory) GetRoomTypeName() RoomCategoryName {
	switch r {
	case RoomCategoryStandard:
		return RoomCategoryNameStandard
	case RoomCategoryComfort:
		return RoomCategoryNameComfort
	case RoomCategoryLuxe:
		return RoomTypeNameLuxe
	}
	return RoomCategoryNameUnknown
}

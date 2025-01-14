package list_rooms_usecase

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strings"
)

type RoomRepository interface {
	ListRooms(ctx context.Context) ([]domain.Room, error)
}

type ListRoomsUseCase struct {
	roomRepo RoomRepository
}

func NewListRoomsUseCase(roomRepo RoomRepository) *ListRoomsUseCase {
	return &ListRoomsUseCase{roomRepo: roomRepo}
}

func (uc *ListRoomsUseCase) ListRooms(c tele.Context) error {
	rooms, err := uc.roomRepo.ListRooms(context.Background())
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	message := strings.Builder{}
	message.WriteString("🛏️ Номера:")

	if len(rooms) == 0 {
		message.WriteString("\nНомера не найдены")
		return c.Send(message.String())
	}

	for _, room := range rooms {
		message.WriteString(fmt.Sprintf("\n\nНомер '%s':\n\t\tКатегория: '%s'\n\t\tЦена за сутки: %d₽\n\t\tУборка: %s",
			room.Number,
			room.Type.GetRoomTypeName(),
			room.Price,
			lo.If(!room.Cleaned, "Нужна").Else("Не нужна"),
		))

		if room.Description != "" {
			message.WriteString(fmt.Sprintf("\n\t\tКомментарий: '%s'", room.Description))
		}
	}
	return c.Send(message.String())
}

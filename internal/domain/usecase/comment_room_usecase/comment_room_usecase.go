package comment_room_usecase

import (
	"context"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strings"
)

type RoomRepository interface {
	ChangeRoomDescription(ctx context.Context, number string, description string) error
}

type CommentRoomUseCase struct {
	roomRepo RoomRepository
}

func NewCommentRoomUseCase(roomRepo RoomRepository) *CommentRoomUseCase {
	return &CommentRoomUseCase{roomRepo: roomRepo}
}

func (uc *CommentRoomUseCase) CommentRoom(c tele.Context) error {
	args := c.Args()
	if len(args) < 2 {
		return c.Send("Должно быть минимум 2 аргумента: Номер, Комментарий")
	}

	// Номер
	number := args[0]

	// Комментарий
	comment := strings.Join(args[1:], " ")

	err := uc.roomRepo.ChangeRoomDescription(context.Background(), number, comment)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send(domain.PrefixSuccess + fmt.Sprintf("Комментарий номера '%s' успешно обновлен на '%s'!", number, comment))
}

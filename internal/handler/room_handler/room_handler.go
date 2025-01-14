package room_handler

import (
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
)

type AddRoomUseCase interface {
	AddRoom(c tele.Context) error
}

type ListRoomsUseCase interface {
	ListRooms(c tele.Context) error
}

type ChangeRoomPriceUseCase interface {
	ChangeRoomPrice(c tele.Context) error
}

type CleanRoomUseCase interface {
	CleanRoom(c tele.Context) error
}

type RoomCleanedUseCase interface {
	RoomCleaned(c tele.Context) error
}

type CommentRoomUseCase interface {
	CommentRoom(c tele.Context) error
}

// RoomHandler has room managing commands
type RoomHandler struct {
	bot                    *tele.Bot
	addRoomUseCase         AddRoomUseCase
	listRoomsUseCase       ListRoomsUseCase
	changeRoomPriceUseCase ChangeRoomPriceUseCase
	cleanRoomUseCase       CleanRoomUseCase
	roomCleanedUseCase     RoomCleanedUseCase
	commentRoomUseCase     CommentRoomUseCase
}

func NewRoomHandler(bot *tele.Bot, addRoomUseCase AddRoomUseCase, listRoomsUseCase ListRoomsUseCase, changeRoomPriceUseCase ChangeRoomPriceUseCase, cleanRoomUseCase CleanRoomUseCase, roomCleanedUseCase RoomCleanedUseCase, commentRoomUseCase CommentRoomUseCase) *RoomHandler {
	return &RoomHandler{bot: bot, addRoomUseCase: addRoomUseCase, listRoomsUseCase: listRoomsUseCase, changeRoomPriceUseCase: changeRoomPriceUseCase, cleanRoomUseCase: cleanRoomUseCase, roomCleanedUseCase: roomCleanedUseCase, commentRoomUseCase: commentRoomUseCase}
}

func (h *RoomHandler) RegisterHandlers() {
	h.bot.Handle(domain.CommandAddRoom, h.addRoomUseCase.AddRoom)
	h.bot.Handle(domain.CommandListRooms, h.listRoomsUseCase.ListRooms)
	h.bot.Handle(domain.CommandChangeRoomPrice, h.changeRoomPriceUseCase.ChangeRoomPrice)
	h.bot.Handle(domain.CommandCleanRoom, h.cleanRoomUseCase.CleanRoom)
	h.bot.Handle(domain.CommandRoomCleaned, h.roomCleanedUseCase.RoomCleaned)
	h.bot.Handle(domain.CommandCommentRoom, h.commentRoomUseCase.CommentRoom)
}

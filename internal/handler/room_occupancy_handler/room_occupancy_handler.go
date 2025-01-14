package room_occupancy_handler

import (
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
)

type AddRoomOccupancyUseCase interface {
	AddRoomOccupancy(c tele.Context) error
}

type ListRoomOccupancyUseCase interface {
	ListRoomOccupancy(c tele.Context) error
}

type EndRoomOccupancyUseCase interface {
	EndRoomOccupancy(c tele.Context) error
}

// RoomOccupancyHandler has room occupancy managing commands
type RoomOccupancyHandler struct {
	bot                      *tele.Bot
	addRoomOccupancyUseCase  AddRoomOccupancyUseCase
	listRoomOccupancyUseCase ListRoomOccupancyUseCase
	endRoomOccupancyUseCase  EndRoomOccupancyUseCase
}

func NewRoomOccupancyHandler(bot *tele.Bot, addRoomOccupancyUseCase AddRoomOccupancyUseCase, listRoomOccupancyUseCase ListRoomOccupancyUseCase, endRoomOccupancyUseCase EndRoomOccupancyUseCase) *RoomOccupancyHandler {
	return &RoomOccupancyHandler{bot: bot, addRoomOccupancyUseCase: addRoomOccupancyUseCase, listRoomOccupancyUseCase: listRoomOccupancyUseCase, endRoomOccupancyUseCase: endRoomOccupancyUseCase}
}

func (h *RoomOccupancyHandler) RegisterHandlers() {
	h.bot.Handle(domain.CommandAddRoomOccupancy, h.addRoomOccupancyUseCase.AddRoomOccupancy)
	h.bot.Handle(domain.CommandListRoomOccupancy, h.listRoomOccupancyUseCase.ListRoomOccupancy)
	h.bot.Handle(domain.CommandEndRoomOccupancy, h.endRoomOccupancyUseCase.EndRoomOccupancy)
}

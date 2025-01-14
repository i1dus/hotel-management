package handler

import (
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase/about_usecase"
	"hotel-management/internal/domain/usecase/add_client_usecase"
	"hotel-management/internal/domain/usecase/add_employee_usecase"
	"hotel-management/internal/domain/usecase/add_room_occupancy_usecase"
	"hotel-management/internal/domain/usecase/add_room_usecase"
	"hotel-management/internal/domain/usecase/change_room_price_usecase"
	"hotel-management/internal/domain/usecase/clean_room_usecase"
	"hotel-management/internal/domain/usecase/comment_room_usecase"
	"hotel-management/internal/domain/usecase/end_room_occupancy_usecase"
	"hotel-management/internal/domain/usecase/help_usecase"
	"hotel-management/internal/domain/usecase/list_employees_usecase"
	"hotel-management/internal/domain/usecase/list_room_occupancies_usecase"
	"hotel-management/internal/domain/usecase/list_rooms_usecase"
	"hotel-management/internal/domain/usecase/remove_employee_usecase"
	"hotel-management/internal/domain/usecase/room_cleaned_usecase"
	"hotel-management/internal/domain/usecase/send_all_message_usecase"
	"hotel-management/internal/domain/usecase/start_usecase"
	"hotel-management/internal/domain/usecase/statatistics_usecase"
	"hotel-management/internal/handler/client_handler"
	"hotel-management/internal/handler/employee_handler"
	"hotel-management/internal/handler/room_handler"
	"hotel-management/internal/handler/room_occupancy_handler"
	"hotel-management/internal/handler/start_handler"
	"hotel-management/internal/repository"
)

type StatisticsUseCase interface {
	Statistics(c tele.Context) error
}

type HelpUseCase interface {
	Help(c tele.Context) error
}

type HandlerController struct {
	bot      *tele.Bot
	handlers []Handler

	statisticsUseCase StatisticsUseCase
	helpUseCase       HelpUseCase
}

type Handler interface {
	RegisterHandlers()
}

func NewHandlerController(bot *tele.Bot, conn *pgx.Conn) *HandlerController {
	clientRepository := repository.NewClientRepository(conn)
	employeeRepository := repository.NewEmployeeRepository(conn)
	roomRepository := repository.NewRoomRepository(conn)
	roomOccupancyRepository := repository.NewRoomOccupancyRepository(conn)
	menu := start_handler.NewMenuWrapper()

	employeeHandler := employee_handler.NewEmployeeHandler(bot,
		add_employee_usecase.NewAddEmployeeUseCase(employeeRepository),
		remove_employee_usecase.NewRemoveEmployeeUseCase(employeeRepository),
		list_employee_usecase.NewListEmployeesUseCase(employeeRepository),
		send_all_message_usecase.NewSendAllMessageUseCase(bot, employeeRepository))

	clientHandler := client_handler.NewClientHandler(bot,
		add_client_usecase.NewAddClientUseCase(clientRepository))

	roomHandler := room_handler.NewRoomHandler(bot,
		add_room_usecase.NewAddRoomUseCase(roomRepository),
		list_rooms_usecase.NewListRoomsUseCase(roomRepository),
		change_room_price_usecase.NewChangeRoomPriceUseCase(roomRepository),
		clean_room_usecase.NewCleanRoomUseCase(roomRepository),
		room_cleaned_usecase.NewRoomCleanedUseCase(roomRepository),
		comment_room_usecase.NewCommentRoomUseCase(roomRepository),
	)

	roomOccupancyHandler := room_occupancy_handler.NewRoomOccupancyHandler(bot,
		add_room_occupancy_usecase.NewAddRoomOccupancyUseCase(roomOccupancyRepository, roomRepository, clientRepository),
		list_room_occupancies_usecase.NewListRoomOccupancyUseCase(roomOccupancyRepository),
		end_room_occupancy_usecase.NewEndRoomOccupancyUseCase(roomOccupancyRepository))

	startHandler := start_handler.NewStartHandler(bot, menu,
		help_usecase.NewHelpUseCase(),
		about_usecase.NewAboutUseCase(),
		start_usecase.NewStartUseCase(employeeRepository, menu.Menu))

	return &HandlerController{
		bot: bot,
		handlers: []Handler{
			employeeHandler,
			clientHandler,
			roomHandler,
			roomOccupancyHandler,
			startHandler,
		},
		helpUseCase:       help_usecase.NewHelpUseCase(),
		statisticsUseCase: statatistics_usecase.NewStatisticsUseCase(roomOccupancyRepository),
	}
}

func (ctrl *HandlerController) RegisterHandlers() {
	ctrl.bot.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(domain.UnknownMessage)
	})
	ctrl.bot.Handle(domain.CommandHelp, ctrl.helpUseCase.Help)
	ctrl.bot.Handle(domain.CommandStatistics, ctrl.statisticsUseCase.Statistics)

	// Зарегистрировать все хэндлеры
	for _, handler := range ctrl.handlers {
		handler.RegisterHandlers()
	}
}

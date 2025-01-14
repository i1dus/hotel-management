package domain

// Employee managing commands
const (
	CommandAddEmployee    = "/add_employee"
	CommandRemoveEmployee = "/remove_employee"
	CommandListEmployees  = "/employees"
	CommandSendAllMessage = "/send_all_message"
)

// Client managing commands
const (
	CommandAddClient = "/add_client"
)

// Room managing commands
const (
	CommandAddRoom         = "/add_room"
	CommandListRooms       = "/rooms"
	CommandChangeRoomPrice = "/change_room_price"
	CommandCleanRoom       = "/clean_room"
	CommandRoomCleaned     = "/room_cleaned"
	CommandCommentRoom     = "/comment_room"
)

// Room occupancy managing commands
const (
	CommandAddRoomOccupancy  = "/add_occupancy"
	CommandListRoomOccupancy = "/occupancies"
	CommandEndRoomOccupancy  = "/end_occupancy"
)

const (
	CommandStart      = "/start"
	CommandHelp       = "/help"
	CommandStatistics = "/stats"
)

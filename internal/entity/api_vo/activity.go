package api_vo

type ActivitItem struct {
}

type AddActivityRequest struct {
	ActName   string `json:"act_name" binding:"required,min=1,max=30"`
	ActType   int    `json:"act_type" binding:"required_if=Version 0"`
	ActStatus int    `json:"act_status" binding:"required_unless=Version 0"`
	StartTime int64  `json:"start_time" binding:"required"`
	EndTime   int64  `json:"end_time" binding:"required"`
	Version   int    `json:"version" binding:"omitempty,gt=0"`
}

type AddActivityReply struct {
	ID int64 `json:"id"`
}

type EditActivityRequest struct {
	ID        int64  `json:"id" binding:"required"`
	ActName   string `json:"act_name" binding:"required,min=1,max=30"`
	ActType   int    `json:"act_type" binding:"required_if=Version 0"`
	ActStatus int    `json:"act_status" binding:"required_unless=Version 0"`
	StartTime int64  `json:"start_time" binding:"required"`
	EndTime   int64  `json:"end_time" binding:"required"`
	Version   int    `json:"version" binding:"omitempty,gt=0"`
}
type EditActivityReply struct {
	ID int64 `json:"id"`
}

type RemoveActivityRequest struct {
}
type RemoveActivityReply struct {
}

type DetailActivityRequest struct {
}
type DetailActivityReply struct {
}

type ListActivityRequest struct {
}
type ListActivityReply struct {
}

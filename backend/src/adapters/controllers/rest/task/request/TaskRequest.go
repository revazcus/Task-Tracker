package request

import (
	"bytes"
	"encoding/json"
	taskDto "task-tracker/boundary/dto/task"
)

type CreateTaskRequest struct {
	Data struct {
		Id         string `json:"id"`
		Attributes struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			Status      string   `json:"status"`
			Priority    string   `json:"priority"`
			Tags        []string `json:"tags"`
			CreatorId   string   `json:"creatorId"`
			PerformerId string   `json:"performerId"`
			Deadline    string   `json:"deadline"`
			Assessment  int      `json:"assessment"`
			TimeCosts   int      `json:"timeCosts"`
			Comments    []string `json:"comments"`
		} `json:"attributes"`
	} `json:"data"`
}

func (r *CreateTaskRequest) FillFromBytes(jsonBytes []byte) error {
	return json.NewDecoder(bytes.NewReader(jsonBytes)).Decode(r)
}

func (r *CreateTaskRequest) CreateTaskDto() *taskDto.TaskDto {
	return &taskDto.TaskDto{
		Id:          r.Data.Id,
		Title:       r.Data.Attributes.Title,
		Description: r.Data.Attributes.Description,
		Status:      r.Data.Attributes.Status,
		Priority:    r.Data.Attributes.Priority,
		Tags:        r.Data.Attributes.Tags,
		CreatorId:   r.Data.Attributes.CreatorId,
		PerformerId: r.Data.Attributes.PerformerId,
		DeadLine:    r.Data.Attributes.Deadline,
		Assessment:  r.Data.Attributes.Assessment,
		TimeCosts:   r.Data.Attributes.TimeCosts,
		Comments:    r.Data.Attributes.Comments,
	}
}

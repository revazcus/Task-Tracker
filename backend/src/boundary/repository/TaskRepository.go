package repositoryInterface

import (
	"context"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	taskEntity "task-tracker/domain/entity/task"
	taskComment "task-tracker/domain/entity/task/comment"
	taskTimeCosts "task-tracker/domain/entity/task/cost"
	taskStatus "task-tracker/domain/entity/task/spec/status"
)

type TaskRepository interface {
	Init(ctx context.Context) error

	Create(ctx context.Context, task *taskEntity.Task) error

	GetAll(ctx context.Context) ([]*taskEntity.Task, error)
	GetById(ctx context.Context, taskId *idPrimitive.EntityId) (*taskEntity.Task, error)

	Update(ctx context.Context, task *taskEntity.Task) (*taskEntity.Task, error)
	UpdatePerformer(ctx context.Context, taskId *idPrimitive.EntityId, performerId *idPrimitive.EntityId) (*taskEntity.Task, error)
	UpdatePerformerAndStatus(ctx context.Context, taskId *idPrimitive.EntityId, performerId *idPrimitive.EntityId, status taskStatus.Status) (*taskEntity.Task, error)
	UpdateTimeCosts(ctx context.Context, taskId *idPrimitive.EntityId, timeCost *taskTimeCosts.TimeCost) (*taskEntity.Task, error)
	UpdateComments(ctx context.Context, taskId *idPrimitive.EntityId, comment *taskComment.Comment) (*taskEntity.Task, error)

	DeleteById(ctx context.Context, taskId *idPrimitive.EntityId) error
}

package taskEntity

import (
	userObject "task-tracker/common/domainObject/shortUser"
	descriptionPrimitive "task-tracker/common/domainPrimitive/description"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	titlePrimitive "task-tracker/common/domainPrimitive/title"
	assessmentPrimitive "task-tracker/domain/entity/task/assessment"
	taskComment "task-tracker/domain/entity/task/comment"
	taskTimeCosts "task-tracker/domain/entity/task/cost"
	taskPriority "task-tracker/domain/entity/task/spec/priority"
	taskStatus "task-tracker/domain/entity/task/spec/status"
	taskTag "task-tracker/domain/entity/task/spec/tag"
	commonTime "task-tracker/infrastructure/tools/time"
)

type Task struct {
	id          *idPrimitive.EntityId
	title       *titlePrimitive.Title
	description *descriptionPrimitive.Description
	status      taskStatus.Status
	priority    taskPriority.Priority
	tags        []*taskTag.Tag
	creator     *userObject.ShortUser
	performer   *userObject.ShortUser
	createAt    *commonTime.Time
	updateAt    *commonTime.Time
	deadline    *commonTime.Time
	//attachments  string // скрины / видео TODO подумать над реализацией
	assessment *assessmentPrimitive.Assessment
	timeCosts  *taskTimeCosts.TimeCosts
	comments   *taskComment.Comments
}

func (t *Task) ID() *idPrimitive.EntityId {
	return t.id
}

func (t *Task) Title() *titlePrimitive.Title {
	return t.title
}

func (t *Task) Description() *descriptionPrimitive.Description {
	return t.description
}

func (t *Task) Status() taskStatus.Status {
	return t.status
}

func (t *Task) Priority() taskPriority.Priority {
	return t.priority
}

func (t *Task) Tags() []*taskTag.Tag {
	return t.tags
}

func (t *Task) Creator() *userObject.ShortUser {
	return t.creator
}

func (t *Task) Performer() *userObject.ShortUser {
	return t.performer
}

func (t *Task) CreateAt() *commonTime.Time {
	return t.createAt
}

func (t *Task) UpdateAt() *commonTime.Time {
	return t.updateAt
}

func (t *Task) Deadline() *commonTime.Time {
	return t.deadline
}

func (t *Task) Assessment() *assessmentPrimitive.Assessment {
	return t.assessment
}

func (t *Task) TimeCosts() *taskTimeCosts.TimeCosts {
	return t.timeCosts
}

func (t *Task) Comments() *taskComment.Comments {
	return t.comments
}

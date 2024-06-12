package usecase

import "github.com/Fumiya-Tahara/dev-navi/internal/entity"

type ProjectRepository interface {
	Create(project *entity.Project) error
	FindById(id uint) (*entity.Project, error)
	FindByUserId(userID uint) ([]*entity.Project, error)
	Update(project *entity.Project) error
	Delete(project *entity.Project, id uint) error
}

type MilestoneRepository interface {
	Create(milestone *entity.Milestone) error
	FindByProjectId(id uint) (*entity.Milestone, error)
	Update(milestone *entity.Milestone) error
	Delete(milestone *entity.Milestone, id uint) error
}

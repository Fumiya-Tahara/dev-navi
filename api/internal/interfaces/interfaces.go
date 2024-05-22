package interfaces

import "github.com/Fumiya-Tahara/dev-navi/internal/entity"

// Add method for CRUD
//Service is the interface for web APIs, while repository is the interface for database operations

type UserService interface {
}

type ProjectService interface {
}

type MilestoneService interface {
}

// Adding methods grouped under 'issues'
type UserRepository interface{}

type ProjectRepository interface {
	Create(project *entity.Project) error
	FindById(id uint) (*entity.Project, error)
	FindByUserId(userId uint) ([]*entity.Project, error)
	Update(project *entity.Project) error
	Delete(project *entity.Project, id uint) error
}

type MilestoneRepository interface {
	Create(milestone *entity.Milestone) error
	FindByProjectId(id uint) (*entity.Milestone, error)
	Update(milestone *entity.Milestone) error
	Delete(milestone *entity.Milestone, id uint) error
}

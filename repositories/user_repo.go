package repositories

import (
	"beew/datasource"
	"beew/models"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	GetByPhone(string) (models.User, error)
	ExistsByID(int) (bool, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &UserRepository{db: datasource.NewDbInstance()}
}

func (u *UserRepository) ExistsByID(id int) (bool, error) {
	var user models.User
	err := u.db.Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (u *UserRepository) GetByPhone(phone string) (result models.User, err error) {
	where := map[string]interface{}{"phone": phone}
	err = u.db.Debug().Where(where).First(&result).Error
	return
}

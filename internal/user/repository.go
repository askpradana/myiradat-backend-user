package user

import (
	"gorm.io/gorm"
	"myiradat-backend-auth/internal/user/dto"
	"myiradat-backend-auth/internal/user/model"
)

type Repository interface {
	GetServicesWithRoles() ([]model.Service, error)
	FindProfileByEmail(profile *model.Profile, email string) error
	FindProfileByID(profile *model.Profile, id int) error
	GetProfileServicesWithRoles(profileID int) ([]dto.ServiceWithRole, error)
	ListProfiles(limit, offset int) ([]model.Profile, int64, error)
	RoleBelongsToService(roleID, serviceID int) bool
	CreateProfileWithRolesTx(profile *model.Profile, relations []model.ProfileServiceRole) error
	UpdateProfileWithRolesTx(profile *model.Profile, relations []model.ProfileServiceRole) error
	CreateProfileServiceRoles(relations []model.ProfileServiceRole) error
	UpdateProfile(profile *model.Profile) error
	//ipro
	FindIproTestByProfileID(test *model.IproTest, profileID int) error
	FindIprosTestByProfileID(test *model.IprosTest, profileID int) error
	FindIprobTestByProfileID(test *model.IprobTest, profileID int) error
	CreateIproTest(test *model.IproTest) error
	CreateIprosTest(test *model.IprosTest) error
	CreateIprobTest(test *model.IprobTest) error
	UpdateIproTest(test *model.IproTest) error
	UpdateIprosTest(test *model.IprosTest) error
	UpdateIprobTest(test *model.IprobTest) error
	//improve care
	CreateConsultWithCommentsTx(consult *model.Consult, comments []model.Comment) error
	FindLatestConsultByProfileID(profileID int) (*model.Consult, error)
	FindCommentsByConsultID(consultID int) ([]model.Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetServicesWithRoles() ([]model.Service, error) {
	var services []model.Service
	err := r.db.
		Preload("Roles").
		Where("is_deleted = false").
		Find(&services).Error

	return services, err
}

func (r *repository) FindProfileByEmail(profile *model.Profile, email string) error {
	return r.db.Where("email = ?", email).First(profile).Error
}

func (r *repository) FindProfileByID(profile *model.Profile, id int) error {
	return r.db.First(profile, id).Error
}

func (r *repository) GetProfileServicesWithRoles(profileID int) ([]dto.ServiceWithRole, error) {
	var results []dto.ServiceWithRole

	err := r.db.Table("profile_service_roles psr").
		Select(`psr.service_id, s.service_name, psr.role_id, r.role_name`).
		Joins("JOIN services s ON s.id = psr.service_id").
		Joins("JOIN roles r ON r.id = psr.role_id AND r.master_service_id = psr.service_id").
		Where("psr.profile_id = ?", profileID).
		Scan(&results).Error

	return results, err
}

func (r *repository) ListProfiles(limit, offset int) ([]model.Profile, int64, error) {
	var profiles []model.Profile
	var total int64

	query := r.db.Model(&model.Profile{}).Where("is_deleted = false")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at desc").Find(&profiles).Error; err != nil {
		return nil, 0, err
	}

	return profiles, total, nil
}

func (r *repository) RoleBelongsToService(roleID, serviceID int) bool {
	var count int64
	r.db.Model(&model.Role{}).
		Where("id = ? AND master_service_id = ?", roleID, serviceID).
		Count(&count)
	return count > 0
}

func (r *repository) CreateProfileWithRolesTx(profile *model.Profile, relations []model.ProfileServiceRole) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Insert profile
		if err := tx.Create(profile).Error; err != nil {
			return err
		}

		// Tambahkan profileID ke setiap relasi
		for i := range relations {
			relations[i].ProfileID = profile.ID
		}

		// Insert relations
		if err := tx.Create(&relations).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *repository) UpdateProfileWithRolesTx(profile *model.Profile, relations []model.ProfileServiceRole) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Update profile
		if err := tx.Save(profile).Error; err != nil {
			return err
		}

		// Hapus relasi lama
		if err := tx.Where("profile_id = ?", profile.ID).Delete(&model.ProfileServiceRole{}).Error; err != nil {
			return err
		}

		// Tambahkan profileID ke setiap relasi baru
		for i := range relations {
			relations[i].ProfileID = profile.ID
		}

		// Insert relasi baru
		if err := tx.Create(&relations).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *repository) UpdateProfile(profile *model.Profile) error {
	return r.db.Save(profile).Error
}

func (r *repository) CreateProfileServiceRoles(relations []model.ProfileServiceRole) error {
	return r.db.Create(&relations).Error
}

func (r *repository) FindIproTestByProfileID(test *model.IproTest, profileID int) error {
	return r.db.Where("profile_id = ?", profileID).First(test).Error
}

func (r *repository) FindIprosTestByProfileID(test *model.IprosTest, profileID int) error {
	return r.db.Where("profile_id = ?", profileID).First(test).Error
}

func (r *repository) FindIprobTestByProfileID(test *model.IprobTest, profileID int) error {
	return r.db.Where("profile_id = ?", profileID).First(test).Error
}

func (r *repository) CreateIproTest(test *model.IproTest) error {
	return r.db.Create(test).Error
}

func (r *repository) CreateIprosTest(test *model.IprosTest) error {
	return r.db.Create(test).Error
}

func (r *repository) CreateIprobTest(test *model.IprobTest) error {
	return r.db.Create(test).Error
}

func (r *repository) UpdateIproTest(test *model.IproTest) error {
	return r.db.Save(test).Error
}

func (r *repository) UpdateIprosTest(test *model.IprosTest) error {
	return r.db.Save(test).Error
}

func (r *repository) UpdateIprobTest(test *model.IprobTest) error {
	return r.db.Save(test).Error
}

func (r *repository) CreateConsultWithCommentsTx(consult *model.Consult, comments []model.Comment) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(consult).Error; err != nil {
			return err
		}

		for i := range comments {
			comments[i].ConsultID = &consult.ID
			if err := tx.Create(&comments[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *repository) FindLatestConsultByProfileID(profileID int) (*model.Consult, error) {
	var consult model.Consult
	err := r.db.
		Where("profile_id = ?", profileID).
		Order("created_at desc").
		First(&consult).Error
	if err != nil {
		return nil, err
	}
	return &consult, nil
}

func (r *repository) FindCommentsByConsultID(consultID int) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Where("consult_id = ?", consultID).Find(&comments).Error
	return comments, err
}

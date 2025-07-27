package user

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"myiradat-backend-auth/internal/user/dto"
	"myiradat-backend-auth/internal/user/model"
	"time"
)

type Service interface {
	GetProfileSummary(email string) (dto.ProfileSummaryResponse, error)
	GetProfileDetail(email string) (*dto.GetProfileDetailResponse, error)
	GetProfileDetailByID(id int) (*dto.GetProfileDetailResponse, error)
	GetServicesWithRoles() ([]dto.ServiceWithRolesDTO, error)
	ListProfiles(req dto.ListProfilesRequest) (dto.PaginatedResponse[dto.ProfileResponse], error)
	CreateProfile(input dto.CreateProfileRequest) error
	UpdateProfileWithRoles(input dto.UpdateProfileWithRolesRequest) error
	UpdateProfile(input dto.UpdateProfileRequest) error
	CreateIproTest(input dto.CreateIproTestRequest) error
	CreateIprosTest(input dto.CreateIprosTestRequest) error
	CreateIprobTest(input dto.CreateIprobTestRequest) error
	CreateConsultWithComments(input dto.CreateConsultRequest) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetProfileSummary(email string) (dto.ProfileSummaryResponse, error) {
	var profile model.Profile
	if err := s.repo.FindProfileByEmail(&profile, email); err != nil {
		return dto.ProfileSummaryResponse{}, errors.New("profile not found")
	}

	// ============================
	// Consult & Comments Section
	// ============================
	var consultDate *time.Time
	var analysisResult string
	var commentDTOs []dto.CommentDTO

	if consult, err := s.repo.FindLatestConsultByProfileID(profile.ID); err == nil && consult != nil {
		consultDate = consult.ConsultDate
		analysisResult = consult.AnalysisResult

		if comments, err := s.repo.FindCommentsByConsultID(consult.ID); err == nil {
			for _, c := range comments {
				commentDTOs = append(commentDTOs, dto.CommentDTO{
					ID:      c.ID,
					Comment: c.Comment,
				})
			}
		}
	}

	// ============================
	// Test Sections
	// ============================
	var iproResult dto.IproTestResult
	var iproDate *time.Time
	if ipro := s.getIproTest(profile.ID); ipro != nil {
		iproDate = ipro.TestTakenDate
		_ = json.Unmarshal(ipro.Result, &iproResult)
	}

	var iprobResult dto.IproTestResult
	var iprobDate *time.Time
	if iprob := s.getIprobTest(profile.ID); iprob != nil {
		iprobDate = iprob.TestTakenDate
		_ = json.Unmarshal(iprob.Result, &iprobResult)
	}

	var iprosResult dto.IprosTestResult
	var iprosDate *time.Time
	if ipros := s.getIprosTest(profile.ID); ipros != nil {
		iprosDate = ipros.TestTakenDate
		_ = json.Unmarshal(ipros.Result, &iprosResult)
	}

	// ============================
	// Final Response Assembly
	// ============================
	return dto.ProfileSummaryResponse{
		Profile: dto.ProfileInfoDTO{
			Email: profile.Email,
			NoHP:  profile.NoHP,
			Name:  profile.Name,
		},
		Consults: dto.ConsultDTO{
			ConsultDate:          consultDate,
			LatestAnalysisResult: analysisResult,
			LatestComments:       commentDTOs,
		},
		Tests: dto.TestsDTO{
			Ipro: dto.TestResultDTO[dto.IproTestResult]{
				TestTakenDate: iproDate,
				Result:        iproResult,
			},
			Iprob: dto.TestResultDTO[dto.IproTestResult]{
				TestTakenDate: iprobDate,
				Result:        iprobResult,
			},
			Ipros: dto.TestResultDTO[dto.IprosTestResult]{
				TestTakenDate: iprosDate,
				Result:        iprosResult,
			},
		},
	}, nil
}

func (s *service) getIproTest(profileID int) *model.IproTest {
	var t model.IproTest
	if err := s.repo.FindIproTestByProfileID(&t, profileID); err != nil {
		return nil
	}
	return &t
}

func (s *service) getIprobTest(profileID int) *model.IprobTest {
	var t model.IprobTest
	if err := s.repo.FindIprobTestByProfileID(&t, profileID); err != nil {
		return nil
	}
	return &t
}

func (s *service) getIprosTest(profileID int) *model.IprosTest {
	var t model.IprosTest
	if err := s.repo.FindIprosTestByProfileID(&t, profileID); err != nil {
		return nil
	}
	return &t
}

func (s *service) GetProfileDetail(email string) (*dto.GetProfileDetailResponse, error) {

	var profile model.Profile
	if err := s.repo.FindProfileByEmail(&profile, email); err != nil {
		return nil, err
	}

	services, err := s.repo.GetProfileServicesWithRoles(profile.ID)
	if err != nil {
		return nil, err
	}

	return &dto.GetProfileDetailResponse{
		ID:       profile.ID,
		Name:     profile.Name,
		Email:    profile.Email,
		NoHP:     profile.NoHP,
		Services: services,
	}, nil
}

func (s *service) GetProfileDetailByID(id int) (*dto.GetProfileDetailResponse, error) {
	var profile model.Profile
	if err := s.repo.FindProfileByID(&profile, id); err != nil {
		return nil, err
	}

	services, err := s.repo.GetProfileServicesWithRoles(profile.ID)
	if err != nil {
		return nil, err
	}

	return &dto.GetProfileDetailResponse{
		ID:       profile.ID,
		Name:     profile.Name,
		Email:    profile.Email,
		NoHP:     profile.NoHP,
		Services: services,
	}, nil
}

func (s *service) GetServicesWithRoles() ([]dto.ServiceWithRolesDTO, error) {
	services, err := s.repo.GetServicesWithRoles()
	if err != nil {
		return nil, err
	}

	var result []dto.ServiceWithRolesDTO
	for _, svc := range services {
		var roles []dto.RoleDTO
		for _, r := range svc.Roles {
			roles = append(roles, dto.RoleDTO{
				RoleID:   r.ID,
				RoleName: r.RoleName,
			})
		}
		result = append(result, dto.ServiceWithRolesDTO{
			ServiceID:   svc.ID,
			ServiceName: svc.ServiceName,
			Roles:       roles,
		})
	}

	return result, nil
}

func (s *service) CreateProfile(input dto.CreateProfileRequest) error {
	// 1. Cek email sudah digunakan?
	var existing model.Profile
	if err := s.repo.FindProfileByEmail(&existing, input.Email); err == nil {
		return errors.New("email already used")
	}

	// 2. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. Validasi setiap roleId cocok dengan serviceId
	for _, srv := range input.Services {
		valid := s.repo.RoleBelongsToService(srv.RoleID, srv.ServiceID)
		if !valid {
			return errors.New("invalid role-service combination")
		}
	}

	// 4. Buat objek profile
	profile := model.Profile{
		Name:     input.Name,
		Email:    input.Email,
		NoHP:     input.NoHP,
		Password: string(hashedPassword),
	}

	// 5. Buat relasi
	var relations []model.ProfileServiceRole
	for _, srv := range input.Services {
		relations = append(relations, model.ProfileServiceRole{
			ServiceID: srv.ServiceID,
			RoleID:    srv.RoleID,
		})
	}

	// 6. Jalankan sebagai transaksi
	return s.repo.CreateProfileWithRolesTx(&profile, relations)
}

func (s *service) ListProfiles(req dto.ListProfilesRequest) (dto.PaginatedResponse[dto.ProfileResponse], error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	profiles, total, err := s.repo.ListProfiles(pageSize, offset, req.Search)
	if err != nil {
		return dto.PaginatedResponse[dto.ProfileResponse]{}, err
	}

	var result []dto.ProfileResponse
	for _, p := range profiles {
		result = append(result, dto.ProfileResponse{
			ID:    p.ID,
			Name:  p.Name,
			Email: p.Email,
			NoHP:  p.NoHP,
		})
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	return dto.PaginatedResponse[dto.ProfileResponse]{
		Data:       result,
		Page:       page,
		PageSize:   pageSize,
		TotalRows:  total,
		TotalPages: totalPages,
	}, nil
}

func (s *service) UpdateProfileWithRoles(input dto.UpdateProfileWithRolesRequest) error {
	// 1. Cari profile berdasarkan ID
	var existing model.Profile
	if err := s.repo.FindProfileByID(&existing, input.ProfileID); err != nil {
		return errors.New("profile not found")
	}

	// 2. Validasi kombinasi role-service
	for _, srv := range input.Services {
		if !s.repo.RoleBelongsToService(srv.RoleID, srv.ServiceID) {
			return errors.New("invalid role-service combination")
		}
	}

	// 3. Update field profile
	existing.Name = input.Name
	existing.Email = input.Email
	existing.NoHP = input.NoHP

	// 4. Hash password jika diisi
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existing.Password = string(hashedPassword)
	}

	// 5. Susun relasi baru
	var relations []model.ProfileServiceRole
	for _, srv := range input.Services {
		relations = append(relations, model.ProfileServiceRole{
			ServiceID: srv.ServiceID,
			RoleID:    srv.RoleID,
		})
	}

	// 6. Jalankan dalam transaction
	return s.repo.UpdateProfileWithRolesTx(&existing, relations)
}

func (s *service) UpdateProfile(input dto.UpdateProfileRequest) error {
	var existing model.Profile
	if err := s.repo.FindProfileByID(&existing, input.ProfileID); err != nil {
		return errors.New("profile not found")
	}

	existing.Name = input.Name
	existing.Email = input.Email
	existing.NoHP = input.NoHP

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existing.Password = string(hashedPassword)
	}

	return s.repo.UpdateProfile(&existing)
}

func (s *service) CreateIproTest(input dto.CreateIproTestRequest) error {
	var profile model.Profile
	if err := s.repo.FindProfileByEmail(&profile, input.Email); err != nil {
		return errors.New("profile not found")
	}

	resultJSON, err := json.Marshal(input.Result)
	if err != nil {
		return err
	}

	var existing model.IproTest
	err = s.repo.FindIproTestByProfileID(&existing, profile.ID)
	if err == nil {
		// ✅ Update
		existing.TestID = input.TestID
		existing.Result = resultJSON
		existing.TestTakenDate = input.TestTakenDate
		existing.ModifiedBy = "system"

		return s.repo.UpdateIproTest(&existing)
	}

	// ❌ Create baru
	test := model.IproTest{
		TestID:        input.TestID,
		Result:        resultJSON,
		TestTakenDate: input.TestTakenDate,
		ProfileID:     &profile.ID,
		CreatedBy:     "system",
	}
	return s.repo.CreateIproTest(&test)
}

func (s *service) CreateIprosTest(input dto.CreateIprosTestRequest) error {
	var profile model.Profile
	if err := s.repo.FindProfileByEmail(&profile, input.Email); err != nil {
		return errors.New("profile not found")
	}

	resultJSON, err := json.Marshal(input.Result)
	if err != nil {
		return err
	}

	var existing model.IprosTest
	err = s.repo.FindIprosTestByProfileID(&existing, profile.ID)
	if err == nil {
		// ✅ Update
		existing.TestID = input.TestID
		existing.Result = resultJSON
		existing.TestTakenDate = input.TestTakenDate
		existing.ModifiedBy = "system"

		return s.repo.UpdateIprosTest(&existing)
	}

	// ❌ Create baru
	test := model.IprosTest{
		TestID:        input.TestID,
		Result:        resultJSON,
		TestTakenDate: input.TestTakenDate,
		ProfileID:     &profile.ID,
		CreatedBy:     "system",
	}
	return s.repo.CreateIprosTest(&test)
}

func (s *service) CreateIprobTest(input dto.CreateIprobTestRequest) error {
	var profile model.Profile
	if err := s.repo.FindProfileByEmail(&profile, input.Email); err != nil {
		return errors.New("profile not found")
	}

	resultJSON, err := json.Marshal(input.Result)
	if err != nil {
		return err
	}

	var existing model.IprobTest
	err = s.repo.FindIprobTestByProfileID(&existing, profile.ID)
	if err == nil {
		// ✅ Jika sudah ada → update
		existing.TestID = input.TestID
		existing.Result = resultJSON
		existing.TestTakenDate = input.TestTakenDate
		existing.ModifiedBy = "system"

		return s.repo.UpdateIprobTest(&existing)
	}

	// ❌ Jika belum ada → insert baru
	test := model.IprobTest{
		TestID:        input.TestID,
		Result:        resultJSON,
		TestTakenDate: input.TestTakenDate,
		ProfileID:     &profile.ID,
		CreatedBy:     "system",
	}
	return s.repo.CreateIprobTest(&test)
}

func (s *service) CreateConsultWithComments(input dto.CreateConsultRequest) error {
	var profile model.Profile
	if err := s.repo.FindProfileByEmail(&profile, input.Email); err != nil {
		return errors.New("profile not found")
	}

	consult := model.Consult{
		Owner:          input.Owner,
		ConsultDate:    input.ConsultDate,
		AnalysisResult: input.AnalysisResult,
		ProfileID:      &profile.ID,
		CreatedBy:      "system",
	}

	var comments []model.Comment
	for _, c := range input.Comments {
		comments = append(comments, model.Comment{
			Comment:   c.Comment,
			CreatedBy: "system",
		})
	}

	return s.repo.CreateConsultWithCommentsTx(&consult, comments)
}

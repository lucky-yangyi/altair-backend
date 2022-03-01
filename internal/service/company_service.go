package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/pkg/utils"

	"context"
	"errors"
)

var GlobalCompanyService *CompanyService

//初始化
func GetCompanyService() *CompanyService {
	GlobalCompanyService = newCompanyService(dao.GetDao())
	return GlobalCompanyService
}

type CompanyService struct {
	baseService
}

func newCompanyService(dao *dao.Dao) *CompanyService {
	return &CompanyService{baseService{dao: dao, ctx: context.Background()}}
}

//获取企业列表
func (s *CompanyService) GetCompanyList(ctx context.Context, req request.CompanyListReq) (resp *response.CompanyListResp, err error) {

	list, err := s.dao.GetCompanyList(req.Params, utils.GetPageOffset(req.PageNo, req.PageSize), req.PageSize)
	if err != nil {
		return
	}
	count, err := s.dao.GetCompanyCount(req.Params)
	if err != nil {
		return
	}
	resp = &response.CompanyListResp{
		List: list,
		Page: response.Page{
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
			TotalCount: uint64(count),
		},
	}
	return resp, nil
}

func GetCompany(id uint64) (company model.Company) {
	dao.DB.Table(company.TableName()).Where("id = ?", id).First(&company)
	return
}

//新增企业
func (s *CompanyService) CreateCompany(ctx context.Context, req request.CreateCompanyReq) (company *model.Company, err error) {
	company = &model.Company{
		Name: req.CompanyName,
	}
	company.Name = req.CompanyName
	err = s.dao.CreateCompany(company)
	return
}

//企业修改
func (s *CompanyService) UpdateConpany(ctx context.Context, req request.UpdateCompanyReq) error {
	company := &model.Company{
		Name: req.CompanyName,
	}
	err1 := s.dao.UpdateCompany(req.CompanyId, company)
	if err1 != nil {
		return err1
	}
	return nil
}

//企业禁用/启用
func (s *CompanyService) EnableCompany(ctx context.Context, req request.EnableCompanyReq) (err error) {
	err = s.dao.UpdateCompanyEnable(req.Status, req.CompanyId)
	if err != nil {
		return err
	}
	return nil
}

//新增企业成员
func (s *CompanyService) AddCompanyMember(ctx context.Context, req request.AddCompanyMemberReq) error {
	//检测邮箱的唯一性
	ok, err := s.dao.CheckEmailExist(req.Email)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("邮箱已存在！")
	}
	pwdMd5 := utils.GenerateMd5(req.MemberPwd)
	companyMember := &model.CompanyMember{
		CompanyID: req.CompanyID,
		Name:      req.MemberName,
		Password:  pwdMd5,
		Email:     req.Email,
		Desc:      req.Desc,
	}
	err1 := s.dao.AddCompanyMember(companyMember)
	if err != nil {
		return err1
	}
	return nil
}

//获取企业成员列表
func (s *CompanyService) GetMemberList(ctx context.Context, req request.MemberListReq) (resp *response.MemberListResp, err error) {

	list, err := s.dao.GetMemberList(req.CompanyId, utils.GetPageOffset(req.PageNo, req.PageSize), req.PageSize)
	if err != nil {
		return
	}
	count, err := s.dao.GetMemberCount(req.CompanyId)
	if err != nil {
		return
	}
	resp = &response.MemberListResp{
		List: list,
		Page: response.Page{
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
			TotalCount: uint64(count),
		},
	}
	return resp, nil
}

//设置管理员
func (s *CompanyService) AdminMember(ctx context.Context, req request.AdminMemberReq) (err error) {
	err = s.dao.UpdateMemberAdmin(req.Status, req.MemberId)
	if err != nil {
		return err
	}
	return nil
}

//新增成员
func (s *CompanyService) CreateMember(ctx context.Context, req request.CreateMemberReq) error {
	//检测邮箱的唯一性
	ok, err := s.dao.CheckEmailExist(req.Email)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("邮箱已存在！")
	}
	pwdMd5 := utils.GenerateMd5(req.MemberPwd)
	companyMember := &model.CompanyMember{
		CompanyID: req.CompanyID,
		Name:      req.MemberName,
		Password:  pwdMd5,
		Email:     req.Email,
		Desc:      req.Desc,
	}
	err1 := s.dao.CreateMember(nil, companyMember)
	if err != nil {
		return err1
	}
	return nil
}

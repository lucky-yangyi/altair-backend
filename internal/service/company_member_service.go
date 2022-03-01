package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/log"
	"altair-backend/pkg/utils"
	"context"
	"errors"
)

var GlobalCompanyMemberService *CompayMemberService

//初始化
func GetCompanyMemberService() *CompayMemberService {
	GlobalCompanyMemberService = newCompanyMemberService(dao.GetDao())
	return GlobalCompanyMemberService
}

type CompayMemberService struct {
	baseService
}

func newCompanyMemberService(dao *dao.Dao) *CompayMemberService {
	return &CompayMemberService{baseService{dao: dao, ctx: context.Background()}}
}

func (s *CompayMemberService) GetCompanyMember(memberId uint64) (member model.CompanyMember) {
	dao.DB.Table(member.TableName()).Where("id = ?", memberId).First(&member)
	return
}

//获取成员列表
func (s *CompayMemberService) GetCompanyMemberList(ctx context.Context, companyId uint64, req request.CompanyMemberListReq) (resp *response.CompanyMemberListResp, err error) {

	list, err := s.dao.GetCompanyMemberList(companyId, utils.GetPageOffset(req.PageNo, req.PageSize), req.PageSize)
	if err != nil {
		return
	}
	count, err := s.dao.GetCompanyMemberCount(companyId)
	if err != nil {
		return
	}
	resp = &response.CompanyMemberListResp{
		List: list,
		Page: response.Page{
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
			TotalCount: uint64(count),
		},
	}
	return resp, nil
}

//新增成员
func (s *CompayMemberService) CreateCompanyMember(companyId uint64, req request.CreateCompanyMemberReq) error {
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
		CompanyID: companyId,
		Name:      req.MemberName,
		Password:  pwdMd5,
		Email:     req.Email,
		Desc:      req.Desc,
		IsAdmin:   false,
		Enabled:   1,
	}
	err1 := s.dao.CreateCompanyMember(nil, companyMember)
	if err != nil {
		return err1
	}
	return nil
}

//增加钱包权限
func (s *CompayMemberService) AddWalletPermission(ctx context.Context, req request.AddWalletPermissionReq) error {
	//检测钱包是否存在
	ok := s.dao.CheckWalletExist(req.Ids)
	if !ok {
		return errors.New("钱包不存在！")
	}
	////检测钱包权限是否存在
	//ok, err := s.dao.CheckWalletAuthExist(req.Ids, req.MemberId)
	//if err != nil {
	//	return err
	//}
	//if ok {
	//	return errcode.WalletAuthExistError
	//}
	var list []*model.WalletAuthorize
	for _, v := range req.Ids {
		list = append(list, &model.WalletAuthorize{
			CompanyMemberID: req.MemberId,
			WalletID:        uint64(v),
		})
	}
	return s.dao.CreateWalletPermission(nil, list)
}

//删除钱包权限
func (s *CompayMemberService) DelWalletPermission(ctx context.Context, memberId uint64) error {
	err := s.dao.DelWalletPermission(memberId)
	if err != nil {
		return err
	}
	return nil
}

//获取钱包权限
func (s *CompayMemberService) GetMemberWalletPermission(ctx context.Context, memberId uint64) (resp *response.MemberWalletAuthResp, err error) {
	list, err := s.dao.GetWalletByMemberId(memberId)
	if err != nil {
		return nil, err
	}
	resp = &response.MemberWalletAuthResp{
		List: list,
	}
	return
}

//成员禁用/启用
func (s *CompayMemberService) EnableMember(ctx context.Context, req request.EnableMemberReq) (err error) {
	err = s.dao.UpdateMemberEnable(req.Status, req.MemberId)
	if err != nil {
		return err
	}
	return nil
}

//成员修改
func (s *CompayMemberService) UpdateConpanyMember(ctx context.Context, req request.UpdateCompanyMemberReq) error {
	//检测邮箱的唯一性
	ok, err := s.dao.CheckEmailExist(req.Email)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("邮箱已存在！")
	}
	member := &model.CompanyMember{
		Name:  req.MemberName,
		Email: req.Email,
		Desc:  req.Desc,
	}
	err1 := s.dao.UpdateCompanyMember(req.MemberId, member)
	if err1 != nil {
		return err1
	}
	return nil
}
func GetMember(email string) (user model.CompanyMember, err error) {
	return dao.GlobalDao.GetMember(email)
}

//成员密码修改
func (s *CompayMemberService) ChangePassword(ctx context.Context, req request.ChangePassword) error {
	data, err := s.dao.CheckPassword(req.MemberId)
	if req.MemberPwd == data.Password {
		log.Fatal("无此用户：", req.MemberId)
		return errors.New("新密码与旧密码不能一样！")
	}
	member := &model.CompanyMember{
		Password: req.MemberPwd,
	}
	err1 := s.dao.ChangePassword(req.MemberId, member)
	if err != nil {
		return err1
	}
	return nil
}

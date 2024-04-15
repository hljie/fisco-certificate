package api

import (
	"certificate-backend/sdk"
	"certificate-backend/serializer"
	"certificate-backend/util"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type CertificateVo struct {
	Id              *big.Int
	StudentName     string
	Grade           *big.Int
	Course          string
	StudentAddress  string
	ApplicationTime *big.Int
	ApprovalTime    *big.Int
	Status          string
	Approver        string
}

// ### 证书相关接口
// 根据id获取证书
type GetCertificateAo struct {
	Id *big.Int `json:"id" validate:"required"`
}

func GetCertificate(c *gin.Context) {
	ao := new(GetCertificateAo)
	if err := c.ShouldBind(ao); err == nil {
		id, studentName, grade, course, studentAddress, applicationTime, approvalTime, status, approver, err := sdk.Contract.GetCertificateDetails(ao.Id)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		vo := CertificateVo{
			Id:              id,
			StudentName:     studentName,
			Grade:           grade,
			Course:          course,
			StudentAddress:  studentAddress.Hex(),
			ApplicationTime: applicationTime,
			ApprovalTime:    approvalTime,
			Status:          status,
			Approver:        approver.Hex(),
		}
		c.JSON(200, serializer.Success(vo))
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// 申请证书
type ApplyCertificateAo struct {
	Course      string `json:"course" validate:"required"`
	StudentName string `json:"studentName" validate:"required"`
}

func ApplyCertificate(c *gin.Context) {
	ao := new(ApplyCertificateAo)
	if err := c.ShouldBind(ao); err == nil {
		user := CurrentUser(c)
		_, _, err := sdk.Contract.ApplyForCertificate(common.HexToAddress(user.Address), ao.StudentName, ao.Course)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, serializer.Success(nil))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 修改待审核的证书
type UpdatePendingCertificateAo struct {
	Id          *big.Int `json:"id" validate:"required"`
	Course      string   `json:"course" validate:"required"`
	StudentName string   `json:"studentName" validate:"required"`
}

func UpdatePendingCertificate(c *gin.Context) {
	ao := new(UpdatePendingCertificateAo)
	if err := c.ShouldBind(ao); err == nil {
		user := CurrentUser(c)
		_, _, err := sdk.Contract.ModifyPendingCertificate(common.HexToAddress(user.Address), ao.Id, ao.StudentName, ao.Course)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, serializer.Success(nil))
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// 我的证书列表
type MyCertificateListAo struct {
	Course string `json:"course" required:"false" default:""`
	Status string `json:"status" required:"false" default:""`
}

func GetMyCertificateList(c *gin.Context) {
	ao := new(MyCertificateListAo)
	if err := c.ShouldBind(ao); err == nil {
		user := CurrentUser(c)
		certificates, err := sdk.Contract.GetStudentCertificates(common.HexToAddress(user.Address), ao.Course, ao.Status)
		if err != nil {
			util.Log().Error(err.Error())
			c.JSON(200, ErrorResponse(err))
			return
		}
		var vos []CertificateVo
		for _, certificate := range certificates {
			vo := CertificateVo{
				Id:              certificate.Id,
				StudentName:     certificate.StudentName,
				Grade:           certificate.Grade,
				Course:          certificate.Course,
				StudentAddress:  certificate.StudentAddress.Hex(),
				ApplicationTime: certificate.ApplicationTime,
				ApprovalTime:    certificate.ApprovalTime,
				Status:          certificate.Status,
				Approver:        certificate.Approver.Hex(),
			}
			vos = append(vos, vo)
		}
		c.JSON(200, serializer.Success(vos))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 待审核的证书列表
type PendingCertificateListAo struct {
	Course         string `json:"course" required:"false" default:""`
	StudentName    string `json:"studentName" validate:"required"`
	StudentAddress string `json:"studentAddress" validate:"required"`
}

func GetPendingCertificateList(c *gin.Context) {
	ao := new(PendingCertificateListAo)
	if err := c.ShouldBind(ao); err == nil {
		certificates, err := sdk.Contract.GetPendingCertificates(common.HexToAddress(ao.StudentAddress), ao.Course, ao.StudentName)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		var vos []CertificateVo
		for _, certificate := range certificates {
			vo := CertificateVo{
				Id:              certificate.Id,
				StudentName:     certificate.StudentName,
				Grade:           certificate.Grade,
				Course:          certificate.Course,
				StudentAddress:  certificate.StudentAddress.Hex(),
				ApplicationTime: certificate.ApplicationTime,
				ApprovalTime:    certificate.ApprovalTime,
				Status:          certificate.Status,
				Approver:        certificate.Approver.Hex(),
			}
			vos = append(vos, vo)
		}
		c.JSON(200, serializer.Success(vos))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 审核证书
type ApproveCertificateAo struct {
	Id     *big.Int `json:"id" validate:"required"`
	Grade  *big.Int `json:"grade" validate:"required"`
	Status string   `json:"status" validate:"required"`
}

func ApproveCertificate(c *gin.Context) {
	ao := new(ApproveCertificateAo)
	if err := c.ShouldBind(ao); err == nil {
		user := CurrentUser(c)
		if ao.Status == "Approved" {
			_, _, err := sdk.Contract.ApproveCertificate(common.HexToAddress(user.Address), ao.Id, ao.Grade)
			if err != nil {
				c.JSON(200, ErrorResponse(err))
				return
			}
		} else {
			_, _, err := sdk.Contract.RejectCertificate(common.HexToAddress(user.Address), ao.Id)
			if err != nil {
				c.JSON(200, ErrorResponse(err))
				return
			}
		}
		c.JSON(200, serializer.Success(nil))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

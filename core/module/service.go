package module

import (
	"net/http"
	"strings"

	"github.com/nus-utils/nus-peer-review/models"
	"github.com/nus-utils/nus-peer-review/utils"
	"gorm.io/gorm"
)

func (controller ModuleController) ModuleCreateHandleFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		module := r.Context().Value(utils.DecodeBodyContextKey).(*models.Module)
		user := r.Context().Value(utils.JWTClaimContextKey).(*models.User)
		
		txError := controller.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(module).Create(module).Error; err != nil {
				return err
			}
			supervision := &models.Supervision{ModuleID: module.ID, StaffID: user.ID}
			if err := tx.Model(supervision).Create(supervision).Error; err != nil {
				return err
			}
			return nil
		})
		if txError != nil {
			var errMessage string
			if strings.Contains(txError.Error(), "duplicate key value violates unique constraint") {
				errMessage = "Creation failed: Module already exists."
			} else {
				errMessage = "Creation failed: " + txError.Error()
			}
			utils.HandleResponse(w, errMessage, http.StatusBadRequest)
			return
		}
		utils.HandleResponseWithObject(w, module, http.StatusOK)
	}
}

func (controller ModuleController) EnrollmentCreateHandleFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := r.Context().Value(utils.DecodeBodyContextKey).(*models.Enrollment)

		result := controller.DB.Model(data).Create(data)
		if result.Error != nil {
			utils.HandleResponse(w, result.Error.Error(), http.StatusBadRequest)
			return
		}
		utils.HandleResponseWithObject(w, data, http.StatusOK)
	}
}

func (controller ModuleController) SupervisionCreateHandleFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := r.Context().Value(utils.DecodeBodyContextKey).(*models.Supervision)

		result := controller.DB.Model(data).Create(data)
		if result.Error != nil {
			utils.HandleResponse(w, result.Error.Error(), http.StatusBadRequest)
			return
		}
		utils.HandleResponseWithObject(w, data, http.StatusOK)
	}
}

func (controller ModuleController) GetStudentEnrollmentsHandleFunc() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var modules []models.Module
		user := r.Context().Value(utils.JWTClaimContextKey).(*models.User)
		controller.DB.Joins("inner join enrollments e on modules.id = e.module_id").Where("e.student_id = ?", user.ID).Find(&modules)
		utils.HandleResponseWithObject(w, modules, http.StatusOK)
	})
}

func (controller ModuleController) GetStaffSupervisionsHandleFunc() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var modules []models.Module
		user := r.Context().Value(utils.JWTClaimContextKey).(*models.User)
		controller.DB.Joins("inner join supervisions e on modules.id = e.module_id").Where("e.staff_id = ?", user.ID).Find(&modules)
		utils.HandleResponseWithObject(w, modules, http.StatusOK)
	})
}

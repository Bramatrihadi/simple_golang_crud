package services

import (
	"fmt"
	"log"

	"ppid/dtos"
	"ppid/models"
	"ppid/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAdmin(admin *models.Admin, repository repositories.AdminRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	admin.ID = uuidResult.String()

	operationResult := repository.Save(admin)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Admin)

	return dtos.Response{Success: true, Data: data}
}

func FindAllAdmins(repository repositories.AdminRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Admins)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneAdminById(id string, repository repositories.AdminRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Admin)

	return dtos.Response{Success: true, Data: data}
}

func UpdateAdminById(id string, admin *models.Admin, repository repositories.AdminRepository) dtos.Response {
	existingAdminResponse := FindOneAdminById(id, repository)

	if !existingAdminResponse.Success {
		return existingAdminResponse
	}

	existingAdmin := existingAdminResponse.Data.(*models.Admin)

	existingAdmin.Nama = admin.Nama
	existingAdmin.Email = admin.Email

	operationResult := repository.Save(existingAdmin)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneAdminById(id string, repository repositories.AdminRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteAdminByIds(multiId *dtos.MultiID, repository repositories.AdminRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func PaginationAdmin(repository repositories.AdminRepository, context *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := repository.Pagination(pagination)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*dtos.Pagination)

	// get current url path
	urlPath := context.Request.URL.Path

	// search query params
	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 0 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return dtos.Response{Success: true, Data: data}
}

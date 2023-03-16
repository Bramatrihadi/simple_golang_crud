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

func CreateUser(user *models.User, repository repositories.UserRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	user.ID = uuidResult.String()

	operationResult := repository.Save(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	return dtos.Response{Success: true, Data: data}
}

func FindAllUsers(repository repositories.UserRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Users)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneUserById(id string, repository repositories.UserRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	return dtos.Response{Success: true, Data: data}
}

func UpdateUserById(id string, user *models.User, repository repositories.UserRepository) dtos.Response {
	existingUserResponse := FindOneUserById(id, repository)

	if !existingUserResponse.Success {
		return existingUserResponse
	}

	existingUser := existingUserResponse.Data.(*models.User)

	existingUser.Nama_lengkap = user.Nama_lengkap
	existingUser.Email = user.Email
	existingUser.Alamat_lengkap = user.Alamat_lengkap

	operationResult := repository.Save(existingUser)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneUserById(id string, repository repositories.UserRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteUserByIds(multiId *dtos.MultiID, repository repositories.UserRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func Pagination(repository repositories.UserRepository, context *gin.Context, pagination *dtos.Pagination) dtos.Response {
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

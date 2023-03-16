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

func CreatePendaftaran(pendaftaran *models.Pendaftaran, repository repositories.PendaftaranRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	pendaftaran.ID = uuidResult.String()

	operationResult := repository.Save(pendaftaran)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Pendaftaran)

	return dtos.Response{Success: true, Data: data}
}

func FindAllPendaftarans(repository repositories.PendaftaranRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Pendaftarans)

	return dtos.Response{Success: true, Data: datas}
}

func FindOnePendaftaranById(id string, repository repositories.PendaftaranRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Pendaftaran)

	return dtos.Response{Success: true, Data: data}
}

func UpdatePendaftaranById(id string, pendaftaran *models.Pendaftaran, repository repositories.PendaftaranRepository) dtos.Response {
	existingPendaftaranResponse := FindOnePendaftaranById(id, repository)

	if !existingPendaftaranResponse.Success {
		return existingPendaftaranResponse
	}

	existingPendaftaran := existingPendaftaranResponse.Data.(*models.Pendaftaran)

	existingPendaftaran.Nik = pendaftaran.Nik

	operationResult := repository.Save(existingPendaftaran)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOnePendaftaranById(id string, repository repositories.PendaftaranRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeletePendaftaranByIds(multiId *dtos.MultiID, repository repositories.PendaftaranRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func PaginationPendaftaran(repository repositories.PendaftaranRepository, context *gin.Context, pagination *dtos.Pagination) dtos.Response {
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

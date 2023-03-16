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

func CreateDokumen(dokumen *models.Dokumen, repository repositories.DokumenRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	dokumen.ID = uuidResult.String()

	operationResult := repository.Save(dokumen)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Dokumen)

	return dtos.Response{Success: true, Data: data}
}

func FindAllDokumens(repository repositories.DokumenRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Dokumens)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneDokumenById(id string, repository repositories.DokumenRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Dokumen)

	return dtos.Response{Success: true, Data: data}
}

func UpdateDokumenById(id string, dokumen *models.Dokumen, repository repositories.DokumenRepository) dtos.Response {
	existingDokumenResponse := FindOneDokumenById(id, repository)

	if !existingDokumenResponse.Success {
		return existingDokumenResponse
	}

	existingDokumen := existingDokumenResponse.Data.(*models.Dokumen)

	existingDokumen.Jenis = dokumen.Jenis
	existingDokumen.Judul = dokumen.Judul

	operationResult := repository.Save(existingDokumen)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneDokumenById(id string, repository repositories.DokumenRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteDokumenByIds(multiId *dtos.MultiID, repository repositories.DokumenRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func PaginationDokumen(repository repositories.DokumenRepository, context *gin.Context, pagination *dtos.Pagination) dtos.Response {
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

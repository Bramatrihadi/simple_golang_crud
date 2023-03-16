package repositories

import (
	"fmt"
	"math"
	"strings"

	"ppid/dtos"
	"ppid/models"

	"github.com/jinzhu/gorm"
)

type PendaftaranRepository struct {
	db *gorm.DB
}

func NewPendaftaranRepository(db *gorm.DB) *PendaftaranRepository {
	return &PendaftaranRepository{db: db}
}

func (r *PendaftaranRepository) Save(pendaftaran *models.Pendaftaran) RepositoryResult {
	err := r.db.Save(pendaftaran).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: pendaftaran}
}

func (r *PendaftaranRepository) FindAll() RepositoryResult {
	var pendaftarans models.Pendaftarans

	err := r.db.Find(&pendaftarans).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &pendaftarans}
}

func (r *PendaftaranRepository) FindOneById(id string) RepositoryResult {
	var pendaftaran models.Pendaftaran

	err := r.db.Where(&models.Pendaftaran{ID: id}).Take(&pendaftaran).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &pendaftaran}
}

func (r *PendaftaranRepository) DeleteOneById(id string) RepositoryResult {
	err := r.db.Delete(&models.Pendaftaran{ID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *PendaftaranRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("ID IN (?)", *ids).Delete(&models.Pendaftarans{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *PendaftaranRepository) Pagination(pagination *dtos.Pagination) (RepositoryResult, int) {
	var pendaftarans models.Pendaftarans

	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break

			}
		}
	}

	find = find.Find(&pendaftarans)

	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = pendaftarans

	// count all data
	errCount := r.db.Model(&models.Pendaftaran{}).Count(&totalRows).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > totalRows {
		// set to row with total rows
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}

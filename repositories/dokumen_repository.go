package repositories

import (
	"fmt"
	"math"
	"strings"

	"ppid/dtos"
	"ppid/models"

	"github.com/jinzhu/gorm"
)

type DokumenRepository struct {
	db *gorm.DB
}

func NewDokumenRepository(db *gorm.DB) *DokumenRepository {
	return &DokumenRepository{db: db}
}

func (r *DokumenRepository) Save(dokumen *models.Dokumen) RepositoryResult {
	err := r.db.Save(dokumen).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: dokumen}
}

func (r *DokumenRepository) FindAll() RepositoryResult {
	var dokumens models.Dokumens

	err := r.db.Find(&dokumens).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &dokumens}
}

func (r *DokumenRepository) FindOneById(id string) RepositoryResult {
	var dokumen models.Dokumen

	err := r.db.Where(&models.Dokumen{ID: id}).Take(&dokumen).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &dokumen}
}

func (r *DokumenRepository) DeleteOneById(id string) RepositoryResult {
	err := r.db.Delete(&models.Dokumen{ID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *DokumenRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("ID IN (?)", *ids).Delete(&models.Dokumens{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *DokumenRepository) Pagination(pagination *dtos.Pagination) (RepositoryResult, int) {
	var dokumens models.Dokumens

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

	find = find.Find(&dokumens)

	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = dokumens

	// count all data
	errCount := r.db.Model(&models.User{}).Count(&totalRows).Error

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

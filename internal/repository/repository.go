package repository

import (
	"GeoServiseAppDate/internal/models"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type Repository interface {
	GetDataAddress(request *models.SearchRequest) ([]*models.AddressSearch, error)
	CheckCacheAddress(request *models.SearchRequest) (bool, error)
	AddDataAddressToDB(request *models.SearchRequest, addresses []*models.AddressSearch) error

	GetDataGEO(request *models.GeocodeRequest) (*models.AddressGeo, error)
	CheckCacheGEO(request *models.GeocodeRequest) (bool, error)
	AddDataGEOToDB(request *models.GeocodeRequest, geo *models.AddressGeo) error
}

type repository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func New(database *sql.DB) Repository {
	return &repository{
		db:         database,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *repository) GetDataAddress(request *models.SearchRequest) ([]*models.AddressSearch, error) {
	var results []*models.AddressSearch

	query := r.sqlBuilder.Select("data").
		From("address_data").Where(sq.Eq{"address": request.Query})

	row := query.RunWith(r.db).QueryRow()

	if err := row.Scan(&results); err != nil {
		return nil, err
	}
	return results, nil

}

func (r *repository) CheckCacheAddress(request *models.SearchRequest) (bool, error) {
	query := r.sqlBuilder.Select("COUNT(*)").
		From("address_data").Where(sq.Eq{"address": request.Query})

	row := query.RunWith(r.db).QueryRow()
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *repository) AddDataAddressToDB(request *models.SearchRequest, addresses []*models.AddressSearch) error {
	query := r.sqlBuilder.Insert("address_data").
		Columns("address", "data").
		Values(request.Query, addresses)

	if _, err := query.RunWith(r.db).Exec(); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetDataGEO(request *models.GeocodeRequest) (*models.AddressGeo, error) {
	var result *models.AddressGeo
	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)
	query := r.sqlBuilder.Select("data").
		From("geo_data").Where(sq.Eq{"geo": geoRequest})

	row := query.RunWith(r.db).QueryRow()

	if err := row.Scan(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) CheckCacheGEO(request *models.GeocodeRequest) (bool, error) {
	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)
	query := r.sqlBuilder.Select("COUNT(*)").
		From("geo_data").Where(sq.Eq{"geo": geoRequest})

	row := query.RunWith(r.db).QueryRow()
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *repository) AddDataGEOToDB(request *models.GeocodeRequest, geo *models.AddressGeo) error {
	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)
	query := r.sqlBuilder.Insert("geo_data").
		Columns("geo", "data").
		Values(geoRequest, geo)

	if _, err := query.RunWith(r.db).Exec(); err != nil {
		return err
	}

	return nil
}

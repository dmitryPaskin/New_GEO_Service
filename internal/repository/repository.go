package repository

import (
	"GeoServiseAppDate/internal/metrics"
	"GeoServiseAppDate/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/prometheus/client_golang/prometheus"
	"time"
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
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "GetDataAddress"}).
			Observe(duration)
	}()

	var result []*models.AddressSearch
	var resultString string
	query := r.sqlBuilder.Select("data").
		From("address_data").Where(sq.Eq{"address": request.Query})

	row := query.RunWith(r.db).QueryRow()

	if err := row.Scan(&resultString); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(resultString), &result); err != nil {
		return nil, err
	}
	return result, nil

}

func (r *repository) CheckCacheAddress(request *models.SearchRequest) (bool, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "CheckCacheAddress"}).
			Observe(duration)
	}()

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
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "AddDataAddressToDB"}).
			Observe(duration)
	}()

	data, err := json.Marshal(addresses)
	if err != nil {
		return err
	}

	query := r.sqlBuilder.Insert("address_data").
		Columns("address", "data").
		Values(request.Query, string(data))

	if _, err := query.RunWith(r.db).Exec(); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetDataGEO(request *models.GeocodeRequest) (*models.AddressGeo, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "GetDataGEO"}).
			Observe(duration)
	}()

	var result *models.AddressGeo
	var resultString string
	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)
	query := r.sqlBuilder.Select("data").
		From("geo_data").Where(sq.Eq{"geo": geoRequest})

	row := query.RunWith(r.db).QueryRow()

	if err := row.Scan(&resultString); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(resultString), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) CheckCacheGEO(request *models.GeocodeRequest) (bool, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "CheckCacheGEO"}).
			Observe(duration)
	}()

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
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "AddDataGEOToBD"}).
			Observe(duration)
	}()

	data, err := json.Marshal(geo)
	if err != nil {
		return err
	}

	geoRequest := fmt.Sprintf("%s, %s", request.Lon, request.Lat)
	query := r.sqlBuilder.Insert("geo_data").
		Columns("geo", "data").
		Values(geoRequest, data)

	if _, err := query.RunWith(r.db).Exec(); err != nil {
		return err
	}

	return nil
}

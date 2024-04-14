package authRepository

import (
	"GeoServiseAppDate/internal/metrics"
	"GeoServiseAppDate/internal/models"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type AuthRepository interface {
	SaveUser(user models.User) error
	CheckUser(user models.User) (bool, error)
	GetUser(user models.User) (models.User, error)
}

type authRepository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func New(database *sql.DB) AuthRepository {
	return &authRepository{
		db:         database,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (ar *authRepository) SaveUser(user models.User) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "SaveUser"}).
			Observe(duration)
	}()

	query := ar.sqlBuilder.Insert("users").
		Columns("login", "password").
		Values(user.Login, user.Password)

	if _, err := query.RunWith(ar.db).Exec(); err != nil {
		return err
	}

	return nil
}

func (ar *authRepository) CheckUser(user models.User) (bool, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "CheckUser"}).
			Observe(duration)
	}()

	query := ar.sqlBuilder.Select("COUNT(*)").
		From("users").
		Where(sq.Eq{"login": user.Login})

	row := query.RunWith(ar.db).QueryRow()
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (ar *authRepository) GetUser(user models.User) (models.User, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"function": "GetUser"}).
			Observe(duration)
	}()

	query := ar.sqlBuilder.Select("login", "password").
		From("users").Where(sq.Eq{"login": user.Login})

	row := query.RunWith(ar.db).QueryRow()
	newUser := models.User{}
	if err := row.Scan(&newUser.Login, &newUser.Password); err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

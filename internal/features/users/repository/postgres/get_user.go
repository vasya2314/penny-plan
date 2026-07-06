package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/vasya2314/penny-plan/internal/core/domain"
	core_errors "github.com/vasya2314/penny-plan/internal/core/errors"
	core_postgres_pool "github.com/vasya2314/penny-plan/internal/core/repository/postgres/pool"
)

func (s *UsersRepository) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM penny_plan.users
	WHERE id = $1;
	`

	row := s.pool.QueryRow(ctx, query, id)

	var userModel domain.User

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan row error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}

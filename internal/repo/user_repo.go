package repo

import (
	"database/sql"
	"errors"
	"transaction_ex1/internal/entity"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindUser(userId int) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, username, balance, version_update FROM users WHERE id = $1", userId)

	var (
		id            int
		username      string
		balance       int
		versionUpdate int
	)

	if err := row.Scan(&id, &username, &balance, &versionUpdate); err != nil {
		return nil, err
	}

	user := entity.NewUser(id, username, balance, versionUpdate)

	return &user, nil
}

func (r *UserRepo) Save(user *entity.User) error {
	expectedVersion := user.GetVersionUpdate()
	updatedVersion := expectedVersion + 1

	// тут главная идея в version_update
	// при получении пользователя через FindUser мы получаем версию пользователя из базы данных
	// при обновлении пользователя мы увеличиваем версию пользователя на единицу
	// при этом в where подставляется версия пользователя из базы данных, а в set - новая версия пользователя
	// таким образом мы гарантируем, что пользователь с такой же версией не будет обновляться
	result, err := r.db.Exec(
		"UPDATE users SET balance = $1, version_update = $2 WHERE id = $3 AND version_update = $4",
		user.GetBalance(), updatedVersion, user.GetId(), expectedVersion,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Ошибка при обновлении пользователя")
	}

	return nil
}

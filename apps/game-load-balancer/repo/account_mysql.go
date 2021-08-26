package repo

import (
	"context"
	"database/sql"
)

type accountMySQLRepo struct {
	db *sql.DB

	accountByUserStmt *sql.Stmt
	updateAccountStmt *sql.Stmt
}

func NewAccountMySQLRepo(db *sql.DB) (AccountRepo, error) {
	accountByUserStmt, err := db.Prepare("SELECT id, username, salt, verifier, session_key_auth, locked, last_ip FROM account WHERE username = ?")
	if err != nil {
		return nil, err
	}

	updateAccountStmt, err := db.Prepare("UPDATE account SET username = ?, salt = ?, verifier = ?, session_key_auth = ?, locked = ?, last_ip = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	return &accountMySQLRepo{
		db:                db,
		accountByUserStmt: accountByUserStmt,
		updateAccountStmt: updateAccountStmt,
	}, nil
}

func (r *accountMySQLRepo) AccountByUserName(ctx context.Context, username string) (*Account, error) {
	account := &Account{}
	row := r.accountByUserStmt.QueryRowContext(ctx, username)
	err := row.Scan(&account.ID, &account.Username, &account.Salt, &account.Verifier, &account.SessionKeyAuth, &account.Locked, &account.LastIP)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return account, nil
}

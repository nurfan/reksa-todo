package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/nurfan/reksa-todo/model"
)

// GetReminderLog ...
func (r *RepositoryPsql) GetUser(ctx context.Context, username string) (result model.Users, err error) {

	query := `SELECT
					*
				FROM 
					m_branch_user bu
				WHERE
					bu.email = ?`

	err = r.Conn.GetContext(ctx, &result, query, username)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("repo.GetAccount : %s", err)
		return
	}

	return
}

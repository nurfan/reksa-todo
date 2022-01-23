package model

import "database/sql"

type Account struct {
	UserID      int64          `db:"user_id"`
	Username    string         `db:"username"`
	Password    string         `db:"password"`
	Nickname    string         `db:"nickname"`
	Email       string         `db:"email"`
	SectionID   int32          `db:"section_id"`
	SectionName string         `db:"section_name"`
	RoleID      int32          `db:"role_id"`
	RoleName    string         `db:"role_name"`
	IsActive    bool           `db:"is_active"`
	LastLogin   sql.NullString `db:"last_login"`
	CreatedBy   int64          `db:"created_by"`
	CreatedAt   string         `db:"created_at"`
	UpdatedBy   int64          `db:"updated_by"`
	UpdatedAt   string         `db:"updated_at"`
}

type Users struct {
	UserID        int64  `db:"branch_user_id"`
	UserType      int32  `db:"branch_user_type"`
	BranchTypeID  int32  `db:"branch_type_id"`
	MksUserID     int64  `db:"mks_user_id"`
	Email         string `db:"email"`
	Password      string `db:"password"`
	Username      string `db:"user_name"`
	Phone         string `db:"no_hp"`
	UserStatus    int32  `db:"user_status"`
	KodeRa        string `db:"kode_ra"`
	ProfileID     string `db:"profile_id"`
	Remarks       string `db:"remarks"`
	DatetimeAdded string `db:"datetime_added"`
	DatetimeLogin string `db:"datetime_login"`
}

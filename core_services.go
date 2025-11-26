package platform

import (
	"context"
)

// ---------- DB ----------

type DBConn interface {
	Exec(ctx context.Context, query string, args ...any) error
}

type DatabaseService interface {
	WithTenant(ctx context.Context) DBConn
}

// ---------- File ----------

type FileStorageService interface {
	Upload(ctx context.Context, filename string, data []byte) (string, error)
}

// ---------- Auth ----------

type AuthzService interface {
	Can(ctx context.Context, action string, userID int) bool
}

// ---------- User ----------

type UserService interface {
	GetUserName(ctx context.Context, userID int) (string, error)
}

// ---------- CoreServices (플러그인이 주입받는 묶음) ----------

type CoreServices struct {
	DB   func(ctx context.Context) DBConn
	File func(ctx context.Context) FileStorageService
	Auth AuthzService
	User UserService
}

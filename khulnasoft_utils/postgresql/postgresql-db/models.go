// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package postgresql_db

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ApiToken struct {
	ID              int64     `json:"id"`
	ApiToken        uuid.UUID `json:"api_token"`
	Name            string    `json:"name"`
	CompanyID       int32     `json:"company_id"`
	GroupID         int32     `json:"group_id"`
	RoleID          int32     `json:"role_id"`
	CreatedByUserID int64     `json:"created_by_user_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type AuditLog struct {
	ID        int64     `json:"id"`
	Event     string    `json:"event"`
	Action    string    `json:"action"`
	Resources string    `json:"resources"`
	Success   bool      `json:"success"`
	UserEmail string    `json:"user_email"`
	UserRole  string    `json:"user_role"`
	CreatedAt time.Time `json:"created_at"`
}

type Company struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	EmailDomain string    `json:"email_domain"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Namespace   string    `json:"namespace"`
}

type ContainerRegistry struct {
	ID              int32           `json:"id"`
	Name            string          `json:"name"`
	RegistryType    string          `json:"registry_type"`
	EncryptedSecret json.RawMessage `json:"encrypted_secret"`
	NonSecret       json.RawMessage `json:"non_secret"`
	Extras          json.RawMessage `json:"extras"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type GenerativeAiIntegration struct {
	ID                 int32           `json:"id"`
	IntegrationType    string          `json:"integration_type"`
	Label              string          `json:"label"`
	LastSentTime       sql.NullTime    `json:"last_sent_time"`
	Config             json.RawMessage `json:"config"`
	ErrorMsg           sql.NullString  `json:"error_msg"`
	DefaultIntegration bool            `json:"default_integration"`
	CreatedByUserID    int64           `json:"created_by_user_id"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

type Integration struct {
	ID              int32           `json:"id"`
	Resource        string          `json:"resource"`
	Filters         json.RawMessage `json:"filters"`
	IntegrationType string          `json:"integration_type"`
	IntervalMinutes int32           `json:"interval_minutes"`
	LastSentTime    sql.NullTime    `json:"last_sent_time"`
	Config          json.RawMessage `json:"config"`
	ErrorMsg        sql.NullString  `json:"error_msg"`
	CreatedByUserID int64           `json:"created_by_user_id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type License struct {
	ID                              int32           `json:"id"`
	LicenseKey                      uuid.UUID       `json:"license_key"`
	StartDate                       time.Time       `json:"start_date"`
	EndDate                         time.Time       `json:"end_date"`
	NoOfHosts                       int64           `json:"no_of_hosts"`
	CurrentHosts                    int64           `json:"current_hosts"`
	IsActive                        bool            `json:"is_active"`
	LicenseType                     string          `json:"license_type"`
	KhulnasoftSupportEmail           string          `json:"khulnasoft_support_email"`
	NotificationThresholdPercentage int32           `json:"notification_threshold_percentage"`
	NotificationThresholdUpdatedAt  sql.NullTime    `json:"notification_threshold_updated_at"`
	RegistryCredentials             json.RawMessage `json:"registry_credentials"`
	Message                         string          `json:"message"`
	Description                     string          `json:"description"`
	CreatedAt                       time.Time       `json:"created_at"`
	UpdatedAt                       time.Time       `json:"updated_at"`
	LicenseEmail                    string          `json:"license_email"`
	LicenseEmailDomain              string          `json:"license_email_domain"`
	NoOfCloudAccounts               int64           `json:"no_of_cloud_accounts"`
	NoOfRegistries                  int64           `json:"no_of_registries"`
	NoOfImagesInRegistry            int64           `json:"no_of_images_in_registry"`
}

type PasswordReset struct {
	ID        int32     `json:"id"`
	UserID    int64     `json:"user_id"`
	Code      uuid.UUID `json:"code"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Scheduler struct {
	ID          int64           `json:"id"`
	Action      string          `json:"action"`
	Description string          `json:"description"`
	CronExpr    string          `json:"cron_expr"`
	Payload     json.RawMessage `json:"payload"`
	IsEnabled   bool            `json:"is_enabled"`
	IsSystem    bool            `json:"is_system"`
	Status      string          `json:"status"`
	LastRanAt   sql.NullTime    `json:"last_ran_at"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type Setting struct {
	ID            int64           `json:"id"`
	Key           string          `json:"key"`
	Value         json.RawMessage `json:"value"`
	IsVisibleOnUi bool            `json:"is_visible_on_ui"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type User struct {
	ID                  int64           `json:"id"`
	FirstName           string          `json:"first_name"`
	LastName            string          `json:"last_name"`
	Email               string          `json:"email"`
	RoleID              int32           `json:"role_id"`
	GroupIds            json.RawMessage `json:"group_ids"`
	CompanyID           int32           `json:"company_id"`
	PasswordHash        string          `json:"password_hash"`
	IsActive            bool            `json:"is_active"`
	PasswordInvalidated bool            `json:"password_invalidated"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
}

type UserGroup struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	IsSystem  bool      `json:"is_system"`
	CompanyID int32     `json:"company_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInvite struct {
	ID              int32     `json:"id"`
	Email           string    `json:"email"`
	Code            uuid.UUID `json:"code"`
	CreatedByUserID int64     `json:"created_by_user_id"`
	RoleID          int32     `json:"role_id"`
	CompanyID       int32     `json:"company_id"`
	Accepted        bool      `json:"accepted"`
	Expiry          time.Time `json:"expiry"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

package email

import (
	"github.com/khulnasoft/kengine/khulnasoft_server/reporters"
	"github.com/go-playground/validator/v10"
)

type Email struct {
	Config           Config                  `json:"config"`
	IntegrationType  string                  `json:"integration_type"`
	NotificationType string                  `json:"notification_type"`
	Filters          reporters.FieldsFilters `json:"filters"`
	Message          string                  `json:"message"`
	Resource         string                  `json:"resource"`
}

type Config struct {
	EmailID string `json:"email_id"  validate:"required,email" required:"true"`
}

func (e Email) ValidateConfig(validate *validator.Validate) error {
	return validate.Struct(e.Config)
}

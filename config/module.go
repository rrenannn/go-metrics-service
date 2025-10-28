package config

import (
	"hungtech-go/infra/crypt"
	"hungtech-go/internal/scheduling"
	"hungtech-go/internal/users"

	"go.uber.org/fx"
)

var AllModules = fx.Options(
	crypt.Module,
	scheduling.Module,
	users.Module,
)

package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

var (
	// Mode project mode: development/staging/production
	// RuntimeRoot runtime root path prefix
	Mode        string
	RuntimeRoot string

	Tracer         *trace.TracerProvider
	Mysql          *gorm.DB
	Redis          redis.UniversalClient
	CasbinEnforcer *casbin.Enforcer
)

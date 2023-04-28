package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nft_platform/config"
)

var Conf *config.Config
var Logger *zap.Logger
var SLogger *zap.SugaredLogger
var Rdb *redis.Client
var DB *gorm.DB

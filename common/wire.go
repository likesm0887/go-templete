package common

import (
	"CouchSpace/configs"
	cache2 "CouchSpace/infrastructure/cache"
	"CouchSpace/infrastructure/db"
	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(configs.NewConfig)

var DbSet = wire.NewSet(db.NewMongoDb, wire.Bind(new(db.IMongoDb), new(*db.MongoDb)))

var CacheSet = wire.NewSet(cache2.NewInmemoryCache, wire.Bind(new(cache2.ICachePeer), new(*cache2.InmemoryCache)))

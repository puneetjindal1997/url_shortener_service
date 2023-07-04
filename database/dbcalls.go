package database

import (
	"context"
	"urlshortner/constant"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)
	return result.InsertedID, err
}

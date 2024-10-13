package system

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ricky97gr/homeOnline/internal/webservice/database/mongo"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
)

func GetOperationLog(q *paginate.PageQuery) ([]model.OperationLogInfo, int64, error) {
	collection, err := mongo.GetMongoClient("operation_log")
	if err != nil {

		return nil, 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, count, err
	}
	opt := &options.FindOptions{}
	opt.SetLimit(int64(q.PageSize))
	opt.SetSkip(int64((q.Page - 1) * q.PageSize))
	opt.SetSort(bson.D{{
		"createTime", -1}})
	result, err := collection.Find(ctx, bson.D{}, opt)
	if err != nil {

		return nil, 0, err
	}
	var logs []model.OperationLogInfo
	err = result.All(ctx, &logs)
	if err != nil {
		return nil, 0, err
	}
	return logs, count, err

}

func GetSystemLog(q *paginate.PageQuery) ([]model.OperationLogInfo, int64, error) {
	collection, err := mongo.GetMongoClient("system_log")
	if err != nil {

		return nil, 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, count, err
	}
	opt := &options.FindOptions{}
	opt.SetLimit(int64(q.PageSize))
	opt.SetSkip(int64((q.Page - 1) * q.PageSize))
	opt.SetSort(bson.D{{
		"createTime", -1}})
	result, err := collection.Find(ctx, bson.D{}, opt)
	if err != nil {

		return nil, 0, err
	}
	var logs []model.OperationLogInfo
	err = result.All(ctx, &logs)
	if err != nil {
		return nil, 0, err
	}
	return logs, count, err

}

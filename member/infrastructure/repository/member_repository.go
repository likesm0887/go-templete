package database

import (
	"awesomeProject1/member/adapter/repository/data"
	db "awesomeProject1/member/infrastructure"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberRepository struct {
	DB          db.IMongoDb
	TimePattern string "2006-01-02 15:04:05"
	collection  *mongo.Collection
}

func NewMemberRepository(db db.IMongoDb) *MemberRepository {
	return &MemberRepository{
		DB:         db,
		collection: db.GetDb("information").Collection("user"),
	}
}


func (memberRepository *MemberRepository) GetUserInfo(userId string) (*data.Information, error) {
	ctx := context.TODO()
	result := &data.Information{}
	err := memberRepository.collection.FindOne(ctx, bson.M{"_id": userId}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (memberRepository *MemberRepository) AddUserInformation(member data.Information) (*string, error) {
	ctx := context.TODO()
	_, err := memberRepository.collection.InsertOne(ctx, member)
	if err != nil {
		return nil, err
	}
	return &member.UserID, nil
}


func (memberRepository *MemberRepository) UpdateMemberInfo(data data.Information) error {
	ctx := context.TODO()
	result := memberRepository.collection.FindOneAndReplace(ctx, bson.M{"_id": data.UserID}, data)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (memberRepository *MemberRepository) GetAll() (*[]data.Information, error) {
	ctx := context.Background()
	cursor, err := memberRepository.collection.Find(ctx, bson.D{})
	var result []data.Information
	err = cursor.All(ctx, result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

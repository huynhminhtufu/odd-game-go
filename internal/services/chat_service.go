package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/oddx-team/odd-game-server/internal/constants"
	"github.com/oddx-team/odd-game-server/pb"
)

func (s *Service) GetChats(ctx context.Context, _ *pb.GetChatsRequest) (*pb.GetChatsResponse, error) {
	chatCollection := s.mongo.Collection(constants.ChatCollection)
	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"$natural": -1})
	findOptions.SetLimit(50)
	cur, err := chatCollection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*pb.ChatEntity
	for cur.Next(ctx) {
		var result pb.ChatEntity
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &pb.GetChatsResponse{
		Chats: results,
	}, nil
}

func (s *Service) InsertChat(ctx context.Context, newChat *pb.ChatEntity) (*pb.InsertChatResponse, error) {
	chatCollection := s.mongo.Collection(constants.ChatCollection)

	_, err := chatCollection.InsertOne(ctx, newChat)
	if err != nil {
		return nil, err
	}

	return &pb.InsertChatResponse{
		Success: true,
	}, nil
}

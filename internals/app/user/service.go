package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"microservice-one/grpcproto"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type service struct {
	repo             Repository
	redisClient      *redis.Client
	serviceTwoClient grpcproto.MicroServiceTwoServiceClient
}

type Service interface {
	RegisterUser(ctx context.Context, user *User) (int64, error)
	GetUserProfileDetails(ctx context.Context, userId int) (*UserProfileDetails, error)
	UpdateUserProfile(ctx context.Context, id int, user UserProfileDetails) error
	DeleteUserProfile(ctx context.Context, id int) error
	ListUsers(ctx context.Context, listUserReq ListUserRequest) ([]string, error)
	SetToCache(ctx context.Context, cacheKey string, data interface{}, expiration time.Duration) error
	GetFromCache(ctx context.Context, cacheKey string, result interface{}) error
}

func NewService(repo Repository, redisClient *redis.Client, serviceTwoClient grpcproto.MicroServiceTwoServiceClient) Service {
	return &service{
		repo:             repo,
		redisClient:      redisClient,
		serviceTwoClient: serviceTwoClient,
	}
}

func (s *service) GetFromCache(ctx context.Context, cacheKey string, result interface{}) error {
	val, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal cache data: %w", err)
	}

	return nil
}

func (s *service) SetToCache(ctx context.Context, cacheKey string, data interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data for cache: %w", err)
	}

	err = s.redisClient.Set(ctx, cacheKey, jsonData, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set cache: %w", err)
	}

	return nil
}

// User
func (s *service) RegisterUser(ctx context.Context, user *User) (int64, error) {
	res, err := s.repo.GetUserByEmail(ctx, user.Email)
	if res != nil && err == nil {
		return 0, errors.New("email already exist")
	}
	hashPass := HashPassword(user.Password)
	user.Password = hashPass

	userId, err := s.repo.CreateUser(ctx, *user)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *service) GetUserProfileDetails(ctx context.Context, userId int) (*UserProfileDetails, error) {
	cacheKey := fmt.Sprintf("user:%d", userId)
	var cacheUser User
	err := s.GetFromCache(ctx, cacheKey, &cacheUser)
	if err == nil {
		fmt.Println("user found in cache")
		return &UserProfileDetails{
			Username:    cacheUser.Username,
			FirstName:   cacheUser.FirstName,
			LastName:    cacheUser.LastName,
			PhoneNumber: cacheUser.PhoneNumber,
			Email:       cacheUser.Email,
			DateOfBirth: cacheUser.DateOfBirth,
			Gender:      cacheUser.Gender,
		}, nil
	} else if err != redis.Nil {
		fmt.Printf("Redis error: %v\n", err)
	}
	user, err := s.repo.GetUserDetailsById(ctx, userId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("movie not found with the id %d", userId)
	}
	err = s.SetToCache(ctx, cacheKey, user, 10*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache movie details: %v\n", err)
	}
	return &UserProfileDetails{
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
	}, nil
}

func (s *service) UpdateUserProfile(ctx context.Context, id int, user UserProfileDetails) error {
	err := s.repo.UpdateUserProfile(ctx, user, id)
	if err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("user:%d", id)
	err = s.redisClient.Del(ctx, cacheKey).Err()
	if err != nil {
		fmt.Printf("Failed to invalidate cache for user %d: %v\n", id, err)
	}
	updatedUser, err := s.repo.GetUserDetailsById(ctx, id)
	if err != nil {
		return err
	}
	err = s.SetToCache(ctx, cacheKey, updatedUser, time.Hour)
	if err != nil {
		fmt.Printf("Failed to set updated cache for user %d: %v\n", id, err)
		return fmt.Errorf("failed to update cache for user: %w", err)
	}
	return nil
}

func (s *service) DeleteUserProfile(ctx context.Context, id int) error {
	err := s.repo.DeleteUserProfile(ctx, int64(id))
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	cacheKey := fmt.Sprintf("user:%d", id)
	err = s.redisClient.Del(ctx, cacheKey).Err()
	if err != nil {
		fmt.Printf("Failed to invalidate cache for user %d: %v\n", id, err)
	}
	return nil
}
func (s *service) ListUsers(ctx context.Context, listUserReq ListUserRequest) ([]string, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	names := make([]*grpcproto.Name, len(users))
	for i, user := range users {
		names[i] = &grpcproto.Name{Name: user.Username}
	}

	var res *grpcproto.MethodResponse
	switch listUserReq.Method {
	case 1:
		res, err = s.serviceTwoClient.MethodOne(ctx, &grpcproto.MethodRequest{
			MethodNumber: listUserReq.Method,
			WaitTime:     listUserReq.WaitTime,
			Names:        names,
		})
	case 2:
		res, err = s.serviceTwoClient.MethodTwo(ctx, &grpcproto.MethodRequest{
			MethodNumber: listUserReq.Method,
			WaitTime:     listUserReq.WaitTime,
			Names:        names,
		})
	default:
		return nil, errors.New("invalid method number")
	}

	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	userNames := make([]string, len(res.UserNames))
	for i, username := range res.UserNames {
		userNames[i] = username.Name
	}

	if len(userNames) == 0 {
		return nil, errors.New("no users found")
	}
	return userNames, nil
}

package pkg

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go-ibooking/src/pkg/logger"
	"go-ibooking/src/utils"

	"github.com/go-redis/redis/v8"
)

type RedisDatastore struct {
	Client redis.UniversalClient
}

func NewRedisDatastore() *RedisDatastore {
	isClusterEnabled := os.Getenv("REDIS_CLUSTER_ENABLED") == "true"
	var client redis.UniversalClient

	if isClusterEnabled {
		client = connectRedisCluster()
	} else {
		client = connectRedisStandalone()
	}

	return &RedisDatastore{Client: client}
}

func (r *RedisDatastore) Close() error {
	err := r.Client.Close()

	if err != nil {
		return err
	}

	return nil
}

func connectRedisStandalone() *redis.Client {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	db := os.Getenv("REDIS_DB")

	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       utils.StringToInt(db),
	}

	client := redis.NewClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Log.Err(err).Msg("failed to connect to Redis (Standalone)")
		log.Printf("Failed to connect to Redis (Standalone): %v", err)
	}

	logger.Write.Info().Msg("Successfully connected to Redis (Standalone)")
	return client
}

func connectRedisCluster() *redis.ClusterClient {
	nodes := os.Getenv("REDIS_CLUSTER_NODES")
	addrs := strings.Split(nodes, ",")
	password := os.Getenv("REDIS_PASSWORD")

	options := &redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	}

	client := redis.NewClusterClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Log.Err(err).Msg("failed to connect to Redis Cluster")
		log.Printf("Failed to connect to Redis Cluster: %v", err)
	}

	logger.Write.Info().Msg("Successfully connected to Redis Cluster")
	return client
}

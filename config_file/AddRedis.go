package configfile

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func AddRedis() (string, error) {

	redisCmd := exec.Command("go", "get", "github.com/go-redis/redis/v8")

	fmt.Println("Getting redis client from: github.com/go-redis/redis/v8")

	redisCmdErr := redisCmd.Run()

	if redisCmdErr != nil {
		return "", errors.New(redisCmdErr.Error())
	}

	makeRedisFile, makeErr := os.Create("config/redis.go")

	if makeErr != nil {
		return "", errors.New(makeErr.Error())
	}

	defer makeRedisFile.Close()

	_, writeErr := makeRedisFile.WriteString(RedisTemplate)
	
	if writeErr != nil {
		return "", errors.New(writeErr.Error())
	}

	return "Redis added to project !", nil
}
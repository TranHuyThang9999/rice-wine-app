package rgdb

//var (
//	rdb *redis.Client
//)
//
//func ConnectRedis() error {
//	rdb = redis.NewClient(&redis.Options{
//		Addr:     configs.Get().AddressRedis,
//		Password: configs.Get().PasswordRedis,
//		DB:       0,
//	})
//
//	_, err := rdb.Ping(context.Background()).Result()
//	if err != nil {
//		return fmt.Errorf("could not connect to Redis: %v", err)
//	}
//
//	fmt.Println("Connected to Redis")
//	return nil
//}

//func GetRedisClient() *redis.Client {
//	return rdb
//}

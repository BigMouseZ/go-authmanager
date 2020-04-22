package redisitem

import (
	`fmt`
	`github.com/gomodule/redigo/redis`
	`go-authmanager/cache`
)

/**
 * @param key redis关键字
 * @return Object
 * @Description 功能：通过key直接获取value值
 **/
func GetByKey(k string) (interface{}, error) {
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	v, err := rc.Do("GET", k)
	if err != nil {
		return nil, err
	}
	return v, nil

}

/**
 * @param key   redis关键字
 * @param value redis值
 * @return boolean
 * @Description 功能：直接通过key存储value
 **/
func SetByKey(k string, v interface{}) bool {
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	_, err := rc.Do("SET", k, v)
	if err != nil {
		fmt.Println("redis SET failed:", err)
		return false
	}
	return true
}

/**
 * @param key redis关键字
 * @return boolean
 * @Description 功能：判断key值是否从存在
 **/
func Exists(k string) (bool, error) {
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	//检查是否存在key值
	exists, err := redis.Bool(rc.Do("EXISTS", k))
	if err != nil {
		fmt.Println("error:", err)
		exists = false
	} else {
		fmt.Printf("exists or not: %v \n", exists)
	}
	return exists, err
}

/**
 * @param key redis关键字
 * @return boolean
 * @Description 功能：删除key
 **/
func DeleteByKey(k string) bool {
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	//检查是否存在key值
	_, err := redis.Bool(rc.Do("DEL", k))
	if err != nil {
		fmt.Println("error:", err)
		return false
	}
	return true
}

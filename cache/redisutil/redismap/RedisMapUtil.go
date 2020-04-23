package redismap

import (
	`fmt`
	`github.com/gomodule/redigo/redis`
	`go-authmanager/cache`
)

/**
 * @param table 表名
 * @param key   key值
 * @return T
 * @Description 功能： 通过table获取其中key值的value
 */
func GetByKey(table, key string) (interface{}, error) {
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	v, err := rc.Do("HGET", table, key)
	if err != nil {
		return nil, err
	}
	return v, nil
}

/**
 * @param table  表名
 * @param key    key值
 * @param object value值
 * @return T
 * @Description 功能：通过table设置某key值的value
 */
func SetByKey(table, key string, value interface{}) bool {
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	_, err := rc.Do("HSET", table, key, value)
	if err != nil {
		fmt.Println("redis HSET failed:", err)
		return false
	}
	return true
}

/**
 * @param table 表名
 * @param key   表中的key值
 * @return boolean
 * @Description 功能：删除table 的key数据
 */
func DeleteByKey(table, key string) bool {
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	_, err := rc.Do("HDEL", table, key)
	if err != nil {
		fmt.Println("redis HDEL failed:", err)
		return false
	}
	return true
}

func DeleteTable(table string) bool {
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	resKeys, err := redis.Strings(rc.Do("hkeys", table))
	if err != nil {
		fmt.Println("hkeys failed", err.Error())
	} else {
		//fmt.Printf("myhash's keys is :")
		for _, v := range resKeys {
			_, err := rc.Do("HDEL", table, v)
			//fmt.Println("键值：" + v)
			if err != nil {
				fmt.Println("redis HDEL failed:", err)
				return false
			}
		}
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
	exists, err := rc.Do("EXISTS", k)
	if err != nil {
		fmt.Println("illegal exception")
	}
	fmt.Printf("exists or not: %v \n", exists)
	return true, err
}

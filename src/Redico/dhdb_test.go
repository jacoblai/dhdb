package Redico_test

import (
	"github.com/garyburd/redigo/redis"
	"testing"
	"fmt"
	"time"
)

var (
	RedisClient     *redis.Pool
)

func TestRedico(t *testing.T) {
	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:  5,
		MaxActive:   5,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6380")
			if err != nil {
				return nil, err
			}
			_, err = c.Do("AUTH", "icoolpy.com")
			fmt.Println(err)
			return c, nil
		},
	}
	//c, err := redis.Dial("tcp", "127.0.0.1:6380")
	//if err != nil {
	//	t.Error(err)
	//}

	//_, err = c.Do("AUTH", "foo", "bar")
	//fmt.Println(err != nil, "no password set")
	//
	//_, err = c.Do("PING", "foo", "bar")
	//fmt.Println(err != nil, "need AUTH")
	//
	//_, err = c.Do("AUTH", "wrongpasswd")
	//fmt.Println(err != nil, "wrong password")

	c1 := RedisClient.Get()
	//_, err := c1.Do("AUTH", "icoolpy.com")
	//fmt.Println(err)

	_, err := c1.Do("PING")
	fmt.Println(err)

	r, err := redis.Int(c1.Do("DEL", "incrs", "aap"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)

	c2 := RedisClient.Get()
	_, err = c2.Do("SET", "foo", "bar")
	if err != nil {
		t.Error(err)
	}

	_, err = c2.Do("SET", "joo", "bar")
	if v, err := redis.Strings(c2.Do("KEYS", "j*")); err != nil || v[0] != "joo" {
		t.Error("Keys not fire *")
	}

	if v, err := redis.String(c2.Do("GET", "foo")); err == nil {
		fmt.Println(v)
	}

	if v, err := redis.Strings(c2.Do("KEYSSTART", "jo")); err == nil {
		fmt.Println("KEYSSTART")
		for _, val := range v {
			fmt.Println(val)
		}
	}

	fmt.Println("datetime range test")
	tm, err := time.Parse(time.RFC3339Nano, "2013-06-05T14:10:43.678Z")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		key := tm.Add(time.Second * time.Duration(i))
		nkey := key.Format(time.RFC3339Nano)
		var nb []byte
		for _, r := range "1,2," {
			nb = append(nb, byte(r))
		}
		for _, r := range nkey {
			nb = append(nb, byte(r))
		}
		_, err = c2.Do("SET", string(nb), "")
	}
	v, err := redis.Strings(c2.Do("KEYSRANGE", "1,2,2013-06-05T14:10:41", "1,2,2013-06-05T14:11:46"))
	fmt.Println("KEYSRANGE")
	fmt.Println(err)
	for _, val := range v {
		fmt.Println(val)
	}

	if _, err = redis.String(c2.Do("SELECT", "5")); err != nil {
		t.Error(err)
	}

	if _, err = redis.String(c2.Do("SET", "foo", "baz")); err != nil {
		t.Error(err)
	}

	fmt.Println("select 15")
	fmt.Println("pop start")
	c3 := RedisClient.Get()
	if _, err = redis.String(c3.Do("SELECT", "15")); err != nil {
		t.Error(err)
	}
	myv, err := redis.Int(c3.Do("RPUSH", "foo2","bar3"));
	if  err !=nil{
		t.Error(err)
	}
	fmt.Println("push finish count")
	fmt.Println(myv)
	c4 := RedisClient.Get()
	if _, err = redis.String(c4.Do("SELECT", "15")); err != nil {
		t.Error(err)
	}
	for i := 0; i < 3; i++ {
		av, err := redis.String(c4.Do("LPOP"));
		if  err !=nil{
			fmt.Println("pop end of queu be show err msg")
			fmt.Println(err)
			break;
		}
		fmt.Println("pop finish", i)
		fmt.Println(av)
	}
}
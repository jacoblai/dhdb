package Redico

import (
	"github.com/garyburd/redigo/redis"
	"testing"
	"fmt"
	"time"
)

func TestRedico(t *testing.T) {
	// Configure you application to connect to redis at s.Addr()
	// Any redis client should work, as long as you use redis commands which
	c, err := redis.Dial("tcp", "127.0.0.1:6380")
	if err != nil {
		t.Error(err)
	}

	_, err = c.Do("AUTH", "foo", "bar")
	fmt.Println(err != nil, "no password set")

	_, err = c.Do("PING", "foo", "bar")
	fmt.Println(err != nil, "need AUTH")

	_, err = c.Do("AUTH", "wrongpasswd")
	fmt.Println(err != nil, "wrong password")

	_, err = c.Do("AUTH", "icoolpy.com")
	fmt.Println(err)

	_, err = c.Do("PING")
	fmt.Println(err)

	r, err := redis.Int(c.Do("DEL", "incrs", "aap"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)

	_, err = c.Do("SET", "foo", "bar")
	if err != nil {
		t.Error(err)
	}

	_, err = c.Do("SET", "joo", "bar")
	if v, err := redis.Strings(c.Do("KEYS", "j*")); err != nil || v[0] != "joo" {
		t.Error("Keys not fire *")
	}

	if v, err := redis.String(c.Do("GET", "foo")); err == nil {
		fmt.Println(v)
	}

	if v, err := redis.Strings(c.Do("KEYSSTART", "jo")); err == nil {
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
		_, err = c.Do("SET", string(nb), "")
	}
	v, err := redis.Strings(c.Do("KEYSRANGE", "1,2,2013-06-05T14:10:41", "1,2,2013-06-05T14:11:46"))
	fmt.Println("KEYSRANGE")
	fmt.Println(err)
	for _, val := range v {
		fmt.Println(val)
	}

	if _, err = redis.String(c.Do("SELECT", "5")); err != nil {
		t.Error(err)
	}

	if _, err = redis.String(c.Do("SET", "foo", "baz")); err != nil {
		t.Error(err)
	}

	if _, err = redis.String(c.Do("SELECT", "15")); err != nil {
		t.Error(err)
	}
	fmt.Println("select 15")
	myv, err := redis.Int(c.Do("RPUSH", "foo2"));
	if  err !=nil{
		t.Error(err)
	}
	fmt.Println("push finish count")
	fmt.Println(myv)
	//fmt.Println("pop finish count")
	//av, err := redis.String(c.Do("LPOP"));
	//if  err !=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("pop finish")
	//fmt.Println(av)

}
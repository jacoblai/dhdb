# DHDB - A fast NoSQL database for storing big list of data

[![Author](https://img.shields.io/badge/author-@jacoblai-blue.svg?style=flat)](http://www.icoolpy.com/) [![Platform](https://img.shields.io/badge/platform-Linux,%20OpenWrt,%20Mac,%20Windows-green.svg?style=flat)](https://github.com/jacoblai/dhdb) [![NoSQL](https://img.shields.io/badge/db-NoSQL-pink.svg?tyle=flat)](https://github.com/jacoblai/dhdb)


DHDB is a high performace key-value(key-string, List-keys) NoSQL database, __an alternative to Redis__.

## Features

* Pure Go 
* Big data list to 10 billion
* Redis all clients are supported
* Persistent queue service
* Android or OpenWrt os supported (ARM/MIPS)

## redis clients supported 

[All redis clients supported](https://redis.io/clients)

## Windows executable

[Download dhdb from here (Windows,Mac,Linux,Android,OpenWrt)](https://github.com/jacoblai/dhdb/tree/master/dhdb-bin)

## Executable flags

 - flags
   - p -- dhdb server host port number (default 6380)
   - a -- dhdb client auth password (default icoolpy.com)

## DHDB supported redis implemented commands

 - Connection (complete)
   - AUTH -- see RequireAuth()
   - ECHO
   - PING
   - SELECT
   - QUIT
 - Key 
   - DEL
   - EXISTS
   - KEYS
   - SCAN
 - String keys (complete)
   - APPEND
   - GET
   - INCR
   - SET
   - RENAME
 - List keys (complete)
   - LPOP
   - RPUSH

## DHDB custom commands

 - Key 
   - KEYSSTART
   - KEYSRANGE
   
## DHDB redigo client sample

## connection
```
import (
 "github.com/garyburd/redigo/redis"
 "fmt"
)

var RedisClient     *redis.Pool
func main() {
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
}
```
## SET
```
c := RedisClient.Get()
_, err := c.Do("SET", "foo", "bar","joo", "bar")
```
## GET 
```
if v, err := redis.String(c.Do("GET", "foo")); err == nil {
 fmt.Println(v)
}
```
## KEYS
```
// find all keys start with 'j' word 
v, err := redis.Strings(c.Do("KEYS", "j*"));
if  err != nil || v[0] != "joo" {
 fmt.Println("Keys not fire *")
}
```
## KEYSSTART (sreach keys from goleveldb Iterator not suport regexp) 
```
v, err := redis.Strings(c.Do("KEYSSTART", "jo"));
if  err == nil {
 fmt.Println("KEYSSTART")
 for _, val := range v {
 	fmt.Println(val)
 }
}
```

## KEYSRANGE (sreach keys from goleveldb Iterator suport range datetime keys express)
```
//gen data 
tm, _ := time.Parse(time.RFC3339Nano, "2017-01-09T14:10:43.678Z")
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
//sreach range keys
v, _ := redis.Strings(c.Do("KEYSRANGE", "1,2,2017-01-09T14:10:41", "1,2,2017-01-09T14:11:46"))
for _, val := range v {
 fmt.Println(val)
}
```

## SELECT (change database List)
```
if _, err = redis.String(c.Do("SELECT", "5")); err != nil {
 fmt.Println(err)
}
```

## RPUSH (Qeueu must init new list by select command)
```
//change new list for queue by select command
_, _ = redis.String(c.Do("SELECT", "15"));
myv, _ := redis.Int(c.Do("RPUSH", "foo2","bar3"));
fmt.Println(myv)//print finish count
```

## LPOP (Qeueu must init list queue type)
```
//change new list for queue by select command
_, _ = redis.String(c.Do("SELECT", "15"));
for i := 0; i < 3; i++ {
 av, err := redis.String(c.Do("LPOP"));
 if  err !=nil{
  fmt.Println(err)// EOF Queue
  break;
 }
 fmt.Println("pop finish", i)
 fmt.Println(av)//print pop value
}
```

## Authors

@jacoblai

## Thanks

* syndtr, github.com/syndtr/goleveldb
* bsm, github.com/bsm/redeo

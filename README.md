# DHDB - A fast NoSQL database for storing big list of data

[![Author](https://img.shields.io/badge/author-@jacoblai-blue.svg?style=flat)](http://www.icoolpy.com/) [![Platform](https://img.shields.io/badge/platform-Linux,%20OpenWrt,%20Mac,%20Windows-green.svg?style=flat)](https://github.com/jacoblai/dhdb) [![NoSQL](https://img.shields.io/badge/db-NoSQL-pink.svg?tyle=flat)](https://github.com/jacoblai/dhdb)


DHDB is a high performace key-value(key-string, List-keys) NoSQL database, __an alternative to Redis__.

## Features

* Pure Go 
* Big data list to 10 billion
* Redis all clients are supported
* Persistent queue service
* Android or OpenWrt os supported (ARM/MIPS)

## redis clients supported [Let's go...](https://redis.io/clients)

## Windows executable

Download from here: https://github.com/jacoblai/dhdb-bin

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

## DHDB customer commands

 - Key 
   - KEYSSTART
   - KEYSRANGE

## Authors

@jacoblai

## Thanks

* syndtr, https://github.com/syndtr/goleveldb
* bsm, github.com/bsm/redeo

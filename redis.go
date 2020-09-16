package redis

import (
	"fmt"
	"strconv"

	"github.com/mediocregopher/radix/v3"
)

var radixClient radix.Client

const (
	// TCP Network
	TCP = "tcp"
	// UDP Network
	UDP = "udp"
)

// RConnect struct doc
type RConnect struct {
}

// DefaultClient radix default client
func DefaultClient(address string) (*RConnect, error) {
	var err error
	radixClient, err = radix.DefaultClientFunc(TCP, address)
	if err != nil {
		return nil, err
	}
	return &RConnect{}, nil
}

// Close connect to redis
func (r RConnect) Close() error {
	return radixClient.Close()
}

// Get redis CLI command
func (r RConnect) Get(key string) (string, error) {
	var value string
	err := radixClient.Do(radix.Cmd(&value, "GET", key))
	if err != nil {
		return value, err
	}
	fmt.Println("Redis: The GET action has been successfully executed")
	return value, nil
}

// Set redis CLI command
func (r RConnect) Set(key, value string) error {
	err := radixClient.Do(radix.Cmd(nil, "SET", key, value))
	if err != nil {
		return err
	}
	fmt.Println("Redis: The SET action has been successfully executed")
	return nil
}

// Del redis CLI command
func (r RConnect) Del(key string) error {
	err := radixClient.Do(radix.Cmd(nil, "DEL", key))
	if err != nil {
		return err
	}
	fmt.Println("Redis: The DEL action has been successfully executed")
	return nil
}

// Expire redis CLI command
func (r RConnect) Expire(key string, ttl int) error {
	var response int
	ttlStr := strconv.Itoa(ttl)
	err := radixClient.Do(radix.Cmd(&response, "EXPIRE", key, ttlStr))
	if err != nil {
		return err
	}
	if response == 0 {
		return fmt.Errorf("The document was not found")
	}
	fmt.Println("Redis: The EXPIRE action has been successfully executed")
	return nil
}

// TTL redis CLI command
func (r RConnect) TTL(key string) (int, error) {
	var ttlStr string
	err := radixClient.Do(radix.Cmd(&ttlStr, "TTL", key))
	if err != nil {
		return 0, err
	}
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		return 0, err
	}
	if ttl == -2 {
		return ttl, fmt.Errorf("The document was not found")
	}
	fmt.Println("Redis: The TTL action has been successfully executed")
	return ttl, nil
}

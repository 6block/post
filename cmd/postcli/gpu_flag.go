package main

import (
	"fmt"
	"strconv"
	"strings"
)

type providerIds []uint64

func (p *providerIds) Set(val string) error {
	strlist := strings.Split(val, ",")
	var int64list providerIds
	for _, item := range strlist {
		val, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			continue
		}
		int64list = append(int64list, val)
	}
	*p = int64list
	return nil
}

func (p *providerIds) String() string {
	var temp = make([]string, len(*p))
	for k, v := range *p {
		temp[k] = fmt.Sprintf("%d", v)
	}
	return strings.Join(temp, ",")
}

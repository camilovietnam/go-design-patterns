package main

import "math/rand"

type singleton struct {
	uniqueID int
}

var instance *singleton

func GetClient() *singleton {
	var cli *singleton

	if instance != nil {
		cli = instance
	} else {
		instance = new(singleton)
		instance.uniqueID = 123456 + rand.Intn(100000)
		cli = instance
	}

	return cli
}

package helpers

import "sync"

var Wg = &sync.WaitGroup{}


var TokenChannel = make(chan string)
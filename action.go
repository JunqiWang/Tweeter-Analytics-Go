package main

import (
    "github.com/hoisie/web"
	"time"
)

func index(val string) string {
	return "<p>Team Dynamos</p>" + 
		   "<p>Twitter Analytics</p>"
}

func hello(val string) string {
    return "Dynamos 0000-0000-0000</br>" + 
			time.Now().Format("2006-01-02 15:04:05") +
		   "</br>"
} 

func main() {
	web.Get("/()", index)
    web.Get("/(q1)", hello)
    web.Run("0.0.0.0:9999")
}

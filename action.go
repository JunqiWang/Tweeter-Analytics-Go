package main

import (
    "github.com/hoisie/web"
    "time"
)

func index(val string) string {
    return "Team Dynamos\nTwitter Analytics"
}

func hello(val string) string {
    return "Dynamos 0000-0000-0000\n" + 
	    time.Now().Format("2006-01-02 15:04:05") + "\n"
} 

func main() {
    web.Get("/()", index)
    web.Get("/(q1)", hello)
    web.Run("0.0.0.0:9999")
}

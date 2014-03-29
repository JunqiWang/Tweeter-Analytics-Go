package main

import (
    "github.com/hoisie/web"
    "time"
)

func q1(val string) string {
    return "Dynamos 0000-0000-0000\n" + 
	    time.Now().Format("2006-01-02 15:04:05")
} 

func main() {
    web.Get("/(q1)", q1)
    web.Run("0.0.0.0:9999")
}

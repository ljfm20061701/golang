package main

import (
    "fmt"
    "pkg/barrage/douyu"
    "pkg/rrframework/logs"
)

func chatmsg(msg *douyu.Message) {
    level := msg.GetStringField("level")
    nn := msg.GetStringField("nn")
    txt := msg.GetStringField("txt")
    logs.Info(fmt.Sprintf("level(%s) - %s >>> %s", level, nn, txt))
}

func main() {
    client, err := douyu.Connect("openbarrage.douyutv.com:8601", nil)
    if err != nil {
        logs.Error(err)
        return
    }

    client.HandlerRegister.Add("chatmsg", douyu.Handler(chatmsg), "chatmsg")
    if err := client.JoinRoom(1691469); err != nil {
        logs.Error(fmt.Sprintf("Join room fail, %s", err.Error()))
        return
    }
    client.Serve()
}

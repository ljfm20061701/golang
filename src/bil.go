package main

import (
    "pkg/barrage/bilibili"
    "pkg/rrframework/logs"
)

func danmu(msg *bilibili.Message) {
    logs.Debug(">>> ", string(msg.Bytes()))
}

func main() {
    // uri, userid, handlerRegister
    client, err := bilibili.Connect("https://live.bilibili.com/239", -1, nil)
    if err != nil {
        logs.Error(err)
        return
    }
    client.HandlerRegister.Add(bilibili.DANMU_MSG, bilibili.Handler(danmu), "danmu")
    client.Serve()
}

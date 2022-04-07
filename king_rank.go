package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type KingRankResp struct {
	Code int64 `json:"code"`
	Data struct {
		Replies []struct {
			Content struct {
				Device  string        `json:"device"`
				JumpURL struct{}      `json:"jump_url"`
				MaxLine int64         `json:"max_line"`
				Members []interface{} `json:"members"`
				Message string        `json:"message"`
				Plat    int64         `json:"plat"`
			} `json:"content"`
			Count  int64 `json:"count"`
			Folder struct {
				HasFolded bool   `json:"has_folded"`
				IsFolded  bool   `json:"is_folded"`
				Rule      string `json:"rule"`
			} `json:"folder"`
			Like    int64 `json:"like"`
			Replies []struct {
				Action  int64 `json:"action"`
				Assist  int64 `json:"assist"`
				Attr    int64 `json:"attr"`
				Content struct {
					Device  string   `json:"device"`
					JumpURL struct{} `json:"jump_url"`
					MaxLine int64    `json:"max_line"`
					Message string   `json:"message"`
					Plat    int64    `json:"plat"`
				} `json:"content"`
				Rcount  int64       `json:"rcount"`
				Replies interface{} `json:"replies"`
			} `json:"replies"`
			Type int64 `json:"type"`
		} `json:"replies"`
	} `json:"data"`
	Message string `json:"message"`
}

func StartKingRank() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/v2/reply/main?next=0&type=1&oid=595029714&mode=3&plat=1", nil)
	if err != nil {
		fmt.Println("req err", err)
	}
	req.Header.Set("authority", "api.bilibili.com")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
	req.Header.Set("accept", "*/*")
	req.Header.Set("referer", "https://www.bilibili.com/bangumi/play/ep424605?from=search&seid=12185563008772548657&spm_id_from=333.337.0.0")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io err", err)
	}

	// JSON 反序列化
	var resultList KingRankResp
	_ = json.Unmarshal(bodyText, &resultList)
	for _, result := range resultList.Data.Replies {
		fmt.Println("一级评论：", result.Content.Message)
		for _, reply := range result.Replies {
			fmt.Println("    二级评论：", reply.Content.Message)
		}
	}
}

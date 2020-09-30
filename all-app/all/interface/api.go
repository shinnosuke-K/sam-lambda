package _interface

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func getTicketsBody() ([]byte, error) {
	types := []string{"problem", "incident", "question", "task"}
	priority := []string{"urgent", "high", "normal", "low"}
	status := []string{"new", "open", "pending", "hold", "solved", "closed"}
	subject := []string{
		"【完全保存版】問い合わせの見積もりとスケジュール管理",
		"問い合わせが驚くほど良くわかる、まとめ記事のまとめ",
		"知らないヒトは絶対損してる！ 4円で遊べる『問い合わせガイド』",
		"質問はエンジニアを幸せにはしなかった",
		"質問割ろうぜ！ 2日6分で質問が手に入る「1分間質問運動」の動画が話題に",
		"もう質問って絶対必要なんだから、3日10円出して「質問」で勉強しときなよ！",
		"コピペで使えるオシャレなATOM見本 81 （全組み合わせ付） ",
		"自分は探すものではなく、つくるもの。それが「ATOM」",
		"ATOM次期システム関連のまとめ",
	}

	test := TicketsResponse{}

	for n := 0; n < 10000; n++ {
		test.Tickets = append(test.Tickets, Ticket{
			ID:             int64(n + 1),
			CreateTime:     time.Now().Add(time.Duration(n)),
			UpdateTime:     time.Now().Add(time.Duration(n)),
			Type:           types[n%4],
			Subject:        subject[n%len(subject)],
			Priority:       priority[n%len(priority)],
			Status:         status[n%len(status)],
			Tags:           []string{"a", "b", "c"},
			RequesterID:    int64(n % 3),
			AssigneeID:     int64(n%3 + 1),
			OrganizationID: int64(n%3 + 2),
		})
	}
	test.NextPage = "localhost"

	j, err := json.Marshal(&test)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func getUsersBody() ([]byte, error) {
	test := UsersResponse{}
	for n := 0; n < 3; n++ {
		test.Users = append(test.Users, User{
			ID:             int64(n + 1),
			Name:           fmt.Sprintf("%d", n+1),
			Email:          "",
			CreateTime:     time.Now().Local().Add(time.Duration(n)),
			OrganizationID: int64(n + 2),
			Alias:          "",
			Role:           "",
		})
	}
	j, err := json.Marshal(test)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func getOrgsBody() ([]byte, error) {
	test := OrgsResponse{}
	for n := 0; n < 3; n++ {
		test.Orgs = append(test.Orgs, Organization{
			ID:         int64(n + 2),
			Name:       fmt.Sprintf("株式会社%d", n),
			CreateTime: time.Now().Add(time.Duration(n)),
		})
	}
	j, err := json.Marshal(test)
	if err != nil {
		return nil, err
	}
	return j, nil

}

func GetBody(format interface{}) ([]byte, error) {
	switch format.(type) {
	case *TicketsResponse:
		fmt.Println("ticket response")
		return getTicketsBody()

	case *UsersResponse:
		return getUsersBody()

	case *OrgsResponse:
		return getOrgsBody()
	}
	return nil, errors.New("invalid format")
}

func Parse(format interface{}, byteBody []byte) (interface{}, error) {

	switch format.(type) {
	case *TicketsResponse:
		var t TicketsResponse
		fmt.Println("ticket")
		err := json.Unmarshal(byteBody, &t)
		if err != nil {
			return nil, err
		}
		return &t, nil

	case *UsersResponse:
		var u UsersResponse
		err := json.Unmarshal(byteBody, &u)
		if err != nil {
			return nil, err
		}
		return &u, nil

	case *OrgsResponse:
		var o OrgsResponse
		err := json.Unmarshal(byteBody, &o)
		if err != nil {
			return nil, err
		}
		return &o, nil
	}

	return nil, errors.New("invalid format")
}

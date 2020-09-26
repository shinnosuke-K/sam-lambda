package table

import (
	"encoding/json"
	"strings"
	"time"

	"gorm.io/gorm"
)

type JsonTicket struct {
	Contents []Content `json:"ticket"`
	NextPage string    `json:"next_page"`
}

type Content struct {
	ID             int64     `json:"id"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
	Type           string    `json:"type"`
	Subject        string    `json:"subject"`
	Priority       string    `json:"priority"`
	Status         string    `json:"status"`
	Tags           []string  `json:"tags"`
	RequesterID    int64     `json:"requester_id"`
	AssigneeID     int64     `json:"assignee_id"`
	OrganizationID int64     `json:"organization_id"`
}

type Ticket struct {
	ID             int64
	CreateTime     time.Time
	UpdateTime     time.Time
	Type           string
	Subject        string
	Priority       string
	Status         string
	Tag            string
	RequesterID    int64
	AssigneeID     int64
	OrganizationID int64
}

func (t *JsonTicket) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&Ticket{})
}

func (t *JsonTicket) CreateTable(db *gorm.DB) error {
	return db.Migrator().CreateTable(&Ticket{})
}

func (t *JsonTicket) GetBody() []byte {

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

	test := JsonTicket{}

	for n := 0; n < 10000; n++ {
		test.Contents = append(test.Contents, Content{
			ID:             int64(n),
			CreateTime:     time.Now().Add(time.Duration(n)),
			UpdateTime:     time.Now().Add(time.Duration(n)),
			Type:           types[n%4],
			Subject:        subject[n%len(subject)],
			Priority:       priority[n%len(priority)],
			Status:         status[n%len(status)],
			Tags:           []string{},
			RequesterID:    int64(n % 3),
			AssigneeID:     int64(n%3 + 1),
			OrganizationID: int64(n%3 + 2),
		})
	}

	j, _ := json.Marshal(test)
	return j
}

func (t *JsonTicket) Mapping(jsonBody []byte) error {
	err := json.Unmarshal(jsonBody, &t)
	if err != nil {
		return err
	}
	return nil
}

func (t *JsonTicket) Insert(db *gorm.DB) {

	for n := range t.Contents {
		ticket := Ticket{
			ID:             t.Contents[n].ID,
			CreateTime:     t.Contents[n].CreateTime,
			UpdateTime:     t.Contents[n].UpdateTime,
			Type:           t.Contents[n].Type,
			Subject:        t.Contents[n].Subject,
			Priority:       t.Contents[n].Priority,
			Status:         t.Contents[n].Status,
			Tag:            strings.Join(t.Contents[n].Tags, ","),
			RequesterID:    t.Contents[n].RequesterID,
			AssigneeID:     t.Contents[n].AssigneeID,
			OrganizationID: t.Contents[n].OrganizationID,
		}
		db.Create(&ticket)
	}
}

func NewTicket() *JsonTicket {
	return new(JsonTicket)
}

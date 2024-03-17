package main

import (
	"fmt"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})
	for _, cmd := range cmds {
		out := make(chan interface{})
		wg.Add(1)
		go func(wg *sync.WaitGroup, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			cmd(in, out)
		}(wg, in, out)
		in = out
	}
	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	// 	in - string
	// 	out - User

	wg := &sync.WaitGroup{}
	users := &sync.Map{}
	for email := range in {
		wg.Add(1)
		go func(wg *sync.WaitGroup, users *sync.Map, email string) {
			defer wg.Done()
			user := GetUser(email)
			if _, exists := users.Load(user.ID); !exists {
				users.Store(user.ID, true)
				out <- user
			}
		}(wg, users, email.(string))
	}
	wg.Wait()
}

// Batch of users to send
type batch struct {
	users []User
}

func (b *batch) Add(user User) {
	if !b.IsFull() {
		b.users = append(b.users, user)
	}
}

func (b batch) IsFull() bool {
	return len(b.users) == 2
}

func (b batch) IsEmpty() bool {
	return len(b.users) == 0
}

func (b *batch) Clear() []User {
	users := b.users
	b.users = nil
	return users
}

func SelectMessages(in, out chan interface{}) {
	// 	in - User
	// 	out - MsgID

	wg := &sync.WaitGroup{}
	var batchUsers batch
	for user := range in {
		batchUsers.Add(user.(User))
		if batchUsers.IsFull() {
			wg.Add(1)
			go func(wg *sync.WaitGroup, users []User) {
				defer wg.Done()
				msgIds, _ := GetMessages(users...)
				for _, msgId := range msgIds {
					out <- msgId
				}
			}(wg, batchUsers.Clear())
		}
	}

	if !batchUsers.IsEmpty() {
		wg.Add(1)
		go func(wg *sync.WaitGroup, users []User) {
			defer wg.Done()
			msgIds, _ := GetMessages(users...)
			for _, msgId := range msgIds {
				out <- msgId
			}
		}(wg, batchUsers.Clear())
	}

	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData

	wg := &sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(wg *sync.WaitGroup, in, out chan interface{}) {
			defer wg.Done()
			for msgId := range in {
				msgId := msgId.(MsgID)
				hasSpam, err := HasSpam(msgId)
				if err != nil {
					fmt.Println(err)
				}
				msgData := MsgData{ID: msgId, HasSpam: hasSpam}
				out <- msgData
			}
		}(wg, in, out)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string
	var data []MsgData
	for msgData := range in {
		data = append(data, msgData.(MsgData))
	}

	sort.Slice(data, func(i int, j int) bool {
		if data[i].HasSpam == data[j].HasSpam {
			return data[i].ID < data[j].ID
		}

		return data[i].HasSpam
	})

	for _, msgData := range data {
		out <- fmt.Sprintf("%v %v", msgData.HasSpam, msgData.ID)
	}
}

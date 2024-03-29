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
		emailStr, ok := email.(string)
		if ok {
			wg.Add(1)
			go func(wg *sync.WaitGroup, users *sync.Map, emailStr string) {
				defer wg.Done()
				user := GetUser(emailStr)
				if _, exists := users.Load(user.ID); !exists {
					users.Store(user.ID, true)
					out <- user
				}
			}(wg, users, emailStr)
		}
	}
	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	// 	in - User
	// 	out - MsgID
	wg := &sync.WaitGroup{}
	batchUsers := []User{}
	for user := range in {
		userStruct, ok := user.(User)
		if ok {
			batchUsers = append(batchUsers, userStruct)
		}

		if len(batchUsers) == GetMessagesMaxUsersBatch {
			wg.Add(1)
			go func(wg *sync.WaitGroup, users []User) {
				defer wg.Done()
				msgIds, err := GetMessages(users...)
				if err == nil {
					for _, msgId := range msgIds {
						out <- msgId
					}
				}
			}(wg, batchUsers)
			batchUsers = nil
		}
	}

	if len(batchUsers) > 0 {
		wg.Add(1)
		go func(wg *sync.WaitGroup, users []User) {
			defer wg.Done()
			msgIds, err := GetMessages(users...)
			if err == nil {
				for _, msgId := range msgIds {
					out <- msgId
				}
			}
		}(wg, batchUsers)
	}

	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData

	wg := &sync.WaitGroup{}
	wg.Add(HasSpamMaxAsyncRequests)
	for i := 0; i < HasSpamMaxAsyncRequests; i++ {
		go func(wg *sync.WaitGroup, in, out chan interface{}) {
			defer wg.Done()
			for msgId := range in {
				msgIdStruct, ok := msgId.(MsgID)
				if ok {
					hasSpam, err := HasSpam(msgIdStruct)
					if err == nil {
						msgData := MsgData{ID: msgIdStruct, HasSpam: hasSpam}
						out <- msgData
					}
				}
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
		msgDataStruct, ok := msgData.(MsgData)
		if ok {
			data = append(data, msgDataStruct)
		}
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

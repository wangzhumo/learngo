package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestLotery(t *testing.T) {
	testClient := httptest.New(t, newApp())

	// 开启一个Group
	var wg sync.WaitGroup
	testClient.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前人数：0\n")

	//100 个协程来处理 - Add User
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// 抽奖人员
			testClient.POST("/import").
				WithFormField("users", fmt.Sprintf("test_u_%d", index)).
				Expect().Status(httptest.StatusOK)

		}(i)

	}
	wg.Wait()

	testClient.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前人数：100\n")

	testClient.GET("/lottery").Expect().Status(httptest.StatusOK).
		Body().Equal("当前人数：99\n")
}

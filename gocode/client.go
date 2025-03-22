package main

import (
	"context"
	"fmt"
	"go_api/api"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, World!")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	exec(conn)
}

func exec(conn *grpc.ClientConn) {
	ctx := context.Background()

	client := api.NewPersonServiceClient(conn)
	wg := &sync.WaitGroup{}
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			p := &api.Person{
				Id: 2,
			}
			res, err := client.GetPerson(ctx, p)
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Time: ", time.Since(start))
}

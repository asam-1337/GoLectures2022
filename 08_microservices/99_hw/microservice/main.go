package main

import (
	context "context"
	"fmt"
	"log"
	"time"
)

func a(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("dmkvsklvklmdsm")
			return
		}
	}
}

const (
	// какой адрес-порт слушать серверу
	listenAddr2 string = "127.0.0.1:8082"

	// кого по каким методам пускать
	ACLData2 string = `{
	"logger":    ["/main.Admin/Logging"],
	"stat":      ["/main.Admin/Statistics"],
	"biz_user":  ["/main.Biz/Check", "/main.Biz/Add"],
	"biz_admin": ["/main.Biz/*"]
}`
)

// чтобы не было сюрпризов когда где-то не успела преключиться горутина и не успело что-то стортовать
func wait2(amout int) {
	time.Sleep(time.Duration(amout) * 10 * time.Millisecond)
}

func main() {
	ctx2, finish2 := context.WithCancel(context.Background())

	go a(ctx2)
	wait2(2)
	finish2()

	ctx, finish := context.WithCancel(context.Background())
	err := StartMyMicroservice2(ctx, listenAddr2, ACLData2)
	if err != nil {
		log.Fatalf("cant start server initial: %v", err)
	}
	wait2(2)
	fmt.Println("dsvsvd")
	finish() // при вызове этой функции ваш сервер должен остановиться и освободить порт
	wait2(2)

	// теперь проверим что вы освободили порт и мы можем стартовать сервер ещё раз
	ctx, finish = context.WithCancel(context.Background())
	go StartMyMicroservice(ctx, listenAddr2, ACLData2)
	if err != nil {
		log.Fatalf("cant start server again: %v", err)
	}
	wait2(1)
	finish()
	wait2(1)
}

func StartMyMicroservice2(ctx context.Context, addr, ACLData string) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("aaaaaaaaaaaaaaa")
			return nil
		}
	}

	// lis, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	return fmt.Errorf("cant listen on port : %w", err)
	// }
	// server := grpc.NewServer()

	// RegisterBizServer(server, NewBizServer())
	// RegisterAdminServer(server, NewAdminServer())

	// fmt.Printf("starting server at %s\n", addr)
	// errCh := make(chan error)
	// go wraper(errCh, lis, server)
	// fmt.Println("aakjj")
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println("mlmlml")
	// 		server.Stop()
	// 		fmt.Println("blblbl")
	// 		err = <-errCh
	// 		fmt.Println(err)
	// 		return nil
	// 	case <-errCh:
	// 		fmt.Println("lsakfksaf")
	// 		return fmt.Errorf("blabla")
	// 	}
	// }

}

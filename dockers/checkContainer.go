package dockers

/*
检测容器运行
**/
// func isRuning(ctx context.Context, cli *client.Client) <-chan bool {
// 	isRun := make(chan bool)
// 	var timer *time.Ticker
// 	go func(ctx context.Context, cli *client.Client) {
// 		for {
// 			//每n s检查一次容器是否运行
// 			timer = time.NewTicker(time.Duration(n) * time.Second)
// 			select {
// 			case <-timer.C:
// 				//获取正在运行的container list
// 				log.Printf("%s is checking the container[%s]is Runing??", os.Args[0], containerName)
// 				contTemp := getContainer(ctx, cli, false)
// 				if contTemp.ID == "" {
// 					log.Print(":NO")
// 					//说明container没有运行
// 					isRun <- true
// 				} else {
// 					log.Print(":YES")
// 					//说明该container正在运行
// 					go printConsole(ctx, cli, contTemp.ID)
// 				}
// 			}

// 		}
// 	}(ctx, cli)
// 	return isRun
// }

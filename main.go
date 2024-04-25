//package main
//
////import (
////	_ "beegoWeb/routers"
////	beego "github.com/beego/beego/v2/server/web"
////	"github.com/beego/beego/v2/server/web/filter/cors"
////)
//
////func main() {
////	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
////		AllowAllOrigins:  true,
////		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
////		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
////		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
////		AllowCredentials: true,
////	}))
////
////	beego.Run()
////}
//
//import (
//	"fmt"
//	"log"
//	"os"
//	"path/filepath"
//	"time"
//
//	"github.com/xuri/excelize/v2"
//)
//
//// listExcelSheets 打开Excel文件并列出所有Sheet的名称
//func listExcelSheets(filePath string) {
//	f, err := excelize.OpenFile(filePath)
//	if err != nil {
//		log.Printf("无法打开文件 '%s': %v\n", filePath, err)
//		return
//	}
//	defer f.Close()
//
//	sheets := f.GetSheetList()
//	fmt.Printf("文件: %s\n", filePath)
//	fmt.Println("包含的Sheets:")
//	for _, sheet := range sheets {
//		fmt.Printf("- %s\n", sheet)
//	}
//	fmt.Println()
//}
//
//// walkExcelFiles 遍历指定目录及子目录，寻找Excel文件
//func walkExcelFiles(rootPath string) {
//	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			log.Printf("访问文件失败 '%s': %v\n", path, err)
//			return err
//		}
//		if !info.IsDir() && (filepath.Ext(path) == ".xlsx" || filepath.Ext(path) == ".xls") {
//			listExcelSheets(path)
//		}
//		return nil
//	})
//
//	if err != nil {
//		log.Printf("遍历目录失败 '%s': %v\n", rootPath, err)
//	}
//}
//
//func main() {
//	// 需要遍历的目录路径，根据需要进行修改
//	rootPath := "D:\\elex\\slgconfiguration\\excel"
//
//	walkExcelFiles(rootPath)
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// 使用time.AfterFunc
//	timer := time.AfterFunc(2*time.Second, func() {
//		fmt.Println("Timer expired")
//	})
//
//	// 为了看到定时器到期的效果，这里等待一段时间
//	time.Sleep(3 * time.Second)
//
//	// 停止定时器（如果它还没过期的话）
//	timer.Stop()
//}

//package main
//
//import (
//	"fmt"
//	"testing"
//	"time"
//)
//
//func main() {
//	ticker := time.NewTicker(1 * time.Second)
//	defer ticker.Stop() // 程序结束时停止定时器
//
//	for i := 0; i < 5; i++ {
//		<-ticker.C // 等待定时器的下一个滴答
//		fmt.Println("Tick")
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// 创建一个打点器，每5秒触发一次
//	ticker := time.NewTicker(5 * time.Second)
//	defer ticker.Stop() // 当不再需要打点器时确保能够停止它
//
//	for {
//		select {
//		case <-ticker.C:
//			// 每5秒会执行到这里
//			fmt.Println("Hello World")
//		}
//	}
//}

package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	// 每5秒执行一次

	_, err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Hello World")
	})
	if err != nil {
		fmt.Println("Error scheduling a job:", err)
		return
	}

	// 开始执行定时器
	c.Start()

	// 防止主程序立即结束，使用select{}无限等待
	select {}
}

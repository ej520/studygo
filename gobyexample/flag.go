package main
import (
    "fmt"
    "flag"
)

func main(){
	//http://docscn.studygolang.com/pkg/
	//func String(name string, value string, usage string) *string
	//String定义了一个有指定名字，默认值，和用法说明的string标签。 返回值是一个存储标签解析值的string变量地址
    data_path := flag.String("D","/home/manu/sample/","DB data path") //指定名字为D  默认值为/home/manu/sample/ ，说明DB data path
    log_file := flag.String("l","/home/manu/sample.log","log file")
    nowait_flag :=flag.Bool("W",false,"do not wait until operation completes")

    flag.Parse()

	var cmd  = flag.Arg(0);
	var cmd2  = flag.Arg(1);

    fmt.Printf("action   : %s %s\n",cmd,cmd2)
    fmt.Printf("data path: %s\n",*data_path)
    fmt.Printf("log file : %s\n",*log_file)
    fmt.Printf("nowait     : %v\n",*nowait_flag)

    fmt.Printf("-------------------------------------------------------\n")

    fmt.Printf("there are %d non-flag input param\n",flag.NArg())
    for i,param := range flag.Args(){
        fmt.Printf("#%d    :%s\n",i,param)
    }


}
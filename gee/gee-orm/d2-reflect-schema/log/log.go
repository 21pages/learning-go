package log

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	errlog  = log.New(os.Stdout, "\033[31m[err]\033[0m ", log.LstdFlags|log.Lshortfile) //red
	warnlog = log.New(os.Stdout, "\033[33m[wan]\033[0m ", log.LstdFlags|log.Lshortfile) //yellow
	infolog = log.New(os.Stdout, "\033[34m[inf]\033[0m ", log.LstdFlags|log.Lshortfile) //blue
	dbglog  = log.New(os.Stdout, "\033[36m[dbg]\033[0m ", log.LstdFlags|log.Lshortfile) //gree
	mu      sync.Mutex
)

//export
var (
	Error  = errlog.Println
	Errorf = errlog.Printf
	Warn   = warnlog.Println
	Warnf  = warnlog.Printf
	Info   = infolog.Println
	Infof  = infolog.Printf
	Dbg    = dbglog.Println
	Dbgf   = dbglog.Printf
)

type Level byte

const (
	LevelDbg Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelDisabled
)

func SetLevel(level Level) {
	mu.Lock()
	defer mu.Unlock()

	loggers := []*log.Logger{errlog, warnlog, infolog, dbglog}
	levels := []Level{LevelError, LevelWarn, LevelInfo, LevelDbg}

	for i := 0; i < len(loggers); i++ {
		if levels[i] < level {
			loggers[i].SetOutput(io.Discard)
		} else {
			loggers[i].SetOutput(os.Stdout)
		}
	}
}

/*
字体颜色

none         = "\033[0m"
black        = "\033[0;30m"
dark_gray    = "\033[1;30m"
blue         = "\033[0;34m"
light_blue   = "\033[1;34m"
green        = "\033[0;32m"
light_green -= "\033[1;32m"
cyan         = "\033[0;36m"
light_cyan   = "\033[1;36m"
red          = "\033[0;31m"
light_red    = "\033[1;31m"
purple       = "\033[0;35m"
light_purple = "\033[1;35m"
brown        = "\033[0;33m"
yellow       = "\033[1;33m"
light_gray   = "\033[0;37m"
white        = "\033[1;37m"

字背景颜色范围: 40--49                   字颜色: 30--39
        40: 黑                          30: 黑
        41:红                          31: 红
        42:绿                          32: 绿
        43:黄                          33: 黄
        44:蓝                          34: 蓝
        45:紫                          35: 紫
        46:深绿                        36: 深绿
        47:白色                        37: 白色

输出特效格式控制：

\033[0m  关闭所有属性
\033[1m   设置高亮度
\03[4m   下划线
\033[5m   闪烁
\033[7m   反显
\033[8m   消隐
\033[30m   --   \033[37m   设置前景色
\033[40m   --   \033[47m   设置背景色

光标位置等的格式控制：

\033[nA  光标上移n行
\03[nB   光标下移n行
\033[nC   光标右移n行
\033[nD   光标左移n行
\033[y;xH设置光标位置
\033[2J   清屏
\033[K   清除从光标到行尾的内容
\033[s   保存光标位置
\033[u   恢复光标位置
\033[?25l   隐藏光标
\33[?25h   显示光标
*/

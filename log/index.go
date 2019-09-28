package log

import (
	"os"
	"path"
	"time"

	"github.com/olivere/elastic"
	"gopkg.in/sohlich/elogrus.v7"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// GlobalLog 全局日志
var GlobalLog = logrus.New()

var logName = "account_log"

func newLfsHook(logLevel log.Level, maxRemainCnt uint) log.Hook {
	writer, err := rotatelogs.New(
		logName+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	log.SetLevel(logLevel)

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})

	return lfsHook
}

// Init 全局日志
func Init() {
	GlobalLog.Formatter = &log.JSONFormatter{}
	GlobalLog.Out = os.Stdout
	GlobalLog.Level = log.InfoLevel

	logPath := "/tmp/iotLog"
	logFileName := "account_log"
	maxAge := time.Hour * 24
	rotationTime := time.Hour
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		GlobalLog.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})
	GlobalLog.AddHook(lfHook)

	// 发送日志到es
	client, _ := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	hook, _ := elogrus.NewElasticHook(client, "localhost", log.InfoLevel, "account_log")
	GlobalLog.AddHook(hook)
}

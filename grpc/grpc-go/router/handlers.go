package router

import (
	"context"
	"fmt"
	"net/http"
	"tutorial/models"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/govalidator"
	"google.golang.org/grpc"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GO HTTP: Hello, World!"))
}

func Compress(w http.ResponseWriter, r *http.Request) {
	LogErr := func(msgErr string, httpStatus int, err interface{}) {
		log.WithFields(log.Fields{
			"status_code": httpStatus,
			"cause":       err,
			"handler":     "ExportPayment",
			"error":       msgErr,
		}).Error(msgErr)
	}
	var compressRules = govalidator.MapData{
		"message": []string{"required"},
	}
	var opts models.CompressBody
	validatorOpts := govalidator.Options{
		Request: r,
		Rules:   compressRules,
		Data:    &opts,
	}
	v := govalidator.New(validatorOpts)
	errs := v.ValidateJSON()
	if len(errs) > 0 {
		LogErr("Compress: Status Bad Request Error at Validating", http.StatusBadRequest, errs)
		w.Write([]byte("Redis error"))
		return
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	flag, err := rdb.Exists(opts.Message).Result()
	if err != nil {
		LogErr("Compress: Redis error", http.StatusBadRequest, errs)
		w.Write([]byte("Redis error"))
		return
	}
	if flag == 1 {
		x, err := rdb.Get(opts.Message).Result()
		if err != nil {
			LogErr("Compress: Redis Error", http.StatusBadRequest, errs)
			return
		}
		w.Write([]byte(x))
		return
	}

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := hellopb.NewHelloServiceClient(cc)
	request := &hellopb.HelloRequest{Name: "Asdas"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)
	w.Write([]byte("Not in cache"))
}

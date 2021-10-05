package router

import (
	"example/models"
	"net/http"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/govalidator"
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
	err = rdb.Set(opts.Message, opts.Message, 0).Err()
	if err != nil {
		LogErr("Compress: Redis error", http.StatusBadRequest, errs)
		w.Write([]byte("Redis error"))
		return
	}
	w.Write([]byte("Not in cache"))
}

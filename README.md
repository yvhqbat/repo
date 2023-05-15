# repo

## redis
- https://github.com/redis/go-redis

## mysql
- https://github.com/go-sql-driver/mysql
- https://github.com/go-gorm/gorm

## kafka
- https://github.com/Shopify/sarama

## 任务队列
- https://github.com/hibiken/asynq
- https://github.com/RichardKnop/machinery

## 限频
- [golang.org/x/time/rate](https://pkg.go.dev/golang.org/x/time/rate)


## 重试延迟策略
```golang
// DefaultRetryDelayFunc is the default RetryDelayFunc used if one is not specified in Config.
// It uses exponential back-off strategy to calculate the retry delay.
func DefaultRetryDelayFunc(n int, e error, t *Task) time.Duration {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := int(math.Pow(float64(n), 4)) + 15 + (r.Intn(30) * (n + 1))
	return time.Duration(s) * time.Second
}
```

## Lua
- http://www.lua.org/
- https://www.coppeliarobotics.com/helpFiles/en/luaCrashCourse.htm



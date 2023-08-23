# repo

## redis
- https://github.com/redis/go-redis

## mysql
- https://github.com/go-sql-driver/mysql
- https://github.com/go-gorm/gorm

环境搭建：
```shell
docker run -dit --rm --name mysql -e "MYSQL_ROOT_PASSWORD=123456" -v /data/mysql/mysql-5.7.36:/var/lib/mysql -p 3306:3306 mysql:5.7.36
```

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

## gateway
- https://github.com/openresty/openresty

## hashids
- https://github.com/speps/go-hashids

## CMD
- https://github.com/spf13/cobra

## Circuit Breaker
- https://github.com/sony/gobreaker

## Design Patterns
- https://learn.microsoft.com/en-us/previous-versions/msp-n-p/dn600223(v=pandp.10)

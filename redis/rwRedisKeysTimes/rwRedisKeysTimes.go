package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var (
	version       bool
	r             *rand.Rand
	clientcluster *redis.ClusterClient
	client        *redis.Client
	ConfCmd       *RedisConfCmd = &RedisConfCmd{}
	WG            sync.WaitGroup
)

const (
	C_Version = "rwRedisKeysTimes V0.1 By WangJiemin.\n\tE_mail: 278667010@qq.com"
)

type RedisConfCmd struct {
	db   int
	host string
	//port  []string
	pass     string
	mode     string
	timedate time.Duration
	total    int
}

// 时间种子
func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
	//rand.Seed(time.Now().UnixNano())
}

func (this *RedisConfCmd) ParseCmdOptions() () {

	defer WG.Done()

	flag.Usage = func() {
		this.PrintUsageMsg()
	}

	flag.BoolVar(&version, "v", false, "print version")
	flag.IntVar(&this.db, "d", 0, "Redis databases. default: 0")
	flag.StringVar(&this.host, "h", "127.0.0.1:6379", "Redis address and port. default: 127.0.0.1:6379")
	flag.StringVar(&this.pass, "p", "", "Redis password. default: \"\"")
	flag.DurationVar(&this.timedate, "t", 30, "Redis keys expiration time. default: 30s")
	flag.IntVar(&this.total, "k", 10, "Redis Keys number. default: 10")
	flag.StringVar(&this.mode, "m", "cluster", "Redis models. options: client/cluster. default: cluster")

	flag.Parse()

	if version {
		fmt.Printf("%s\n", C_Version)
		os.Exit(0)
	}
	if this.mode == "" {
		log.Fatalln("Redis Cluster failed to establish a connection.")
		os.Exit(1)
	}

	if this.mode == "cluster" {
		RedisClusterConnentBaseStr(this)
		RedisClusterReadWirterKeysTimes(clientcluster, this.total, this.timedate)
		time.Sleep(time.Second * 2)
		fmt.Printf("Redis %s 测试完成!\n", "Cluster")
	} else {
		RedisConnentBaseSrt(this)
		RedisClientReadWirterKeysTimes(client, this.total, this.timedate)
		time.Sleep(time.Second * 2)
		fmt.Printf("Redis %s 测试完成!\n", "Master/Slave")
	}
}

func (this *RedisConfCmd) PrintUsageMsg() {
	//flag.Usage()
	var logo = `
***************************************************************************************
	system_command: {$command}                                                        
	system_goos: {$goos}                                                              
	system_arch: {$arch}                                                              
	hostname: {$hostname}                                                             
	hostaddress: {$host_ip}                                                           
	blog: {$url}                                                                      
***************************************************************************************
	`
	arch := fmt.Sprint(runtime.GOARCH)
	good := fmt.Sprint(runtime.GOOS)
	hostname, host_ip := GetSystemHomeNameAndAdderss()
	fmt.Printf("%s\n", C_Version)
	logo = strings.Replace(logo, "{$command}", os.Args[0], -1)
	logo = strings.Replace(logo, "{$arch}", arch, -1)
	logo = strings.Replace(logo, "{$goos}", good, -1)
	logo = strings.Replace(logo, "{$hostname}", hostname, -1)
	logo = strings.Replace(logo, "{$host_ip}", host_ip, -1)
	logo = strings.Replace(logo, "{$url}", "https://jiemin.wang", -1)
	fmt.Println(logo)
	flag.PrintDefaults()

}

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func RedisClusterConnentBaseStr(cfg *RedisConfCmd) (*redis.ClusterClient) {
	url := strings.Split(cfg.host, ",")
	clientcluster = redis.NewClusterClient(&redis.ClusterOptions{
		//Addrs:    []string{fmt.Sprintf("%s", cfg.host)},
		Addrs:    url,
		Password: cfg.pass,
	})
	err := clientcluster.Ping().Err()
	if err == nil {
		log.Println("Redis Cluster established a successful connection.")
	} else {
		log.Fatalln("Redis Cluster failed to establish a connection. ERROR: ", err)
	}
	return clientcluster
}

func RedisConnentBaseSrt(cfg *RedisConfCmd) (*redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s", cfg.host),
		Password: cfg.pass,
		DB:       cfg.db,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

// 获取 系统的hostname 和 ip地址
func GetSystemHomeNameAndAdderss() (hostname string, address string) {
	// get system hostname
	host, err := os.Hostname()
	if err != nil {
		log.Fatalln("fail to get system hostname ERROR: ", err)
	} else {
		hostname = host
	}

	// get system address
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln("fail to get system adderss. ERROR: ", err)
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, add := range addrs {
				if ipnet, ok := add.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						//address = append(address, ipnet.IP.String())
						address = ipnet.IP.String()
					}
				}
			}
		}
	}

	return hostname, address
}

func RedisClusterReadWirterKeysTimes(clientCluster *redis.ClusterClient, total int, timedate time.Duration) () {
	var (
		setavg float64
		setmin float64
		setmax float64
		setsum float64

		getavg float64
		getmin float64
		getmax float64
		getsum float64
	)
	settotaltime := make(map[int]float64)
	gettotaltime := make(map[int]float64)

	for i := 0; i < total; i++ {
		strValues := RandString(8)
		strValuesnum := strValues + "_" + strconv.Itoa(i)
		setsetStart := time.Now()
		//err := clientCluster.Set(strValuesnum, strValues, 30*time.Second).Err()
		//if err != nil {
		//	log.Fatalln("Redis SET Key failed. ERROR: ", err)
		//}
		if timedate == 30 {
			timedate = time.Second * timedate
			err := clientCluster.Set(strValuesnum, strValues, timedate).Err()
			if err != nil {
				log.Fatalln("Redis SET Key failed. ERROR: ", err)
			}
		} else {
			err := clientCluster.Set(strValuesnum, strValues, timedate).Err()
			if err != nil {
				log.Fatalln("Redis SET Key failed. ERROR: ", err)
			}
		}

		setsetEnd := time.Now()
		setsubtime := setsetEnd.Sub(setsetStart)
		settotaltime[i] = float64(setsubtime.Seconds() * 1000)
		//settotaltime[i] = float64(setsubtime.Nanoseconds() / 1000 / 1000)

		getsetStart := time.Now()
		val, err := clientCluster.Get(strValuesnum).Result()
		if err != nil {
			log.Fatalln("Redis GET Key failed. ERROR: ", err)
		}
		getsetEnd := time.Now()
		getsubtime := getsetEnd.Sub(getsetStart)
		gettotaltime[i] = float64(getsubtime.Seconds() * 1000)
		//gettotaltime[i] = float64(getsubtime.Nanoseconds() / 1000 / 1000)
		fmt.Printf("Redis Keys:\t%s,\tRedis values:\t%s\n", strValuesnum, val)
	}

	// set keys times
	for _, v := range settotaltime {
		setsum += v
	}
	for _, v := range settotaltime {
		setmax = settotaltime[0]
		if v > setmax {
			setmax = v
		}
	}

	for _, v := range settotaltime {
		setmax = settotaltime[0]
		if v < setmax {
			setmin = v
		}
	}
	setavg = setsum / float64(len(settotaltime))

	// get keys times
	for _, v := range gettotaltime {
		getsum += v
	}
	for _, v := range gettotaltime {
		getmax = gettotaltime[0]
		if v > getmax {
			getmax = v
		}
	}

	for _, v := range gettotaltime {
		getmax = gettotaltime[0]
		if v < getmax {
			getmin = v
		}
	}
	getavg = getsum / float64(len(gettotaltime))

	fmt.Printf("\nCOMMENT: %d keys\n\n", total)
	fmt.Printf("Performance:\n\tREAD MAX:\t%3.f ms\n\tREAD MIN:\t%3.f ms\n\tREAD AVG:\t%3.f ms\n\t\n\tSET MAX:\t%3.f ms\n\tSET MIN:\t%3.f ms\n\tSET AVG:\t%3.f ms\n\tTOTAL:\n\tREAD SUM:\t%3.f ms\n\tSET SUM:\t%3.f ms\n",
		getmax, getmin, getavg, setmax, setmin, setavg, getsum, setsum)

}

func RedisClientReadWirterKeysTimes(client *redis.Client, total int, timedate time.Duration) () {
	var (
		setavg float64
		setmin float64
		setmax float64
		setsum float64

		getavg float64
		getmin float64
		getmax float64
		getsum float64
	)
	settotaltime := make(map[int]float64)
	gettotaltime := make(map[int]float64)

	for i := 0; i < total; i++ {
		strValues := RandString(8)
		strValuesnum := strValues + "_" + strconv.Itoa(i)
		setsetStart := time.Now()
		//err := client.Set(strValuesnum, strValues, 30*time.Second).Err()
		//if err != nil {
		//	log.Fatalln("Redis SET Key failed. ERROR: ", err)
		//}
		if timedate == 30 {
			timedate = time.Second * timedate
			err := client.Set(strValuesnum, strValues, timedate).Err()
			if err != nil {
				log.Fatalln("Redis SET Key failed. ERROR: ", err)
			}
		} else {
			err := client.Set(strValuesnum, strValues, timedate).Err()
			if err != nil {
				log.Fatalln("Redis SET Key failed. ERROR: ", err)
			}
		}
		setsetEnd := time.Now()
		setsubtime := setsetEnd.Sub(setsetStart)
		settotaltime[i] = float64(setsubtime.Seconds() * 1000)
		//settotaltime[i] = float64(setsubtime.Nanoseconds() / 1000 / 1000)

		getsetStart := time.Now()
		val, err := client.Get(strValuesnum).Result()
		if err != nil {
			log.Fatalln("Redis GET Key failed. ERROR: ", err)
		}
		getsetEnd := time.Now()
		getsubtime := getsetEnd.Sub(getsetStart)
		gettotaltime[i] = float64(getsubtime.Seconds() * 1000)
		//gettotaltime[i] = float64(getsubtime.Nanoseconds() / 1000 / 1000)
		fmt.Printf("Redis Keys:\t%s,\tRedis values:\t%s\n", strValuesnum, val)
	}

	// set keys times
	for _, v := range settotaltime {
		setsum += v
	}
	for _, v := range settotaltime {
		setmax = settotaltime[0]
		if v > setmax {
			setmax = v
		}
	}

	for _, v := range settotaltime {
		setmax = settotaltime[0]
		if v < setmax {
			setmin = v
		}
	}
	setavg = setsum / float64(len(settotaltime))

	// get keys times
	for _, v := range gettotaltime {
		getsum += v
	}
	for _, v := range gettotaltime {
		getmax = gettotaltime[0]
		if v > getmax {
			getmax = v
		}
	}

	for _, v := range gettotaltime {
		getmax = gettotaltime[0]
		if v < getmax {
			getmin = v
		}
	}
	getavg = getsum / float64(len(gettotaltime))

	fmt.Printf("\nCOMMENT: %d keys\n\n", total)
	fmt.Printf("Performance:\n\tREAD MAX:\t%3.f ms\n\tREAD MIN:\t%3.f ms\n\tREAD AVG:\t%3.f ms\n\t\n\tSET MAX:\t%3.f ms\n\tSET MIN:\t%3.f ms\n\tSET AVG:\t%3.f ms\n\tTOTAL:\n\tREAD SUM:\t%3.f ms\n\tSET SUM:\t%3.f ms\n",
		getmax, getmin, getavg, setmax, setmin, setavg, getsum, setsum)

}

func main() {
	WG.Add(1)
	ConfCmd.ParseCmdOptions()
	WG.Wait()

	// 输出随机字符串
	//for i := 0; i < 10; i++ {
	//	fmt.Println(RandString(128))
	//	//time.Sleep(time.Second * 1)
	//}

}

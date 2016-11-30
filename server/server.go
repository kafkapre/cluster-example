package main


import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix.v2/redis"
	"os"
)

type Person struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Comment struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	ParentId string `json:"parent_id"`
	IpAddress string `json:"ip_address"`
}

func createKey(id string)  string {
	return "persons/" + id
}

func (p Person) storePerson(redis *redis.Client)  (string, error){
	key := createKey(p.Id)
	return redis.Cmd("HMSET", key, "id", p.Id, "name", p.Name, "surname", p.Surname).Str()
}

func fetchPerson(redis *redis.Client, key string) (Person, error) {
	res, err := redis.Cmd("HMGET", key, "id", "name", "surname").List()
	return Person{Id:res[0], Name:res[1], Surname:res[2]}, err
}

func existPerson(redis *redis.Client, id string)  (bool, error) {
	res, err := redis.Cmd("EXISTS", createKey(id)).Int()
	return (res == 1), err
}

func httpGetPersons(c *gin.Context, redis *redis.Client) {
	var (
		result gin.H
	)
	id := c.Param("id")

	exist, err := existPerson(redis, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else if  !exist {
		result = gin.H{"Msg": "Not found"}
		c.JSON(http.StatusNotFound, result)
	} else {
		person, err := fetchPerson(redis, createKey(id))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		} else {
			result = gin.H{"person": person}
			c.JSON(http.StatusOK, result)
		}
	}
}

func httpGetAllPersons(c *gin.Context, redis *redis.Client) {
	var (
		persons []Person
	)

	list, err := redis.Cmd("KEYS", "*").List()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	for _, key := range list{
		p, err:= fetchPerson(redis, key)
		if err != nil {
			fmt.Print(err.Error())
		}else{
			persons = append(persons, p)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
		"count":  len(persons),
	})
}




func httpPostPerson(c *gin.Context, redis *redis.Client) {
	var p Person
	c.BindJSON(&p)

	if (len(p.Id) == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id cannot be empty."})
	}else {
		exist, err := existPerson(redis, p.Id)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		} else if  exist {
			result := gin.H{"Msg": fmt.Sprintf("Person with id: %s already exists.", p.Id) }
			c.JSON(http.StatusNotFound, result)
		}else {
			//httpStorePerson(c, redis, p)
		}
	}
}

func httpGetComments(c *gin.Context, commentMap map[string]Comment ) {
	c.JSON(http.StatusOK, gin.H{
		"count":  len(commentMap),
	})
}

func httpPostComment(c *gin.Context, commentMap map[string]Comment) {
	var comment Comment
	c.BindJSON(&comment)

	if (len(comment.Id) == 0) {
		comment.Id = "1"
	}

	if (len(comment.ParentId) == 0) {
			comment.ParentId = nil
	}


	commentMap[comment.Id]=comment

	for key, value := range commentMap {
    fmt.Println("Key:", key, "Value:", value)
  }

}

func main() {

	/*
	redis := connectToRedis()
	if redis == nil {
		fmt.Println("Attempts to connect to redis faild. Server will be stoped")
		syscall.Exit(1)
	}
	*/

	commentMap := make(map[string]Comment)

	x:= commentMap["asda"]
	l:= len(commentMap)
	fmt.Println("Key:", x, l)

	router := gin.Default()

	// GET Comments
	router.GET("/comments", func(c *gin.Context) {httpGetComments(c, commentMap)})

	// POST new Comment
	router.POST("/comments", func(c *gin.Context) {httpPostComment(c, commentMap)})

	router.Run(":3000")
}

func connectToRedis()  (*redis.Client){
		url:=obtainRedisUrl()
		for i:=0; i < 20; i++ {
				redisClient, err := redis.Dial("tcp", url)
				if err == nil {
						fmt.Println("Connected to redis.")
						return redisClient
				}
				fmt.Printf("Tried to connect to redis [attempt: %d].\n", i)
				time.Sleep(5 * time.Second)
		}
		return nil
}

func obtainRedisUrl() string {
	redisIp := os.Getenv("REDIS_IP")
	redisPort := os.Getenv("REDIS_PORT")

	if len(redisIp) == 0 {
		redisIp = "localhost"
	}
	if len(redisPort) == 0 {
		redisPort="6379"
	}

	return redisIp + ":" + redisPort
}

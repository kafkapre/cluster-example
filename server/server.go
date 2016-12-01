package main


import (
	"fmt"
	"net/http"
	"time"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
  "github.com/patrickmn/go-cache"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CommentA struct {
	IdMongo bson.ObjectId `bson:"_id,omitempty"`
	Name string `json:"name"`
	Text string `json:"text"`
	ParentId string `json:"parent_id"`
	IpAddress string `json:"ip_address"`
}

type Comment struct {

	Id string `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	ParentId string `json:"parent_id"`
	IpAddress string `json:"ip_address"`
}

func httpGetComments(c *gin.Context, commentMap map[string]Comment ) {

        timestamp := c.Query("timestamp")
        from := c.Query("from")
        to := c.DefaultQuery("to", "aaaaa")
        x := c.DefaultQuery("x", "bbbb")

        fmt.Println("timestamp: ", timestamp)
        fmt.Println("from: ", from)
        fmt.Println("to: ", to)
        fmt.Println("x: ", x)


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
			//comment.ParentId = nil
	}


	commentMap[comment.Id]=comment

	for key, value := range commentMap {
    fmt.Println("Key:", key, "Value:", value)
  }

}

func main() {


	mongo, err := mgo.Dial("mymongo.org")
  if err != nil {
          panic(err)
  }
  defer mongo.Close()
	mongo.SetMode(mgo.Monotonic, true)


	mongoServerDb := mongo.DB("serverdb").C("comments")

	start := time.Now()
	for i := 1; i <= 0; i++ {
		id:=strconv.Itoa(i)

		text:= "xcv " + id + " lablal"
		err = mongoServerDb.Insert(&Comment{"" ,"Ales", text, "parentid neco", "121.22.11.1"})
		 if err != nil {
						 log.Fatal(err)
		 }
 	}
	elapsed := time.Since(start)
  log.Printf("Binomial took %s", elapsed)

	 var result []CommentA
	 err = mongoServerDb.Find(bson.M{"name": "Ales"}).Sort("-$natural").Skip(3).Limit(20).All(&result)
	 if err != nil {
					 log.Fatal(err)
	 }

	 //fmt.Println("Results All: ", result)


	 for key, value := range result {
	 		fmt.Println("Phone:", key, " ", value.IdMongo.Hex(), " ", value )
		}


  c := cache.New(10 * time.Second, 2*time.Second)
  c.Set("foo", "bar", cache.DefaultExpiration)
  if foo, found := c.Get("foo");  found {
      fmt.Println("11111 " , foo)
  }


  // time.Sleep(5 * time.Second)
	//
	//
  // if foo, found := c.Get("foo");  found {
  //     fmt.Println("2222 ", foo)
  // }else {
	//
	//
	//
  // fmt.Println("3333 ", found)
  // }


	commentMap := make(map[string]Comment)

	// x:= commentMap["asda"]
	// l:= len(commentMap)
	// fmt.Println("Key:", x, l)

	router := gin.Default()

	// GET Comments
	router.GET("/comments", func(c *gin.Context) {httpGetComments(c, commentMap)})

	// POST new Comment
	router.POST("/comments", func(c *gin.Context) {httpPostComment(c, commentMap)})

	router.Run(":3000")
}

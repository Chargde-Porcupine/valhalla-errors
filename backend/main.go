package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type verror struct {
	gorm.Model
	Title    string
	Desc     string
	Category uint
	Status   uint
	Solution string
	ImageL   string `gorm:"column:image_l"`
}

func geneateImagesPaths(db *gorm.DB, router *gin.Engine) {
	var Verrori []verror
	db.Find(&Verrori)
	for _, item := range Verrori {
		router.StaticFile("/images/"+item.ImageL, "./images/"+item.ImageL)
	}
}
func customContains(larger string, slice string) int {
	for _, char := range larger {
		toDel := strings.Index(slice, string(char))
		if toDel != -1 {
			slice = slice[:toDel] + slice[toDel+1:]
		}
	}
	return len(slice)
}

func searchByTitle(inputrarray []verror, title string) []verror {
	var result []verror
	resultMap := make(map[int]verror)
	for _, item := range inputrarray {
		resultMap[customContains(item.Title, title)] = item
	}
	if len(resultMap) == 0 {
		return inputrarray
	}
	keys := make([]int, 0, len(resultMap))
	for k := range resultMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, i := range keys {
		result = append(result, resultMap[i])
	}
	return result
}

func main() {
	db, err := gorm.Open(sqlite.Open("verrori.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&verror{})
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	//generate image paths
	geneateImagesPaths(db, router)
	router.StaticFile("/first", "./images/first.jpg")
	//*router.SetTrustedProxies([]string{"192.168.1.2"})
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"Hello": "ValhallaErrors"})
	})
	//read by title
	router.GET("/find/:title", func(c *gin.Context) {
		var verrorfound []verror
		if err := db.Find(&verrorfound).Error; err != nil {
			c.JSON(400, "Content Not Found")
			return
		}
		verrorfound = searchByTitle(verrorfound, c.Param("title"))
		c.IndentedJSON(http.StatusOK, verrorfound)
	})

	//read all posts
	router.GET("/all", func(c *gin.Context) {
		var verrorfound []verror
		if err := db.Find(&verrorfound).Error; err != nil {
			c.IndentedJSON(400, "Content not Found")
			return
		}
		c.JSON(http.StatusOK, verrorfound)
	})
	//create
	router.POST("/create", func(c *gin.Context) {
		var verrorposted verror

		if err := json.Unmarshal([]byte(c.PostForm("data")), &verrorposted); err != nil {
			c.IndentedJSON(400, "br1")
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			c.IndentedJSON(400, "f")
			return
		}
		// Retrieve file information
		extension := filepath.Ext(verrorposted.ImageL)
		// Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension

		if err := c.SaveUploadedFile(file, "./images/"+newFileName); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		db.Create(&verror{Title: verrorposted.Title, Desc: verrorposted.Desc, Category: verrorposted.Category, Status: verrorposted.Status, Solution: verrorposted.Solution, ImageL: newFileName})

		router.StaticFile("/images/"+newFileName, "./images/"+newFileName)
	})

	//update
	router.POST("/solve/:imageL", func(c *gin.Context) {
		var verrorfound verror
		if err := db.First(&verrorfound, "image_l = ?", c.Param("imageL")).Error; err != nil {
			c.IndentedJSON(400, "Content Not Found")
			return
		}
		var verrorposted verror
		if err := c.BindJSON(&verrorposted); err != nil {
			c.IndentedJSON(400, "Bad Request")
			return
		}
		//TODO: I need a better way of doing this(?)
		db.Model(&verrorfound).Update("Solution", verrorposted.Solution)
		db.Model(&verrorfound).Update("Status", verrorposted.Status)
		fmt.Println(verrorfound.Status)
		fmt.Println(verrorposted.Status)

	})
	router.Run()
}

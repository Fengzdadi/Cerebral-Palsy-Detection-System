package controller

import (
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetBaseInfoHis(c *gin.Context) {
	// The problem here is you cannot ensure the childId will be found by others when they request by this GET method
	childId := c.Query("ChildId")
	childIdInt, _ := strconv.Atoi(childId)
	childIdUint := uint(childIdInt)
	b, err := model.GetBaseInfoHis(childIdUint)
	c.JSON(200, gin.H{
		"message": b,
		"error":   err,
	})

	//userNameValue, exists := c.Get("username")
	//if !exists {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "user_name not found"})
	//	return
	//} else {
	//	fmt.Println(userNameValue)
	//}
	//
	//userid, exists := c.Get("userid")
	//
	//// 写一个断言函数
	//if !exists {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "userid not found"})
	//	return
	//} else {
	//	fmt.Println(userid)
	//	fmt.Printf("Type of userid: %T\n", userid)
	//}
	//
	//if floatUserID, ok := userid.(float64); ok {
	//	uintUserID := uint(floatUserID)
	//	b, err := model.GetBaseInfoHis(uintUserID)
	//	c.JSON(200, gin.H{
	//		"message": b,
	//		"error":   err,
	//	})
	//} else {
	//	c.JSON(400, gin.H{"error": "userid not found"})
	//	return
	//}

}

func AddBaseInfoHis(c *gin.Context) {
	var base model.BaseInfoHis
	if err := c.ShouldBind(&base); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err := base.AddBaseInfoHis()
	c.JSON(200, err)
}

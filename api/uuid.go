package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/database"
)

type MessageStruct struct {
	Action  string `json:"action"`  // 执行方法 token  message
	Message string `json:"message"` // 内容 json 或其他
}

func LinkWS(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": errors.New("请传递UUID"),
		})
		return
	}

	//升级get请求为webSocket协议
	conn, err := database.InitConnection(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		conn.Close()
	}()

	//要求前端发送token标签
	j, _ := json.Marshal(MessageStruct{
		Action:  "token",
		Message: "请传递token",
	})
	if err := conn.WriteMessage(j); err != nil {
		goto END
	}

	for {
		j, err := conn.ReadMessage()
		if err != nil {
			goto END
		}

		var readData MessageStruct
		if err := json.Unmarshal(j, &readData); err != nil {
			j, _ := json.Marshal(map[string]string{
				"action":  "message",
				"message": "数据传递错误，请传递json数据 map[string]string",
			})
			if err := conn.WriteMessage(j); err != nil {
				goto END
			}
			continue
		}
		switch readData.Action {
		case "token":
			//处理token相关
		case "message":
			//跟进传递的数据进行处理
			if err := conn.WriteMessage(j); err != nil {
				goto END
			}
		}
	}

END:
	conn.Close()
	return
}

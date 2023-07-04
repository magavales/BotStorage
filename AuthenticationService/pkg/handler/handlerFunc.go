package handler

import (
	"AuthenticationService/pkg/authentication"
	"AuthenticationService/pkg/database"
	"AuthenticationService/pkg/model"
	"AuthenticationService/pkg/session"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func (h *Handler) Authentication(c *gin.Context) {
	var (
		auth      authentication.Authentication
		resp      Response
		cookie    session.Cookie
		rdb       database.RedisDB
		data      string
		condition bool
		err       error
		location  *time.Location
	)
	ctx := context.Background()
	resp.rw = c.Writer
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		resp.SetStatusBadRequest()
		log.Println("Bad request!")
	}

	reqCookie := c.Request.Header.Get("Set-Cookie")
	if reqCookie == "" {
		data, condition = auth.VerificationPassword(authorization)
		if condition {
			location, err = time.LoadLocation("Europe/Moscow")
			if err != nil {
				return
			}
			now := time.Now().In(location)
			cookie.SetCookie(now)
			rdb.Connect()
			err = rdb.Conn.Do(ctx, "SET", data, cookie.Cookie.String(), "EX", time.Hour*4).Err()
			if err != nil {
				log.Println("Хуета получилась!")
			}
			resp.SetStatusOk()
			resp.SetCookie(cookie.Cookie.String())
		} else {
			resp.SetStatusUnauthorized()
		}
	} else {
		rdb.Connect()
		result, err := rdb.Conn.Get(ctx, data).Result()
		if err != nil {
			log.Println("I can't get data from Redis database!")
		}
		if reqCookie == result {
			resp.SetStatusOk()
		} else {
			resp.SetStatusBadRequest()
		}
	}

}

func (h *Handler) SearchUser(c *gin.Context) {
	var (
		db   database.PostgresDB
		resp Response
		data model.TableData
	)
	resp.rw = c.Writer
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
	}
	db.Connect()
	_, err = db.Access.GetPwd(db.Pool, strconv.FormatInt(data.ChatID, 10))
	if err != nil {
		log.Printf("Don't found password for %d", data.ChatID)
		resp.SetStatusUnauthorized()
	} else {
		resp.SetStatusOk()
	}

}

func (h *Handler) SetUser(c *gin.Context) {
	var (
		auth   authentication.Authentication
		resp   Response
		cookie session.Cookie
		rdb    database.RedisDB
		data   model.TableData
	)
	data.DecodeData(c.Request.Body)
	ctx := context.Background()
	resp.rw = c.Writer
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		resp.SetStatusBadRequest()
		log.Println("Bad request!")
	}

	condition := auth.SetPassword(authorization)
	if condition {
		cookie.SetCookie(time.Now())
		rdb.Connect()
		rdb.Conn.Set(ctx, strconv.FormatInt(data.ChatID, 10), cookie, time.Hour*4)
		resp.SetStatusOk()
		resp.SetCookie(cookie.Cookie.String())
	} else {
		resp.SetStatusUnauthorized()
	}
}

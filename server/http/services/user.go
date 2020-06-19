package services

import (
	"encoding/json"
	"github.com/ameniGa/timeTracker/config"
	"github.com/ameniGa/timeTracker/database"
	"github.com/ameniGa/timeTracker/helpers/context"
	mdl "github.com/ameniGa/timeTracker/models"
	"io/ioutil"
	"net/http"
)

// create a handler struct
type Runner struct {
	Config *config.Config
	Db     database.PresenceHandler
}

func (h Runner) GetUser(res http.ResponseWriter, req *http.Request) {
	request, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	input := &mdl.AuthUserInput{}
	err = json.Unmarshal(request, input)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	ch := make(chan mdl.UserWithError,1)
	ctx, cancel := context.AddTimeoutToCtx(req.Context(), 2)
	defer cancel()
	h.Db.DbGetUserByID(ctx, input.UserID, ch)
	select {
	case out := <-ch:
		if out.Error != nil {
			res.Write([]byte(err.Error()))
			return
		}
		if out.UserInfo.Password != input.Password {
			res.Write([]byte("invalid password"))
			return
		}
		response, err := json.Marshal(out.UserInfo)
		if err != nil {
			res.Write([]byte(err.Error()))
			return
		}
		res.Write(response)
		return
	case <-req.Context().Done():
		res.Write([]byte(req.Context().Err().Error()))
		return
	}
}

func (h Runner) UpdateUser(res http.ResponseWriter, req *http.Request) {
	request, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	input := &mdl.AuthUserInput{}
	err = json.Unmarshal(request, input)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	ch := make(chan error,1)
	ctx, cancel := context.AddTimeoutToCtx(req.Context(), 2)
	defer cancel()
	h.Db.DbUpdateUser(ctx, input.UserID, input.Password, ch)
	select {
	case err := <-ch:
		if err != nil {
			res.Write([]byte(err.Error()))
			return
		}
		response, err := json.Marshal(mdl.GenericRes{
			Status: "successfully updated",
		})
		if err != nil {
			res.Write([]byte(err.Error()))
			return
		}
		res.Write(response)
		return
	case <-req.Context().Done():
		res.Write([]byte(req.Context().Err().Error()))
	}
}

// implement `ServeHTTP` method on `Runner` struct
func (h Runner) GetUserTimeTrack(res http.ResponseWriter, req *http.Request) {
}

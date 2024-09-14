package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getOneMovie(rw http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		//app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(rw, err)
		return
	}

	//app.logger.Println("the id is: ", id)

	movie, err := app.models.DB.Get(id)
	err = app.writeJSON(rw, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(rw, err)
		return
	}

}

func (app *application) getAllMovie(rw http.ResponseWriter, r *http.Request) {

	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(rw, err)
		return
	}

	err = app.writeJSON(rw, http.StatusOK, movies, "movie")
	if err != nil {
		app.errorJSON(rw, err)
		return
	}

}

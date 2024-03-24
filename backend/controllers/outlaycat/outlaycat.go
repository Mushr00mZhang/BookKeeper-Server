package outlaycat

import (
	outlaycat_dto "bookkeeper-backend/dtos/outlaycat"
	"bookkeeper-backend/dtos/pagedlist"
	"bookkeeper-backend/dtos/result"
	"bookkeeper-backend/models/database"
	"bookkeeper-backend/services/outlaycat"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	base := r.PathPrefix("/outlaycats").Subrouter()
	base.HandleFunc("", List).Methods("GET")
	base.HandleFunc("/paged-list", PagedList).Methods("GET")
	base.HandleFunc("/{id}", Get).Methods("GET")
	base.HandleFunc("", Create).Methods("POST")
	base.HandleFunc("/{id}", Update).Methods("PUT")
	base.HandleFunc("/{id}", Delete).Methods("DELETE")
}
func List(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	dto := outlaycat_dto.ParseListDto(values)
	list, code, err := outlaycat.List(database.DB, dto)
	res := result.Dto[[]outlaycat_dto.Dto]{
		Code:   code,
		Result: list,
		Tip:    "查询成功",
		Error:  err,
	}
	if err != nil {
		res.Tip = "查询失败"
		res.Return(&w, http.StatusInternalServerError)
		return
	}
	res.Return(&w, http.StatusOK)
}
func PagedList(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	dto := outlaycat_dto.ParsePagedListDto(values)
	list, code, err := outlaycat.PagedList(database.DB, dto)
	res := result.Dto[pagedlist.Dto[outlaycat_dto.Dto]]{
		Code:   code,
		Result: list,
		Tip:    "查询成功",
		Error:  err,
	}
	if err != nil {
		res.Tip = "查询失败"
		res.Return(&w, http.StatusInternalServerError)
		return
	}
	res.Return(&w, http.StatusOK)
}
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
		res := result.Dto[*outlaycat_dto.Dto]{
			Code:  1,
			Tip:   "解析Id失败",
			Error: err,
		}
		res.Return(&w, http.StatusBadRequest)
		return
	}
	item, code, err := outlaycat.Get(database.DB, id)
	res := result.Dto[*outlaycat_dto.Dto]{
		Code:   0,
		Result: item,
		Tip:    "查询成功",
		Error:  err,
	}
	if err != nil {
		res.Code = code + 1
		res.Tip = "查询失败"
		res.Return(&w, http.StatusInternalServerError)
		return
	}
	res.Return(&w, http.StatusOK)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var dto outlaycat_dto.CreateDto
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取Body失败。%s\n", err)
		res := result.Dto[*uuid.UUID]{
			Code:  1,
			Tip:   "读取Body失败",
			Error: err,
		}
		res.Return(&w, http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		log.Printf("反序列化JSON失败。%s\n", err)
		res := result.Dto[*uuid.UUID]{
			Code:  2,
			Tip:   "反序列化JSON失败",
			Error: err,
		}
		res.Return(&w, http.StatusBadRequest)
		return
	}
	id, code, err := outlaycat.Create(database.DB, dto)
	res := result.Dto[*uuid.UUID]{
		Code:   0,
		Result: id,
		Tip:    "创建成功",
		Error:  err,
	}
	if err != nil {
		res.Code = code + 2
		res.Tip = "创建失败"
		res.Return(&w, http.StatusInternalServerError)
		return
	}
	res.Return(&w, http.StatusOK)
}
func Update(w http.ResponseWriter, r *http.Request) {
	var dto outlaycat_dto.UpdateDto
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
		res := result.Dto[bool]{
			Code:  1,
			Tip:   "解析Id失败",
			Error: err,
		}
		res.Return(&w, http.StatusBadRequest)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取Body失败。%s\n", err)
		res := result.Dto[bool]{
			Code:  2,
			Tip:   "读取Body失败",
			Error: err,
		}
		res.Return(&w, http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		log.Printf("反序列化JSON失败。%s\n", err)
		res := result.Dto[bool]{
			Code:  3,
			Tip:   "反序列化JSON失败",
			Error: err,
		}
		res.Return(&w, http.StatusBadRequest)
		return
	}
	if id != dto.Id {
		log.Printf("Query与Body中的Id不一致。\n")
		res := result.Dto[bool]{
			Code:  4,
			Tip:   "Query与Body中的Id不一致",
			Error: err,
		}
		res.Return(&w, http.StatusBadRequest)
		return
	}
	s, code, err := outlaycat.Update(database.DB, dto)
	res := result.Dto[bool]{
		Code:   0,
		Result: s,
		Tip:    "更新成功",
		Error:  err,
	}
	if err != nil {
		res.Code = code + 4
		res.Tip = "更新失败"
		res.Return(&w, http.StatusInternalServerError)
		return
	}
	res.Return(&w, http.StatusOK)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
		res := result.Dto[bool]{
			Code:   1,
			Result: false,
			Tip:    "解析Id失败",
			Error:  err,
		}
		res.Return(&w, http.StatusBadRequest)
		return
	}
	s, code, err := outlaycat.Delete(database.DB, id)
	res := result.Dto[bool]{
		Code:   0,
		Result: s,
		Tip:    "删除成功",
		Error:  err,
	}
	if err != nil {
		res.Code = code + 1
		res.Tip = "删除失败"
		res.Return(&w, http.StatusInternalServerError)
		return
	}
	res.Return(&w, http.StatusOK)
}

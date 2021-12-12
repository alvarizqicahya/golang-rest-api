package handler

import (
	"golang-rest-api/helper"
	"golang-rest-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MahasiswaHandler struct{}

var mhsModel = new(model.MahasiswaModel)

func (M *MahasiswaHandler) ShowAll(c *gin.Context) {
	var message string

	mhs, err := mhsModel.ShowAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("gagal menampilkan semua data mahasiswa", err.Error()))
		return
	}

	count := len(mhs)
	if count > 0 {
		message = "suskes menampilkan semua data mahasiswa"
	} else {
		message = "data mahasiswa kosong"
	}

	c.JSON(http.StatusOK, helper.BuildResponse(message, mhs))
}

func (M *MahasiswaHandler) Show(c *gin.Context) {
	id := c.Params.ByName("id")

	mhs, err := mhsModel.Find(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("gagal menampilkan data mahasiswa", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse("sukses menampilkan data mahasiswa", mhs))
}

func (M *MahasiswaHandler) Create(c *gin.Context) {
	var mhs model.Mahasiswa

	c.BindJSON(&mhs)

	if err := mhsModel.Save(&mhs); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("gagal menambah data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse("sukses menambah data", mhs))
}

func (M *MahasiswaHandler) Update(c *gin.Context) {
	id := c.Params.ByName("id")

	mhs, err := mhsModel.Find(id)

	if err != nil {
		c.JSON(http.StatusNoContent, helper.BuildErrorResponse("data tidak ditemukan", err.Error()))
		return
	}

	c.BindJSON(&mhs)

	if err = mhsModel.Save(&mhs); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("gagal menyimpan data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse("sukses memperbarui data", mhs))
}

func (M *MahasiswaHandler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := mhsModel.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("gagal menghapus data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse("sukses menghapus data", map[string]string{"id": id}))
}

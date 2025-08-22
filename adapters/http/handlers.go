package http

import (
	"disruptiva.org/specruptiva/pkg/core/service"
	"github.com/gin-gonic/gin"
)

type inputSchema struct {
	Schema     string `form:"schema" json:"schema"`
	ApiVersion string `form:"apiVersion" json:"apiVersion"`
	Kind       string `form:"kind" json:"kind"`
}

type inputData struct {
	Data string `form:"data" json:"data"`
}

type SchemaHandler struct {
	service service.SchemaService
}

type DataHandler struct {
	service service.DataService
}

func NewSchemaHandler(service service.SchemaService) *SchemaHandler {
	return &SchemaHandler{service: service}
}

func NewDataHandler(service service.DataService) *DataHandler {
	return &DataHandler{service: service}
}

func (h *SchemaHandler) Create(c *gin.Context) {
	var inSchema inputSchema
	c.Bind(&inSchema)
	success, err := h.service.Create(inSchema.Schema, inSchema.ApiVersion, inSchema.Kind)
	if err == nil {
		c.JSON(201, success)
	} else {
		c.JSON(422, gin.H{"error": err.Error()})
	}
}

func (h *SchemaHandler) List(c *gin.Context) {
	schemas, err := h.service.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error while listing schemas"})
	} else {
		c.JSON(200, schemas)
	}
}

func (h *SchemaHandler) Read(c *gin.Context) {
	id := c.Params.ByName("id")
	schema, err := h.service.Read(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error while reading schemas"})
	} else if schema.Schema == "" {
		c.JSON(404, gin.H{"error": "Schema not found"})
	} else {
		c.JSON(200, schema)
	}
}

func (h *SchemaHandler) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var schema inputSchema
	c.Bind(&schema)
	if schema.Schema == "" {
		c.JSON(422, gin.H{"error": "Fields are empty"})
		return
	}
	success, err := h.service.Update(id, schema.Schema)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error while updating schema"})
	} else if success.Id == "" {
		c.JSON(404, gin.H{"error": "Schema not found"})
	} else {
		c.JSON(200, success)
	}
}

func (h *SchemaHandler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	success, err := h.service.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error while deleting schema"})
	} else if success.Id == "" {
		c.JSON(404, gin.H{"error": "Schema not found"})
	} else {
		c.JSON(200, success)
	}
}

func (h *DataHandler) Create(c *gin.Context) {
	var inData inputData
	c.Bind(&inData)
	success, err := h.service.Create(inData.Data)
	if err == nil {
		c.JSON(201, success)
	} else {
		c.JSON(422, gin.H{"error": err.Error()})
	}
}

func (h *DataHandler) List(c *gin.Context) {
	datas, err := h.service.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error while listing data"})
	} else {
		c.JSON(200, datas)
	}
}

func (h *DataHandler) Read(c *gin.Context) {
	id := c.Params.ByName("id")
	data, err := h.service.Read(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error while reading data"})
	} else if data.Data == "" {
		c.JSON(404, gin.H{"error": "Data not found"})
	} else {
		c.JSON(200, data)
	}
}

func (h *DataHandler) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var data inputData
	c.Bind(&data)
	if data.Data == "" {
		c.JSON(422, gin.H{"error": "Fields are empty"})
		return
	}
	success, err := h.service.Update(id, data.Data)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error while updating data"})
	} else if success.Id == "" {
		c.JSON(404, gin.H{"error": "Data not found"})
	} else {
		c.JSON(200, success)
	}
}

func (h *DataHandler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	success, err := h.service.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error while deleting data"})
	} else if success.Id == "" {
		c.JSON(404, gin.H{"error": "Data not found"})
	} else {
		c.JSON(200, success)
	}
}

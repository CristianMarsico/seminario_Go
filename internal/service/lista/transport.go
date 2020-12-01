package lista

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/lista",
		function: add(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/lista",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/lista/:ID",
		function: getByID(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/lista/:ID",
		function: delete(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/lista/:ID",
		function: edit(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, err := s.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"no es posible acceder a la lista ": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"lista": l,
		})
	}
}

func getByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("ID"), 10, 64)
		l, err := s.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"no existe tal id en la lista": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"lista": l,
		})
	}
}

func delete(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, _ := strconv.ParseInt(c.Param("ID"), 10, 64)

		err := s.Delete(i)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error!!! ": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}

func add(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := NewBand_Art(c.Query("name"))
		l, err := s.AddList(l)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error al intentar agregar ": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"lista": l,
		})
	}
}

func edit(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		l, _ := strconv.ParseInt(c.Param("ID"), 10, 64)
		n := c.Query("name")

		err := s.Edit(n, l)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"No se ha podido editar ": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}

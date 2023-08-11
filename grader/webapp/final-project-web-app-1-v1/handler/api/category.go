package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	id := r.Context().Value("id").(string)

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	intId, _ := strconv.Atoi(id)
	data, err := c.categoryService.GetCategories(r.Context(), intId)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil || category.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		// log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	// TODO: answer here
	idUser := r.Context().Value("id").(string)

	if idUser == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	intIdUs, _ := strconv.Atoi(idUser)
	var ctg = entity.Category{Type: category.Type, UserID: intIdUs}
	data, err := c.categoryService.StoreCategory(r.Context(), &ctg)
	if err != nil {
		w.WriteHeader(500)
		// r.db.Model(entity.User{}).Where("id = ?", user.ID).Find(&data)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(201)
	jsonData, _ := json.Marshal(map[string]interface{}{
		"user_id":     intIdUs,
		"category_id": data.ID,
		"message":     "success create new category",
	})
	w.Write(jsonData)
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	idUser := r.Context().Value("id").(string)

	categoryID := r.URL.Query().Get("category_id")
	intCatID, _ := strconv.Atoi(categoryID)

	err := c.categoryService.DeleteCategory(r.Context(), intCatID)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(map[string]interface{}{
		"user_id":     idUser,
		"category_id": categoryID,
		"message":     "success delete category",
	})
	w.Write(jsonData)
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}

package brand_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/ernanilima/gshopping/app/utils/response"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// InsertBrand insere uma marca
func (repo *brandRepository) InsertBrand(w http.ResponseWriter, r *http.Request) {

	brand := model.Brand{}

	err := json.NewDecoder(r.Body).Decode(&brand)
	if err != nil {
		messageError := "Erro no corpo recebido, valor inválido"
		response.Error(w, r, http.StatusBadRequest, messageError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(brand); err != nil {
		messageError := "Erro de validação: Marca inválida"
		response.Error(w, r, http.StatusBadRequest, messageError)
		return
	}

	brand, err = repo.BrandRepository.Insert(brand)
	if err != nil {
		messageError := "Marca já existe"
		response.Error(w, r, http.StatusBadRequest, messageError)
		return
	}

	messageSuccess := "Marca inserida com sucesso"
	response.Success(w, http.StatusOK, brand, messageSuccess)
}

// EditBrand edita uma marca
func (repo *brandRepository) EditBrand(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		messageError := "ID inválido"
		response.Error(w, r, http.StatusUnprocessableEntity, messageError)
		return
	}

	brand := model.Brand{}
	err = json.NewDecoder(r.Body).Decode(&brand)
	if err != nil {
		messageError := "Erro no corpo recebido, valor inválido"
		response.Error(w, r, http.StatusBadRequest, messageError)
		return
	}

	brand.ID = id
	validate := validator.New()
	if err := validate.Struct(brand); err != nil {
		messageError := "Erro de validação: Marca inválida"
		response.Error(w, r, http.StatusBadRequest, messageError)
		return
	}

	brand, err = repo.BrandRepository.Edit(brand)
	if err != nil {
		messageError := "Marca já existe"
		response.Error(w, r, http.StatusBadRequest, messageError)
		return
	}

	messageSuccess := "Marca editada com sucesso"
	response.Success(w, http.StatusOK, brand, messageSuccess)
}

// FindAllBrands busca uma lista com todas as marcas
func (repo *brandRepository) FindAllBrands(w http.ResponseWriter, r *http.Request) {
	pagination := utils.PaginationFilters(r)
	response.JSON(w, http.StatusOK, repo.BrandRepository.FindAll(pagination))
}

// FindBrandById busca uma marca pelo ID
func (repo *brandRepository) FindBrandById(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		messageError := "ID inválido"
		response.Error(w, r, http.StatusUnprocessableEntity, messageError)
		return
	}

	brand, err := repo.BrandRepository.FindById(id)
	if err != nil {
		messageError := "Marca não encontrada"
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, brand)
}

// FindAllBrandsByDescription busca uma lista de marcas pela descricao
func (repo *brandRepository) FindAllBrandsByDescription(w http.ResponseWriter, r *http.Request) {

	description := chi.URLParam(r, "description")
	pagination := utils.PaginationFilters(r)
	brands, err := repo.BrandRepository.FindByDescription(description, pagination)
	if err != nil || brands.TotalElements == 0 {
		messageError := fmt.Sprintf("Nenhuma Marca encontrada com '%s'", description)
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, brands)
}

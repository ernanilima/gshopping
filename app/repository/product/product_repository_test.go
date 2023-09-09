package product_repository_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ernanilima/gshopping/app/model"
	product_repository "github.com/ernanilima/gshopping/app/repository/product"
	"github.com/ernanilima/gshopping/app/test/mocks"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Deve inserir um produto
func TestInsertProduct_Should_Insert_A_Product(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db1, mock1, err := sqlmock.New()
	assert.NoError(t, err)
	defer db1.Close()
	mock1.ExpectPing()
	db2, mock2, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()
	mock2.ExpectPing()

	// cria um mock da query executada
	mock1.ExpectExec("INSERT INTO product").WillReturnResult(sqlmock.NewResult(1, 1))
	mock2.ExpectExec("DELETE FROM notfound").WillReturnResult(sqlmock.NewResult(1, 1))
	connector.EXPECT().OpenConnection().Return(db1)
	connector.EXPECT().OpenConnection().Return(db2)

	product := model.Product{
		Barcode:     "78941526300",
		Description: "Produto para teste 1",
		Brand:       model.Brand{ID: uuid.New()},
	}

	result, err := product_repository.NewProductRepository(connector).InsertProduct(product)
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, "78941526300", result.Barcode)
	assert.Equal(t, "Produto para teste 1", result.Description)
	assert.NotNil(t, result.Brand.ID)
	assert.NotNil(t, result.CreatedAt)
}

// Deve editar um produto
func TestEditProduct_Should_Edit_A_Product(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db1, mock1, err := sqlmock.New()
	assert.NoError(t, err)
	defer db1.Close()
	mock1.ExpectPing()
	db2, mock2, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()
	mock2.ExpectPing()
	db3, mock3, err := sqlmock.New()
	assert.NoError(t, err)
	defer db3.Close()
	mock3.ExpectPing()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"id", "barcode", "description", "created_at", "brand_id", "brand_description"}).
		AddRow(uuid.New(), "78900000000", "Produto para teste 1", time.Date(2020, time.January, 1, 20, 30, 40, 0, time.UTC), uuid.New(), "Marca para teste 1")

	// cria um mock da query executada
	mock2.ExpectQuery("SELECT (.+) FROM product p (.+) WHERE p.id").WillReturnRows(rows).RowsWillBeClosed()
	mock1.ExpectExec("UPDATE product SET").WillReturnResult(sqlmock.NewResult(1, 1))
	mock3.ExpectExec("DELETE FROM notfound").WillReturnResult(sqlmock.NewResult(1, 1))
	connector.EXPECT().OpenConnection().Return(db1)
	connector.EXPECT().OpenConnection().Return(db2)
	connector.EXPECT().OpenConnection().Return(db3)

	product := model.Product{
		ID:          uuid.New(),
		Barcode:     "78941526300",
		Description: "Produto para teste EDIT",
		Brand:       model.Brand{ID: uuid.New(), Description: "Marca para teste 1"},
		CreatedAt:   time.Date(2020, time.January, 1, 20, 30, 40, 0, time.UTC),
	}

	result, err := product_repository.NewProductRepository(connector).EditProduct(product)
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, "78941526300", result.Barcode)
	assert.Equal(t, "Produto para teste EDIT", result.Description)
	assert.NotNil(t, result.Brand.ID)
	assert.Equal(t, "Marca para teste 1", result.Brand.Description)
	assert.Equal(t, time.Date(2020, time.January, 1, 20, 30, 40, 0, time.UTC), result.CreatedAt)
}

// Deve retornar todos os produtos
func TestFindAllProducts_Should_Return_All_Products(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"count", "id", "barcode", "description", "brand_id", "brand", "created_at"}).
		AddRow(2, uuid.New(), "789000000001", "Produto para teste 1", uuid.New(), "Marca para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC)).
		AddRow(2, uuid.New(), "789000000002", "Produto para teste 2", uuid.New(), "Marca para teste 2", time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC))

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM product").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result := product_repository.NewProductRepository(connector).FindAllProducts("", utils.Pageable{
		Size: 10,
		Sort: "id ASC",
	})

	// converte para a entidade existente em content
	bts, err := json.Marshal(result.Content)
	assert.NoError(t, err)
	var products []model.Product
	err = json.Unmarshal(bts, &products)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 2, result.TotalElements)
	assert.Equal(t, 10, result.Size)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 2, result.NumberOfElements)
	assert.NotNil(t, products[0].ID)
	assert.Equal(t, "789000000001", products[0].Barcode)
	assert.Equal(t, "Produto para teste 1", products[0].Description)
	assert.NotNil(t, products[0].Brand.ID)
	assert.Equal(t, "Marca para teste 1", products[0].Brand.Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), products[0].CreatedAt)
	assert.NotNil(t, products[1].ID)
	assert.Equal(t, "789000000002", products[1].Barcode)
	assert.Equal(t, "Produto para teste 2", products[1].Description)
	assert.NotNil(t, products[1].Brand.ID)
	assert.Equal(t, "Marca para teste 2", products[1].Brand.Description)
	assert.Equal(t, time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC), products[1].CreatedAt)
}

// Deve retornar um produto ao buscar por id
func TestFindProductById_Should_Return_A_Product_When_Searching_By_ID(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"id", "barcode", "description", "created_at", "brand_id", "brand_description"}).
		AddRow(uuid.New(), "789000000001", "Produto para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), uuid.New(), "Marca para teste 1")

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM product p (.+) WHERE p.id").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result, err := product_repository.NewProductRepository(connector).FindProductById(uuid.New())
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, "789000000001", result.Barcode)
	assert.Equal(t, "Produto para teste 1", result.Description)
	assert.NotNil(t, result.Brand.ID)
	assert.Equal(t, "Marca para teste 1", result.Brand.Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result.CreatedAt)
}

// Deve retornar um produto ao pesquisar por codigo de barras
func TestFindProductByBarcode_Should_Return_Product_When_Searching_By_Barcode(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"id", "barcode", "description", "brand_description", "created_at"}).
		AddRow(uuid.New(), "7891020301", "Produto para teste 1", "Marca para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC))

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM product").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result, err := product_repository.NewProductRepository(connector).FindProductByBarcode("7891020301")
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, "7891020301", result.Barcode)
	assert.Equal(t, "Produto para teste 1", result.Description)
	assert.Equal(t, model.Brand{Description: "Marca para teste 1"}, result.Brand)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result.CreatedAt)
}

// Deve retornar um erro e registar a buscar para produto nao localizado por codigo de barras
func TestFindProductByBarcode_Should_Return_A_Error_And_Register_The_Search_For_Product_Not_Found_By_Barcode(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM product").WillReturnRows(sqlmock.NewRows([]string{})).RowsWillBeClosed()
	mock.ExpectExec("INSERT INTO notfound").WillReturnResult(sqlmock.NewResult(1, 1))
	connector.EXPECT().OpenConnection().Return(db)
	connector.EXPECT().OpenConnection().Return(db)

	var _ interface{}
	_, err = product_repository.NewProductRepository(connector).FindProductByBarcode("7891020301")

	// verifica os resultados
	assert.Error(t, err)
}

// Deve retornar um erro ao nao localizar o produto por codigo de barras e tentar registar a busca
func TestFindProductByBarcode_Should_Return_A_Error_When_Not_Locating_The_Product_By_Barcode_And_Trying_To_Register_The_Search(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM product").WillReturnRows(sqlmock.NewRows([]string{})).RowsWillBeClosed()
	mock.ExpectExec("INSERT INTO notfound").WillReturnError(errors.New("Error"))
	connector.EXPECT().OpenConnection().Return(db)
	connector.EXPECT().OpenConnection().Return(db)

	// capturar a saida do metodo
	output := captureOutput(func() {
		product_repository.NewProductRepository(connector).FindProductByBarcode("7891020301")
	})

	// verifica os resultados
	assert.Equal(t, "Erro ao inserir ProductNotFound com o codigo de barras: 7891020301 | Error", output)
}

// Deve retornar todos os produtos nao encontrados
func TestFindAllProductsNotFound_Should_Return_All_ProductsNotFound(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"count", "id", "barcode", "attempts"}).
		AddRow(2, 1, "789000000091", 11).
		AddRow(2, 2, "789000000092", 22)

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM notfound").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result := product_repository.NewProductRepository(connector).FindAllProductsNotFound(utils.Pageable{
		Size: 10,
		Sort: "id ASC",
	})

	// converte para a entidade existente em content
	bts, err := json.Marshal(result.Content)
	assert.NoError(t, err)
	var productsNotFound []model.ProductNotFound
	err = json.Unmarshal(bts, &productsNotFound)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 2, result.TotalElements)
	assert.Equal(t, 10, result.Size)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 2, result.NumberOfElements)
	assert.NotNil(t, productsNotFound[0].ID)
	assert.Equal(t, "789000000091", productsNotFound[0].Barcode)
	assert.Equal(t, int64(11), productsNotFound[0].Attempts)
	assert.NotNil(t, productsNotFound[1].ID)
	assert.Equal(t, "789000000092", productsNotFound[1].Barcode)
	assert.Equal(t, int64(22), productsNotFound[1].Attempts)
}

// Deve retornar todos os produtos nao encontrados ao pesquisar por codigo de barras
func TestFindAllProductsNotFoundByBarcode_Should_Return_ProductsNotFound_When_Searching_By_Barcode(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"count", "id", "barcode", "attempts"}).
		AddRow(2, 12, "789000000091", 10).
		AddRow(2, 21, "789000000092", 20)

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM notfound nf WHERE nf.barcode").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result, err := product_repository.NewProductRepository(connector).FindAllProductsNotFoundByBarcode("78900000009", utils.Pageable{
		Size: 10,
		Sort: "id ASC",
	})
	assert.NoError(t, err)

	// converte para a entidade existente em content
	bts, err := json.Marshal(result.Content)
	assert.NoError(t, err)
	var productsNotFound []model.ProductNotFound
	err = json.Unmarshal(bts, &productsNotFound)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 2, result.TotalElements)
	assert.Equal(t, 10, result.Size)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 2, result.NumberOfElements)
	assert.NotNil(t, productsNotFound[0].ID)
	assert.Equal(t, "789000000091", productsNotFound[0].Barcode)
	assert.Equal(t, int64(10), productsNotFound[0].Attempts)
	assert.NotNil(t, productsNotFound[1].ID)
	assert.Equal(t, "789000000092", productsNotFound[1].Barcode)
	assert.Equal(t, int64(20), productsNotFound[1].Attempts)
}

// Deve retornar o total de produtos
func TestFindTotalProducts_Should_Return_Total_Products(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(16)

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM product").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result := product_repository.NewProductRepository(connector).FindTotalProducts()

	// verifica os resultados
	assert.Equal(t, int32(16), result)
}

// Deve retornar o total de produtos nao encontrados
func TestFindTotalProductsNotFound_Should_Return_Total_ProductsNotFound(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(17)

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM notfound").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result := product_repository.NewProductRepository(connector).FindTotalProductsNotFound()

	// verifica os resultados
	assert.Equal(t, int32(17), result)
}

// Capturar a saida da funcao passada por parametro
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	buf := bytes.Buffer{}
	io.Copy(&buf, r)
	return buf.String()
}

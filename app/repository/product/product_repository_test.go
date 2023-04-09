package product_repository_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	product_repository "github.com/ernanilima/gshopping/app/repository/product"
	"github.com/ernanilima/gshopping/app/test/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Deve retornar um produto ao pesquisar por codigo de barras
func TestFindByBarcode_Should_Return_Product_When_Searching_By_Barcode(t *testing.T) {
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
	rows := sqlmock.NewRows([]string{"id", "barcode", "description", "brand", "created_at"}).
		AddRow(uuid.New(), "7891020301", "Produto para teste 1", "Marca para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC))

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM product").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result, err := product_repository.NewProductRepository(connector).FindByBarcode("7891020301")
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, "7891020301", result.Barcode)
	assert.Equal(t, "Produto para teste 1", result.Description)
	assert.Equal(t, "Marca para teste 1", result.Brand)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result.CreatedAt)
}

// Deve retornar um erro e registar a buscar para produto nao localizado por codigo de barras
func TestFindByBarcode_Should_Return_A_Error_And_Register_The_Search_For_Product_Not_Found_By_Barcode(t *testing.T) {
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
	_, err = product_repository.NewProductRepository(connector).FindByBarcode("7891020301")

	// verifica os resultados
	assert.Error(t, err)
}

// Deve retornar um erro ao nao localizar o produto por codigo de barras e tentar registar a busca
func TestFindByBarcode_Should_Return_A_Error_When_Not_Locating_The_Product_By_Barcode_And_Trying_To_Register_The_Search(t *testing.T) {
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
		product_repository.NewProductRepository(connector).FindByBarcode("7891020301")
	})

	// verifica os resultados
	assert.Equal(t, "Erro ao inserir o produto com o codigo de barras: 7891020301 | Error", output)
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

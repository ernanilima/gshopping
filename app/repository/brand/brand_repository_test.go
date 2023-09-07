package brand_repository_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ernanilima/gshopping/app/model"
	brand_repository "github.com/ernanilima/gshopping/app/repository/brand"
	"github.com/ernanilima/gshopping/app/test/mocks"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Deve inserir uma marca
func TestInsertBrand_Should_Insert_A_Brand(t *testing.T) {
	// cria um mock para conexao com o banco de dados
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	connector := mocks.NewMockDatabaseConnector(ctrl)

	// cria um mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPing()

	brand := model.Brand{
		Description: "Marca para teste 1",
	}

	// cria um mock da query executada
	mock.ExpectExec("INSERT INTO brand").WillReturnResult(sqlmock.NewResult(1, 1))
	connector.EXPECT().OpenConnection().Return(db)

	result, err := brand_repository.NewBrandRepository(connector).InsertBrand(brand)
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Marca para teste 1", result.Description)
	assert.NotNil(t, result.CreatedAt)
}

// Deve editar uma marca
func TestEditBrand_Should_Edit_A_Brand(t *testing.T) {
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

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"total_products", "id", "code", "description", "created_at"}).
		AddRow(100, uuid.New(), 10, "Marca para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC))

	// cria um mock da query executada
	mock2.ExpectQuery("SELECT (.+) FROM brand b (.+) WHERE b.id").WillReturnRows(rows).RowsWillBeClosed()
	mock1.ExpectExec("UPDATE brand SET").WillReturnResult(sqlmock.NewResult(1, 1))
	connector.EXPECT().OpenConnection().Return(db1)
	connector.EXPECT().OpenConnection().Return(db2)

	brand := model.Brand{
		ID:          uuid.New(),
		Description: "Marca para teste EDIT",
		CreatedAt:   time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC),
	}

	result, err := brand_repository.NewBrandRepository(connector).EditBrand(brand)
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Marca para teste EDIT", result.Description)
	assert.NotNil(t, result.CreatedAt)
}

// Deve deletar uma marca
func TestDeleteBrand_Should_Delete_A_Brand(t *testing.T) {
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

	id := uuid.New()

	// cria um mock dos dados que deveram ser retornados
	rows := sqlmock.NewRows([]string{"total_products", "id", "code", "description", "created_at"}).
		AddRow(100, id, 10, "Marca para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC))

	// cria um mock da query executada
	mock2.ExpectQuery("SELECT (.+) FROM brand b (.+) WHERE b.id").WillReturnRows(rows).RowsWillBeClosed()
	mock1.ExpectExec("DELETE FROM brand").WillReturnResult(sqlmock.NewResult(1, 1))
	connector.EXPECT().OpenConnection().Return(db1)
	connector.EXPECT().OpenConnection().Return(db2)

	result, err := brand_repository.NewBrandRepository(connector).DeleteBrand(id)
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Marca para teste 1", result.Description)
	assert.NotNil(t, result.CreatedAt)
}

// Deve retornar todas as marcas
func TestFindAll_Should_Return_All_Brands(t *testing.T) {
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
	rows := sqlmock.NewRows([]string{"count", "total_products", "id", "code", "description", "created_at"}).
		AddRow(2, 10, uuid.New(), 100, "Marca para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC)).
		AddRow(2, 20, uuid.New(), 200, "Marca para teste 2", time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC))

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM brand").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result := brand_repository.NewBrandRepository(connector).FindAllBrands(utils.Pageable{
		Size: 10,
		Sort: "id ASC",
	})

	// converte para a entidade existente em content
	bts, err := json.Marshal(result.Content)
	assert.NoError(t, err)
	var brands []model.Brand
	err = json.Unmarshal(bts, &brands)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 2, result.TotalElements)
	assert.Equal(t, 10, result.Size)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 2, result.NumberOfElements)
	assert.NotNil(t, brands[0].ID)
	assert.Equal(t, int64(100), brands[0].Code)
	assert.Equal(t, "Marca para teste 1", brands[0].Description)
	assert.Equal(t, int64(10), brands[0].TotalProducts)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), brands[0].CreatedAt)
	assert.NotNil(t, brands[1].ID)
	assert.Equal(t, int64(200), brands[1].Code)
	assert.Equal(t, "Marca para teste 2", brands[1].Description)
	assert.Equal(t, int64(20), brands[1].TotalProducts)
	assert.Equal(t, time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC), brands[1].CreatedAt)
}

// Deve retornar uma marca ao buscar por id
func TestFindById_Should_Return_A_Brand_When_Searching_By_ID(t *testing.T) {
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
	rows := sqlmock.NewRows([]string{"total_products", "id", "code", "description", "created_at"}).
		AddRow(100, uuid.New(), 10, "Marca para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC))

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM brand b (.+) WHERE b.id").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result, err := brand_repository.NewBrandRepository(connector).FindBrandById(uuid.New())
	assert.NoError(t, err)

	// verifica os resultados
	assert.NotNil(t, result.ID)
	assert.Equal(t, int64(10), result.Code)
	assert.Equal(t, "Marca para teste 1", result.Description)
	assert.Equal(t, int64(100), result.TotalProducts)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result.CreatedAt)
}

// Deve retornar as marcas localizadas ao buscar por descricao
func TestFindAll_Should_Return_The_Brands_Found_When_Searching_By_Description(t *testing.T) {
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
	rows := sqlmock.NewRows([]string{"count", "total_products", "id", "code", "description", "created_at"}).
		AddRow(2, 100, uuid.New(), 10, "Marca para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC)).
		AddRow(2, 200, uuid.New(), 20, "Marca para teste 2", time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC))

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM brand b (.+) WHERE UPPER\\(unaccent\\(b.description\\)\\)").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result, err := brand_repository.NewBrandRepository(connector).FindAllBrandsByDescription("teste", utils.Pageable{
		Size: 10,
		Sort: "id ASC",
	})
	assert.NoError(t, err)

	// converte para a entidade existente em content
	bts, err := json.Marshal(result.Content)
	assert.NoError(t, err)
	var brands []model.Brand
	err = json.Unmarshal(bts, &brands)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 2, result.TotalElements)
	assert.Equal(t, 10, result.Size)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 2, result.NumberOfElements)
	assert.NotNil(t, brands[0].ID)
	assert.Equal(t, int64(10), brands[0].Code)
	assert.Equal(t, "Marca para teste 1", brands[0].Description)
	assert.Equal(t, int64(100), brands[0].TotalProducts)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), brands[0].CreatedAt)
	assert.NotNil(t, brands[1].ID)
	assert.Equal(t, int64(20), brands[1].Code)
	assert.Equal(t, "Marca para teste 2", brands[1].Description)
	assert.Equal(t, int64(200), brands[1].TotalProducts)
	assert.Equal(t, time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC), brands[1].CreatedAt)
}

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
	rows := sqlmock.NewRows([]string{"count", "id", "description", "created_at"}).
		AddRow(2, uuid.New(), "Marda para teste 1", time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC)).
		AddRow(2, uuid.New(), "Marda para teste 2", time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC))

	// cria um mock da query executada
	mock.ExpectQuery("SELECT (.+) FROM brand").WillReturnRows(rows).RowsWillBeClosed()
	connector.EXPECT().OpenConnection().Return(db)

	result := brand_repository.NewBrandRepository(connector).FindAll(utils.Pageable{
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
	assert.Equal(t, "Marda para teste 1", brands[0].Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), brands[0].CreatedAt)
	assert.NotNil(t, brands[1].ID)
	assert.Equal(t, "Marda para teste 2", brands[1].Description)
	assert.Equal(t, time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC), brands[1].CreatedAt)
}

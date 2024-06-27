package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobbybof/inventory-api/internal/helper"
	"github.com/bobbybof/inventory-api/internal/repository"
	mock "github.com/bobbybof/inventory-api/internal/repository/mock"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateProductApi(t *testing.T) {
	product := randomProduct()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mock.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":        product.Name,
				"price":       product.Price,
				"description": product.Description,
			},
			buildStubs: func(store *mock.MockStore) {
				arg := repository.CreateProductParams{
					Name:        product.Name,
					Price:       product.Price,
					Description: product.Description,
				}
				store.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			store := mock.NewMockStore(ctr)
			tc.buildStubs(store)

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/product"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))

			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}
}

func TestGetProductsApi(t *testing.T) {
	products := []repository.Product{
		randomProduct(),
		randomProduct(),
		randomProduct(),
	}

	type getProductsResponse struct {
		Data  []repository.Product `json:"data"`
		Total int64                `json:"total"`
	}

	testCases := []struct {
		name          string
		buildStubs    func(store *mock.MockStore)
		queryParams   string
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			queryParams: "limit=10&offset=0",
			buildStubs: func(store *mock.MockStore) {
				arg := repository.GetAllProductsParam{
					Limit:  10,
					Offset: 0,
				}
				store.EXPECT().
					GetAllProducts(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(products, int64(3), nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				var gotProducts getProductsResponse
				err := json.Unmarshal(recorder.Body.Bytes(), &gotProducts)
				require.NoError(t, err)

				require.Equal(t, int64(len(products)), gotProducts.Total)
				require.Equal(t, products, gotProducts.Data)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			store := mock.NewMockStore(ctr)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/products?%s", tc.queryParams)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomProduct() repository.Product {
	return repository.Product{
		ID:          int32(helper.RandomInt(1, 100)),
		Name:        helper.RandomString(5),
		Price:       helper.RandomFloat(0, 999999),
		Description: pgtype.Text{String: helper.RandomString(100), Valid: true},
	}
}

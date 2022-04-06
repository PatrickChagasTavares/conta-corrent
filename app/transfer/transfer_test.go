package transfer

import (
	"errors"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/patrickchagastavares/conta-corrent/mocks"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/test"
	"github.com/stretchr/testify/assert"
)

func TestLisByID(t *testing.T) {
	accountID := 1
	time := time.Now()
	dataSucess := []*model.Transfer{
		{ID: 1, OriginID: accountID, DestinationID: 2, Amount: *big.NewInt(100), CreatedAt: time},
		{ID: 3, OriginID: accountID, DestinationID: 3, Amount: *big.NewInt(100), CreatedAt: time},
	}

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData []*model.Transfer

		InputID int

		PrepareMock func(mock *mocks.MockTransferStore)
	}{
		"return sucess": {
			InputID:      accountID,
			ExpectedData: dataSucess,
			PrepareMock: func(mock *mocks.MockTransferStore) {
				mock.EXPECT().ListByID(gomock.Any(), accountID).Return(dataSucess, nil)
			},
		},
		"return error: id required": {
			InputID:     -1,
			ExpectedErr: errListIDNotInformed,
			PrepareMock: func(mock *mocks.MockTransferStore) {},
		},
		"return error:  store": {
			InputID:     accountID,
			ExpectedErr: errListByID,
			PrepareMock: func(mock *mocks.MockTransferStore) {
				mock.EXPECT().ListByID(gomock.Any(), accountID).Return(nil, errors.New("error"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockTransfer := mocks.NewMockTransferStore(ctrl)

			cs.PrepareMock(mockTransfer)

			app := NewApp(mockTransfer, nil)

			data, err := app.ListByID(ctx, cs.InputID)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestCreate(t *testing.T) {
	transfer := &model.Transfer{OriginID: 1, DestinationID: 2, Amount: *big.NewInt(15)}

	cases := map[string]struct {
		ExpectedErr error

		InputTransfer *model.Transfer

		PrepareMock        func(mock *mocks.MockTransferStore)
		PrepareMockAccount func(mock *mocks.MockAccountStore)
	}{
		"return sucess": {
			InputTransfer: transfer,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByID(gomock.Any(), 1).Return(&model.Account{ID: 1, Balance: *big.NewInt(130)}, nil)
				mock.EXPECT().GetByID(gomock.Any(), 2).Return(&model.Account{ID: 2, Balance: *big.NewInt(130)}, nil)

				mock.EXPECT().UpdateBalance(gomock.Any(), &model.Account{ID: 1, Balance: *big.NewInt(115)}).Return(nil)
				mock.EXPECT().UpdateBalance(gomock.Any(), &model.Account{ID: 2, Balance: *big.NewInt(145)}).Return(nil)
			},
			PrepareMock: func(mock *mocks.MockTransferStore) {
				mock.EXPECT().Create(gomock.Any(), transfer).Return(nil)
			},
		},
		"return error: validate originID": {
			InputTransfer:      &model.Transfer{OriginID: -1},
			ExpectedErr:        model.NewError(http.StatusBadRequest, "O id da sua conta é obrigatório", nil),
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {},
			PrepareMock:        func(mock *mocks.MockTransferStore) {},
		},
		"return error: validate DestinationID": {
			InputTransfer:      &model.Transfer{OriginID: 1, DestinationID: -1},
			ExpectedErr:        model.NewError(http.StatusBadRequest, "O id do destinatário é obrigatório", nil),
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {},
			PrepareMock:        func(mock *mocks.MockTransferStore) {},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockTransfer := mocks.NewMockTransferStore(ctrl)
			mockAccount := mocks.NewMockAccountStore(ctrl)

			cs.PrepareMock(mockTransfer)
			cs.PrepareMockAccount(mockAccount)

			app := NewApp(mockTransfer, mockAccount)

			err := app.Create(ctx, cs.InputTransfer)

			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

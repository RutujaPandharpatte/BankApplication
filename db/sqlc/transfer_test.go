package db

import (
	"context"
	"testing"
	"time"

	"github.com/RutujaPandharpatte/BankApplication/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.WithinDuration(t, transfer.CreatedAt, time.Now(), time.Second)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 2; i++ {
		transfer := createRandomTransfer(t)

		args := ListTransfersParams{
			FromAccountID: transfer.FromAccountID,
			ToAccountID:   transfer.ToAccountID,
		}

		transfers, err := testQueries.ListTransfers(context.Background(), args)

		require.NoError(t, err)
		require.Empty(t, transfers) // Offset is 4 & only 1 transfer is created, so it should be empty

		for _, got := range transfers {
			require.NotEmpty(t, got)
		}
	}
}

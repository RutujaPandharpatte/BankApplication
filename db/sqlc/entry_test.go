package db

import (
	"context"
	"testing"
	"time"

	"github.com/RutujaPandharpatte/BankApplication/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.WithinDuration(t, entry.CreatedAt, time.Now(), time.Second)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 2; i++ {
		entry := createRandomEntry(t)

		args := ListEntriesParams{
			AccountID: entry.AccountID,
			Limit:     1,
		}

		entries, err := testQueries.ListEntries(context.Background(), args)

		require.NoError(t, err)
		require.NotEmpty(t, entries)

		require.Len(t, entries, 1)
		require.Equal(t, entry.ID, entries[0].ID)
		require.Equal(t, entry.AccountID, entries[0].AccountID)
		require.Equal(t, entry.Amount, entries[0].Amount)

		require.WithinDuration(t, entry.CreatedAt, entries[0].CreatedAt, time.Second)
	}
}

package db

import (
	"context"
	"testing"
	"time"

	"github.com/angrypenguin1995/simple__bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	CreateRandomEntry(t, account1)
}

func TestGetEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	entry1 := CreateRandomEntry(t, account1)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {
	account1 := createRandomAccount(t)
	for i := 1; i < 11; i++ {
		CreateRandomEntry(t, account1)
	}

	args := ListEntriesParams{
		AccountID: account1.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)

	for i := 0; i < 5; i++ {
		entry := entries[i]
		require.Equal(t, entry.AccountID, account1.ID)
		require.NotEmpty(t, entry)
	}

}

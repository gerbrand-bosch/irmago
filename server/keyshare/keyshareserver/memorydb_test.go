package keyshareserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryDBUserManagement(t *testing.T) {
	db := NewMemoryDatabase()

	user := &KeyshareUser{Username: "testuser"}
	err := db.NewUser(user)
	require.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)

	nuser, err := db.User("testuser")
	require.NoError(t, err)
	assert.Equal(t, "testuser", nuser.Username)

	_, err = db.User("nonexistent")
	assert.Error(t, err)

	user = &KeyshareUser{Username: "testuser"}
	err = db.NewUser(user)
	assert.Error(t, err)

	err = db.UpdateUser(nuser)
	assert.NoError(t, err)

	err = db.AddEmailVerification(nuser, "test@test.com", "testtoken")
	assert.NoError(t, err)

	err = db.AddLog(nuser, PinCheckSuccess, nil)
	assert.NoError(t, err)

	ok, tries, wait, err := db.ReservePincheck(nuser)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.True(t, tries > 0)
	assert.Equal(t, int64(0), wait)

	err = db.SetSeen(nuser)
	assert.NoError(t, err)
}
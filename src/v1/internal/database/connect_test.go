package database

import "testing"

func TestConnectPostgres(t *testing.T) {
	_, err := ConnectPostgres()
	if err != nil {
		t.Error(err)
	}
}

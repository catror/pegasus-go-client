package session

import (
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/assert"
)

func TestReplicaManager_GetReplica(t *testing.T) {
	defer leaktest.Check(t)()

	rm := NewReplicaManager()
	defer rm.Close()

	r1 := rm.GetReplica("127.0.0.1:34802")
	assert.Equal(t, len(rm.replicas), 1)

	r2 := rm.GetReplica("127.0.0.1:34802")
	assert.Equal(t, r1, r2)
}

func TestReplicaManager_Close(t *testing.T) {
	defer leaktest.Check(t)()

	rm := NewReplicaManager()
	defer rm.Close()

	rm.GetReplica("127.0.0.1:34801")
	rm.GetReplica("127.0.0.1:34802")
	rm.GetReplica("127.0.0.1:34803")

	time.Sleep(time.Second)
}

package tests

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/models"
	"testing"
	"time"
)

type TestSubscription struct {
	t *testing.T
	s *database.Subscription
}

func (ts *TestSubscription) create(subscription models.CreateSubscription, status int) string {
	id, res := ts.s.Create(&subscription)
	validateError(ts.t, res, status)
	if status == 200 {
		return id
	}
	return ""
}

func (ts *TestSubscription) update(id string, subscription models.UpdateSubscription, status int) {
	err := ts.s.Update(id, &subscription)
	validateError(ts.t, err, status)
}

func (ts *TestSubscription) getById(id string, status int) {
	_, err := ts.s.GetById(id)
	validateError(ts.t, err, status)
}

func (ts *TestSubscription) getByName(name string, status int) {
	_, err := ts.s.GetByName(name)
	validateError(ts.t, err, status)
}

func (ts *TestSubscription) getAll(status int) {
	_, err := ts.s.GetAll()
	validateError(ts.t, err, status)
}

func (ts *TestSubscription) delete(id string, status int) {
	err := ts.s.Disable(id)
	validateError(ts.t, err, status)
}

func (ts *TestSubscription) Restore(id string, status int) {
	err := ts.s.Restore(id)
	validateError(ts.t, err, status)
}

func TestSubscriptionWorkflow(t *testing.T) {
	t.Run("Initialize DB", testDatabase)
	name := time.Now().String()
	testCorrectSubscription := models.CreateSubscription{
		Name:      name,
		Frequency: "* * * * *",
		RateLimit: 2,
	}
	testInvalidFrequencySubscription := models.CreateSubscription{
		Name:      "test2",
		Frequency: "* * * * * rnd",
		RateLimit: 2,
	}
	testDuplicateSubscription := models.CreateSubscription{
		Name:      name,
		Frequency: "* * * * *",
		RateLimit: 2,
	}
	testWrongRateLimitSubscription := models.CreateSubscription{
		Name:      "test2",
		Frequency: "* * * * *",
		RateLimit: -1,
	}
	testWrongSubscription := models.CreateSubscription{
		Name:      "test3",
		Frequency: "* * * * *",
		RateLimit: 2,
	}

	var id1 string
	var id2 string
	t.Run("Create", func(t *testing.T) {
		ts := TestSubscription{t: t, s: &database.Subscription{}}
		id1 = ts.create(testCorrectSubscription, 200)
		id2 = ts.create(testDuplicateSubscription, 409)
		ts.create(testInvalidFrequencySubscription, 400)
		ts.create(testWrongRateLimitSubscription, 409)
		ts.create(testWrongSubscription, 409)
	})
	t.Run("GetById", func(t *testing.T) {
		ts := TestSubscription{t: t, s: &database.Subscription{}}
		ts.getById(id1, 200)
		ts.getById(id2, 404)
	})
	t.Run("GetByName", func(t *testing.T) {
		ts := TestSubscription{t: t, s: &database.Subscription{}}
		ts.getByName(name, 200)
		ts.getByName("test2", 404)
	})
	t.Run("Update", func(t *testing.T) {
		ts := TestSubscription{t: t, s: &database.Subscription{}}
		testUpdateSubscription := models.UpdateSubscription{
			Name:      name + "1",
			Frequency: "* * * * *",
			RateLimit: 3,
		}
		testUpdateWrongSubscription := models.UpdateSubscription{
			Name:      name + "2",
			RateLimit: -10,
		}
		ts.update(id1, testUpdateSubscription, 200)
		ts.update(id1, testUpdateWrongSubscription, 400)
		ts.update(id2, testUpdateSubscription, 404)
	})
	t.Run("Delete", func(t *testing.T) {
		ts := TestSubscription{t: t, s: &database.Subscription{}}
		ts.delete(id1, 200)
		ts.getById(id1, 404)
		ts.delete(id2, 404)
	})
	t.Run("Restore", func(t *testing.T) {
		ts := TestSubscription{t: t, s: &database.Subscription{}}
		ts.Restore(id1, 200)
		ts.getById(id1, 200)
		ts.Restore(id2, 404)
	})
}

func TestSubscriptionGetAll(t *testing.T) {
	t.Run("Initialize DB", testDatabase)
	ts := TestSubscription{t: t, s: &database.Subscription{}}
	ts.getAll(200)
}

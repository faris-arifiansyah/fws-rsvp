package repository

import (
	"context"
	"time"

	"github.com/faris-arifiansyah/fws-rsvp/constants"

	"github.com/globalsign/mgo/bson"

	rsvp "github.com/faris-arifiansyah/fws-rsvp"
	"github.com/faris-arifiansyah/mgoi"
)

type mongoRsvp struct {
	db mgoi.DatabaseManager
}

func NewMongoRsvp(db mgoi.DatabaseManager) rsvp.RsvpRepo {
	return &mongoRsvp{db}
}

func (mr *mongoRsvp) CreateRsvp(ctx context.Context, rp rsvp.Rsvp) (rsvp.Rsvp, error) {
	rp.ID = bson.NewObjectId()
	rp.CreatedAt = time.Now()

	return rp, mr.db.C("rsvps").Insert(rp)
}

func (mr *mongoRsvp) GetRsvps(ctx context.Context, p *rsvp.Parameter) (*rsvp.RsvpResult, error) {
	var rsvpResult rsvp.RsvpResult

	query := mr.db.C("rsvps").Find(nil)
	query.Sort(p.Sort)

	if p.Limit != constants.NoLimit {
		query.Skip(p.Offset)
		query.Limit(p.Limit)
	}

	total, err := query.Count()
	if err != nil {
		return nil, err
	}

	err = query.All(&rsvpResult.Data)
	rsvpResult.Total = int64(total)

	return &rsvpResult, err
}

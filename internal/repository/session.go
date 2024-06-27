package repository

import (
	"context"
	"github.com/dennis0126/network-monitor/internal/db"
	"github.com/dennis0126/network-monitor/internal/model"
)

type SessionRepository struct {
	ctx     context.Context
	queries *db.Queries
}

func NewSessionRepository(ctx context.Context, queries *db.Queries) SessionRepository {
	return SessionRepository{ctx: ctx, queries: queries}
}

func (r SessionRepository) CreateSession(userId string, ipAddress string, userAgent string) (model.Session, error) {
	dbSession, err := r.queries.CreateSession(r.ctx, db.CreateSessionParams{UserID: userId, IpAddress: ipAddress, UserAgent: userAgent})
	if err != nil {
		return model.Session{}, err
	}
	return unmarshallDbSession(dbSession), nil
}

func (r SessionRepository) GetSessionById(id string) (model.Session, error) {
	dbSession, err := r.queries.GetSessionById(r.ctx, id)
	if err != nil {
		return model.Session{}, err
	}
	return unmarshallDbSession(dbSession), nil
}

func (r SessionRepository) DeleteSessionById(id string) error {
	err := r.queries.DeleteSessionById(r.ctx, id)
	return err
}

func unmarshallDbSession(session db.Session) model.Session {
	return model.Session{
		ID:           session.ID,
		UserID:       session.UserID,
		IpAddress:    session.IpAddress,
		UserAgent:    session.UserAgent,
		LastActivity: session.LastActivity.Time,
	}
}

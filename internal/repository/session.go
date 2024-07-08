package repository

import (
	"context"
	"errors"
	"github.com/dennis0126/network-monitor/internal/db"
	"github.com/dennis0126/network-monitor/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type SessionRepository struct {
	ctx     context.Context
	queries *db.Queries
}

func NewSessionRepository(ctx context.Context, queries *db.Queries) SessionRepository {
	return SessionRepository{ctx: ctx, queries: queries}
}

func (r SessionRepository) CreateSession(id string, userId string, ipAddress string, userAgent string, lastActivity time.Time) (model.Session, error) {
	dbSession, err := r.queries.CreateSession(r.ctx, db.CreateSessionParams{
		ID: id, UserID: userId, IpAddress: ipAddress, UserAgent: userAgent, LastActivity: pgtype.Timestamptz{Time: lastActivity, Valid: true},
	})
	if err != nil {
		return model.Session{}, err
	}
	return unmarshallDbSession(dbSession), nil
}

func (r SessionRepository) GetSessionById(id string) (*model.Session, error) {
	dbSession, err := r.queries.GetSessionById(r.ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return &model.Session{}, err
	}

	session := unmarshallDbSession(dbSession)
	return &session, nil
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

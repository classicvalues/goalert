package graphqlapp

import (
	context "context"

	"github.com/target/goalert/assignment"
	"github.com/target/goalert/graphql2"

	"github.com/pkg/errors"
)

type Target App

func (a *App) Target() graphql2.TargetResolver { return (*Target)(a) }

func (t *Target) Name(ctx context.Context, raw *assignment.RawTarget) (*string, error) {
	if raw.Name != "" {
		return &raw.Name, nil
	}
	switch raw.Type {
	case assignment.TargetTypeNotificationChannel:
		nc, err := (*App)(t).FindOneNC(ctx, raw.ID)
		if err != nil {
			return nil, err
		}
		return &nc.Name, nil
	case assignment.TargetTypeRotation:
		r, err := (*App)(t).FindOneRotation(ctx, raw.ID)
		if err != nil {
			return nil, err
		}
		return &r.Name, nil
	case assignment.TargetTypeUser:
		u, err := (*App)(t).FindOneUser(ctx, raw.ID)
		if err != nil {
			return nil, err
		}
		return &u.Name, nil
	case assignment.TargetTypeEscalationPolicy:
		ep, err := (*App)(t).FindOnePolicy(ctx, raw.ID)
		if err != nil {
			return nil, err
		}
		return &ep.Name, nil
	case assignment.TargetTypeSchedule:
		sched, err := (*App)(t).FindOneSchedule(ctx, raw.ID)
		if err != nil {
			return nil, err
		}
		return &sched.Name, nil
	case assignment.TargetTypeService:
		svc, err := (*App)(t).FindOneService(ctx, raw.ID)
		if err != nil {
			return nil, err
		}
		return &svc.Name, nil

	}

	return nil, errors.New("unhandled target type")
}

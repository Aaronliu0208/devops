package model

import (
	"context"

	"casicloud.com/ylops/backend/internal/app/icontext"
	"casicloud.com/ylops/backend/internal/app/model"
	"casicloud.com/ylops/backend/pkg/errors"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ model.ITrans = new(Trans)

// TransSet 注入Trans
var TransSet = wire.NewSet(wire.Struct(new(Trans), "*"), wire.Bind(new(model.ITrans), new(*Trans)))

// Trans 事务管理
type Trans struct {
	Client *mongo.Client
}

// Exec 执行事务
func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}

	session, err := a.Client.StartSession()
	if err != nil {
		return errors.WithStack(err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		err := fn(icontext.NewTrans(sessCtx, true))
		return nil, err
	})

	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

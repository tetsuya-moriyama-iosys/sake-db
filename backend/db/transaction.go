/**
*  トランザクション用ヘルパ
 */

package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type TransactionFunc[R any] func(sc mongo.SessionContext) (R, error)

func WithTransaction[R any](ctx context.Context, client *mongo.Client, fn TransactionFunc[R]) (R, error) {
	var res R
	session, err := client.StartSession()
	if err != nil {
		return res, err
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
			return err
		}

		// トランザクション関数を実行
		result, err := fn(sc)
		if err != nil {
			if abortErr := session.AbortTransaction(sc); abortErr != nil {
				log.Printf("Failed to abort transaction: %v", abortErr)
			}
			return err
		}

		res = result

		if err := session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	return res, err
}

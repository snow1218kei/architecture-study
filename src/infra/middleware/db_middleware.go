package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/yuuki-tsujimura/architecture-study/src/config"
	"log"
	"net/http"
)

var (
	transactionMethods = []string{
		"POST",
		"PUT",
		"DELETE",
	}
)

func DBMiddleware(conn *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tx *sqlx.Tx
		if isTransactionMethod(c.Request.Method) {
			log.Println("start transaction")
			transaction, err := conn.BeginTxx(context.Background(), nil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction error"})
				c.Abort()
				return
			}
			tx = transaction

			c.Set(config.DBKey, transaction)
		} else {
			c.Set(config.DBKey, conn)
		}

		c.Next()

		if c.IsAborted() && tx != nil {
			tx.Rollback()
			log.Println("rollback")
			return
		}

		if tx != nil {
			if err := tx.Commit(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction commit error"})
				c.Abort()
				return
			}
			log.Println("committed")
		}
	}
}

func isTransactionMethod(httpMethod string) bool {
	for _, tMethod := range transactionMethods {
		if httpMethod == tMethod {
			return true
		}
	}
	return false
}

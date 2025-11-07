/* *
 * Test Dao
 */

package dao

import (
	"context"
	"io"

	"github.com/tonly18/xgin/xerror"
)

type TestDao struct {
	ctx context.Context
}

func NewTestDao(ctx context.Context) *TestDao {
	return &TestDao{
		ctx,
	}
}

func (d *TestDao) GetData(x int) (string, xerror.Error) {
	if x == 0 {
		//return "", xerror.NewXError("test-dao-error")
		return "", xerror.Wrap(io.EOF, "test-dao-error")
	}

	return "test dao", nil
}

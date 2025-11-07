/* *
 * Test Model
 */

package model

import (
	"context"
	"fmt"

	"github.com/tonly18/xgin/example/model/dao"
	"github.com/tonly18/xgin/xerror"
)

type TestModel struct {
	ctx context.Context
	*dao.TestDao
}

func NewTestModel(ctx context.Context) *TestModel {
	return &TestModel{
		ctx,
		dao.NewTestDao(ctx),
	}
}

// 创建收款订单
func (m *TestModel) GetData(x int) (string, xerror.Error) {
	data, xerr := m.TestDao.GetData(x)
	if xerr != nil {
		return "", xerror.Wrap(xerr, "test-model-error")
	}

	return fmt.Sprintf("%s_%s", data, "test-model"), nil
}

/* *
 * Test Service
 */

package service

import (
	"context"
	"fmt"

	"github.com/tonly18/xgin/example/model"
	"github.com/tonly18/xgin/xerror"
)

type TestService struct {
	ctx context.Context
}

func NewTestService(ctx context.Context) *TestService {
	return &TestService{
		ctx: ctx,
	}
}

func (s *TestService) GetData(x int) (string, xerror.Error) {
	testModel := model.NewTestModel(s.ctx)
	data, xerr := testModel.GetData(x)
	if xerr != nil {
		return "", xerror.Wrap(xerr, "test-service-error")
	}

	return fmt.Sprintf("%s_%s", data, "test-service"), nil
}

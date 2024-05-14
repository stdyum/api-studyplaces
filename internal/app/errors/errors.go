package errors

import (
	"net/http"

	grpcErr "github.com/stdyum/api-common/grpc"
	httpErr "github.com/stdyum/api-common/http"
	"github.com/stdyum/api-studyplaces/internal/app/controllers"
	"google.golang.org/grpc/codes"
)

var (
	HttpErrorsMap = map[error]any{
		controllers.ErrNoPermissions: http.StatusForbidden,
		controllers.ErrValidation:    http.StatusUnprocessableEntity,
	}

	GRpcErrorsMap = map[error]any{
		controllers.ErrNoPermissions: codes.PermissionDenied,
		controllers.ErrValidation:    codes.InvalidArgument,
	}
)

func Register() {
	httpErr.RegisterErrors(HttpErrorsMap)
	grpcErr.RegisterErrors(GRpcErrorsMap)
}

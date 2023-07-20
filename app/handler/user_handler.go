package handler

import (
	"net/http"
	"strings"

	"github.com/fuku01/go-test-api/app/usecase"
	"github.com/labstack/echo/v4"
)

// @ 「User」に関する、handlerメソッドの集まり（インターフェース）を定義。
type UserHandler interface {
	GetLoginUser(c echo.Context) error // ログイン中のユーザー情報を取得するメソッドを定義
}

// @ 構造体の型。
type userHandler struct {
	userUsecase usecase.UserUsecase
}

// @ /mainのルーティングで、この構造体を使用する（呼び出す）ための関数を定義。
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

// @ フロントからのHTTPリクエストを受け取り、/usecase層で実装した【具体的な処理】を呼び出し、フロントへ返すレスポンスを生成。

func (h userHandler) GetLoginUser(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization") // リクエストヘッダーからAuthorizationを取得
	token := strings.TrimPrefix(authHeader, "Bearer ")    // Bearerを削除

	// ログイン中のユーザー情報を取得
	user, err := h.userUsecase.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		return err
	}

	// ログイン中のユーザー情報を返す
	return c.JSON(http.StatusOK, user)
}

package hub

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/rss3-global-indexer/internal/database"
)

func (h *Hub) GetEpochsHandler(c echo.Context) error {
	var request GetEpochsRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("bad request: %v", err))
	}

	if err := defaults.Set(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("set default failed: %v", err))
	}

	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate failed: %v", err))
	}

	epoch, err := h.databaseClient.FindEpochs(c.Request().Context(), request.Limit, request.Cursor)
	if err != nil {
		if errors.Is(err, database.ErrorRowNotFound) {
			return c.NoContent(http.StatusNotFound)
		}

		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("get failed: %v", err))
	}

	var cursor string
	if len(epoch) > 0 && len(epoch) == request.Limit {
		cursor = fmt.Sprintf("%d", epoch[len(epoch)-1].ID)
	}

	return c.JSON(http.StatusOK, Response{
		Data:   epoch,
		Cursor: cursor,
	})
}

func (h *Hub) GetEpochHandler(c echo.Context) error {
	var request GetEpochRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("bad request: %v", err))
	}

	if err := defaults.Set(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("set default failed: %v", err))
	}

	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate failed: %v", err))
	}

	epoch, err := h.databaseClient.FindEpoch(c.Request().Context(), request.ID, request.ItemsLimit, request.Cursor)
	if err != nil {
		if errors.Is(err, database.ErrorRowNotFound) {
			return c.NoContent(http.StatusNotFound)
		}

		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("get failed: %v", err))
	}

	var cursor string
	if len(epoch.RewardItems) > 0 && len(epoch.RewardItems) == request.ItemsLimit {
		cursor = fmt.Sprintf("%d", epoch.RewardItems[len(epoch.RewardItems)-1].Index)
	}

	return c.JSON(http.StatusOK, Response{
		Data:   epoch,
		Cursor: cursor,
	})
}

func (h *Hub) GetEpochNodeRewardsHandler(c echo.Context) error {
	var request GetEpochNodeRewardsRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("bad request: %v", err))
	}

	if err := defaults.Set(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("set default failed: %v", err))
	}

	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate failed: %v", err))
	}

	epoch, err := h.databaseClient.FindEpochNodeRewards(c.Request().Context(), request.NodeAddress, request.Limit, request.Cursor)
	if err != nil {
		if errors.Is(err, database.ErrorRowNotFound) {
			return c.NoContent(http.StatusNotFound)
		}

		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("get failed: %v", err))
	}

	var cursor string
	if len(epoch) > 0 && len(epoch) == request.Limit {
		cursor = fmt.Sprintf("%d", epoch[len(epoch)-1].ID)
	}

	return c.JSON(http.StatusOK, Response{
		Data:   epoch,
		Cursor: cursor,
	})
}

type GetEpochsRequest struct {
	Cursor *string `query:"cursor"`
	Limit  int     `query:"limit" validate:"min=1,max=50" default:"10"`
}

type GetEpochRequest struct {
	ID         uint64  `param:"id" validate:"required"`
	ItemsLimit int     `query:"itemsLimit" validate:"min=1,max=50" default:"10"`
	Cursor     *string `query:"cursor"`
}

type GetEpochNodeRewardsRequest struct {
	NodeAddress common.Address `param:"node" validate:"required"`
	Limit       int            `query:"limit" validate:"min=1,max=50" default:"10"`
	Cursor      *string        `query:"cursor"`
}
package routes

import (
	"net/http"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
	"strconv"
)

func NewsList(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		count = 20
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	news, totalCount, err := models.FindAllNews(offset, count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	totalPages := totalCount / count
	if totalPages == 0 {
		totalPages = 1
	}

	httpUtils.RenderAdmin("admin/templates/news/overview.gohtml", OverviewModel{
		Items:      news,
		Count:      count,
		Offset:     offset,
		NextOffset: offset + count,
		PrevOffset: offset - count,
		Page:       offset/count + 1,
		TotalPages: totalPages,
		TotalCount: totalCount,
	}, w)
}

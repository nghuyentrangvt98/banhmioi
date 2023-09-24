package repo

import (
	"context"
	"log/slog"

	"github.com/nghuyentrangvt98/banhmioi/config"
	"github.com/nghuyentrangvt98/banhmioi/ent"
	"github.com/nghuyentrangvt98/banhmioi/ent/news"
	"github.com/nghuyentrangvt98/banhmioi/model"
)

func GetNews(ctx context.Context, id int) *ent.News {
	news, _ := config.Client.News.Get(ctx, id)
	return news
}

func GetMultiNews(ctx context.Context, newsQuery model.NewsQuery) model.NewsList {
	query := config.Client.News.Query()
	if newsQuery.Category != "" {
		query = query.Where(news.Category(newsQuery.Category))
	}
	if newsQuery.Keyword != "" {
		query = query.Where(news.Or(
			news.TitleContains(newsQuery.Keyword),
			news.ContentContains(newsQuery.Keyword),
		))
	}
	query = query.Order(news.OrderOption(news.ByID()))
	news, err := query.All(ctx)
	if err != nil {
		slog.Info(err.Error())
	}
	total := len(news)
	return model.NewsList{Data: news, Total: total}
}

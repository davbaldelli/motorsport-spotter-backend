package entities

import (
	"motorsportspotter.backend/models"
	"time"
)

type NewsDb struct {
	Id       int
	Title    string
	Subtitle string
	Author   string
	Content  string
	Date     time.Time
	Image    string
}

type DbNewsList []NewsDb

func (NewsDb) TableName() string {
	return "news"
}

func (n NewsDb) ToEntity() models.News {
	return models.News{
		Id:       n.Id,
		Title:    n.Title,
		Subtitle: n.Subtitle,
		Author:   n.Author,
		Content:  n.Content,
		Date:     n.Date.Format("2006-01-02"),
		Image:    n.Image,
	}
}

func (l DbNewsList) ConvertAll() []models.News {
	var news []models.News
	for _, newsDb := range l {
		news = append(news, newsDb.ToEntity())
	}
	return news
}

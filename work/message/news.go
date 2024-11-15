package message

//News 图文消息
type News struct {
	CommonToken  `json:"-"`
	ArticleCount int        `xml:"ArticleCount" json:"-"`
	Articles     []*Article `xml:"Articles>item,omitempty" json:"articles"`
}

//NewNews 初始化图文消息
func NewNews(articles []*Article) *News {
	news := new(News)
	news.ArticleCount = len(articles)
	news.Articles = articles
	return news
}

//Article 单篇文章
type Article struct {
	Title       string `xml:"Title,omitempty" json:"title"`
	Description string `xml:"Description,omitempty" json:"description"`
	PicURL      string `xml:"PicUrl,omitempty" json:"picurl"`
	URL         string `xml:"Url,omitempty" json:"url"`
	Appid       string `xml:"-" json:"appid"`    //仅在发送应用消息时需要
	Pagepath    string `xml:"-" json:"pagepath"` //仅在发送应用消息时需要
}

//MpNews 图文消息
type MpNews struct {
	Articles []*MpNewsArticle `xml:"-" json:"articles"`
}

//MpNewsArticle mpnews类型的图文消息，跟普通的图文消息一致，唯一的差异是图文内容存储在企业微信
type MpNewsArticle struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

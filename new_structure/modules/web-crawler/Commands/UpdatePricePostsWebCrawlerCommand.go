package Commands

import (
	"fmt"
	"project-root/modules/web-crawler/Lib/divar"
	"project-root/sys-modules/console/Lib"
)

type UpdatePricePostsWebCrawlerCommand struct{}

func (c *UpdatePricePostsWebCrawlerCommand) Signature() string {
	return "web-crawler:update-price-posts"
}

func (c *UpdatePricePostsWebCrawlerCommand) Description() string {
	return "It checks all the posts and if the price changes, it saves it in priceHistory with the date"
}

func (c *UpdatePricePostsWebCrawlerCommand) Handle(args []string) {
	fmt.Println("crawler update-price-posts start...")
	divar.UpdatePriceCrawler()
}

var _ Lib.CommandInterface = &UpdatePricePostsWebCrawlerCommand{}

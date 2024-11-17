package Commands

import (
	"fmt"
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/database/Facades"
)

type GetPostsWebCrawlerCommand struct{}

func (c *GetPostsWebCrawlerCommand) Signature() string {
	return "web-crawler:get-posts"
}

func (c *GetPostsWebCrawlerCommand) Description() string {
	return "It collects the list of new posts from all three types of posts on different reference sites and creates them in the posts table"
}

func (c *GetPostsWebCrawlerCommand) Handle(args []string) {
	db := Facades.Db()
	//db.
	fmt.Println("Executing MyCommand with Param1:", db)

}

var _ Lib.CommandInterface = &GetPostsWebCrawlerCommand{}

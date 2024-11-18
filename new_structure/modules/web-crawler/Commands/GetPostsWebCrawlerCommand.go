package Commands

import (
	"fmt"
	"project-root/modules/web-crawler/Lib/divar"
	"project-root/sys-modules/console/Lib"
)

type GetPostsWebCrawlerCommand struct{}

func (c *GetPostsWebCrawlerCommand) Signature() string {
	return "web-crawler:get-posts"
}

func (c *GetPostsWebCrawlerCommand) Description() string {
	return "It collects the list of new posts from all three types of posts on different reference sites and creates them in the posts table"
}

func (c *GetPostsWebCrawlerCommand) Handle(args []string) {
	fmt.Println("crawler start...")
	divar.RunCrawlerPosts()
}

var _ Lib.CommandInterface = &GetPostsWebCrawlerCommand{}

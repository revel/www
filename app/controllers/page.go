package controllers

import (
	"fmt"
	//"os/exec"
	//"bufio"
	//"//io/ioutil"
	"html/template"
	//"path/filepath"
	//"strings"

	"github.com/revel/revel"
	//"github.com/russross/blackfriday"
	//"gopkg.in/yaml.v2"
	//"github.com/pksunkara/pygments"
	"github.com/revel/revelframework.com/app/meta"
)



type Page struct {
	*revel.Controller
}

// /robots.txt - Only allow spiders on prod site
func (c Page) RobotsTxt() revel.Result {

	txt := "User-agent: *\n"
	if revel.Config.BoolDefault("site.live", false)  == false {
		txt += "Disallow: /\n"
	}
	txt += "\n"

	return c.RenderText(txt)
}

// main home page
func (c Page) Index() revel.Result {
	return c.Render()
}
func (c Page) Debug(section, page string) revel.Result {
	return c.RenderJson(meta.Site)
}



type CurrPage struct {
	//Title string
	Version string
	Section string
	SectionTitle string
	Page string
	Title string
	Path string
	Lang string
	ContentSS template.HTML
	Content string
}

//var Site *SiteStruct

//func GetCurrPage(section, section_title, version, lang, page string) CurrPage {
//
//	s := CurrPage{Section: section,  PageUrl: page, Version: version, Lang: lang}
//	return s
//}


// render an expected markdown file
func (c Page) Page(section, page string) revel.Result {
	fmt.Println("PAGE==", section, page)

	//cp := CurrPage{Section: section, Page: page, Title: "Title", Path: "PATHH", Content: "CYIPEeeeeeeeeeeeeONTENT"}
	c.RenderArgs["Site"] = meta.Site
	c.RenderArgs["Section"] = meta.Site.Sections[section]


	//c.RenderArgs["Page"] = cp


	pdata := meta.ReadMarkDownPage(section, page)
	c.RenderArgs["Page"] = pdata
	if pdata.Error != nil {
		fmt.Println("error==", pdata.Error)
	}
	fmt.Println("error==", pdata.Title)
	/*
	c.RenderArgs["cPage"] = cPage


	cPage := GetCurrPage(site_section, "Manual", ver, lang, page)

	nav := GetNav(site_section)
	c.RenderArgs["nav"] = nav


	page_no_ext := page
	if filepath.Ext(page) == ".html" { // wtf includes the .
		page_no_ext = page[0: len(page) - 5]
	}

	// use template.HTML to "unescape" encoding.. ie proper html not &lt;escaped
	pdata := ReadMarkdownPage(site_section, page_no_ext)
	c.RenderArgs["page_content"] = pdata.HTML
	cPage.PageTitle = pdata.Title


	c.RenderArgs["cPage"] = cPage
	*/
	return c.Render()

}

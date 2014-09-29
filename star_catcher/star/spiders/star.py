# -*- coding: utf-8 -*-
import scrapy
from scrapy.contrib.spiders import CrawlSpider, Rule
from scrapy.contrib.linkextractors import LinkExtractor

class StarSpider(CrawlSpider):
	name = "star"
	allowed_domains = ["yanjiuseng.net"]

	start_urls = [
		"http://www.yanjiuseng.net/fanhao/986.html"
	]

	rules = (
		Rule(LinkExtractor(allow=('(.*)\.html')), follow=True, callback='parse_item'),
	)

	def parse_item(self, response):
		sels = response.xpath('//div[@class="article_img_mid"]/h1/text()')
		if len(sels) > 0:
			print "catch url: %s" % response.url
			star_name = sels[0].extract()
			print star_name

		yield None
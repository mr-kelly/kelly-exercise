# -*- coding: utf-8 -*-
import scrapy
from scrapy.contrib.spiders import CrawlSpider, Rule
from scrapy.contrib.linkextractors import LinkExtractor
import re, sys
sys.path.append("..")
import database_client

class StarSpider(CrawlSpider):
	name = "star"
	allowed_domains = ["yanjiuseng.net"]

	start_urls = [
		"http://www.yanjiuseng.net/fanhao/986.html"
	]

	rules = (
		Rule(LinkExtractor(allow=('(.*)/fanhao/(.*)\.html')), follow=True, callback='parse_item'),
	)

	def parse_item(self, response):
		title_sels = response.xpath('//div[@class="article_img_mid"]')
		
		if len(title_sels) > 0:
			main_title_sel = title_sels[0]

			print "catch url: %s" % response.url
			star_id = re.match('.*/(\d+).html', response.url).group(1) # Get Id
			star_name = main_title_sel.xpath('h1/text()')[0].extract()
			print "%s---%s" % (star_id, star_name)

			database_client.db_client.replace_into(
				'star', 
				'id, name', 
				'%s, "%s"' % (star_id, star_name))

		yield None
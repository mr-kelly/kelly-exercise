# -*- coding: utf-8 -*-
import scrapy
from scrapy.contrib.spiders import CrawlSpider, Rule
from scrapy.contrib.linkextractors import LinkExtractor
import re, sys
sys.path.append("..")
from database_client import *

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

		# 标题
		title_sels = response.xpath('//div[@class="article_img_mid"]')
		
		if len(title_sels) > 0:
			main_title_sel = title_sels[0]

			print "catch url: %s" % response.url
			star_id = re.match('.*/(\d+).html', response.url).group(1) # Get Id
			star_name = main_title_sel.xpath('h1/text()')[0].extract()
			print "%s---%s" % (star_id, star_name)

			db_client.replace_into(
				'star', 
				'id, name', 
				'%s, "%s"' % (star_id, star_name))


		# 小图片地址
		pic_url = response.xpath('//div[@class="article_img_left"]/img/@src').extract()
		print u'小图片地址: %s' % pic_url

		# 大图片地址
		big_pic_url = response.xpath('//div[@class="arcBody"]/div[@id="disappear"]//img/@src').extract()
		print u'大图地址: %s' %big_pic_url

		# 文字内容
		content = response.xpath('//div[@class="arcBody"]/div[@id="disappear"]').extract()

		# 表格开始
		fh_table_rows = response.xpath('//div[@class="arcBody"]/table/tbody/tr')
		if (len(fh_table_rows) > 0):
			for row in fh_table_rows: # tr
				cols = row.xpath('td/text()') # td
				assert(len(cols) == 5)
				for col in cols:
					print 'Col-Value: %s' % col.extract().strip()
				# response.xpath('//table[@class="fanhao_list_table"]/tbody/tr')[0].extract()



		

		yield None
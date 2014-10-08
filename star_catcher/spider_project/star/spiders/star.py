# -*- coding: utf-8 -*-
import scrapy
from scrapy.contrib.spiders import CrawlSpider, Rule
from scrapy.contrib.linkextractors import LinkExtractor
import re
from database_client import *
import os


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

			# 小图片地址

			small_pic_url = ""
			small_pic_xpath = response.xpath('//div[@class="article_img_left"]/img/@src')
			if len(small_pic_xpath) > 0:
				small_pic_url = "http://%s%s" % (self.allowed_domains[0], small_pic_xpath[0].extract())
			print u'小图片地址: %s' % small_pic_url

			# 大图片地址
			big_pic_url = ""
			big_pic_xpath = response.xpath('//div[@class="arcBody"]/div[@id="disappear"]//img/@src')
			if len(big_pic_xpath) > 0:
				big_pic_url = "http://%s%s" % (self.allowed_domains[0], big_pic_xpath[0].extract())
			print u'大图地址: %s' % big_pic_url

			# 文字内容
			# content = response.xpath('//div[@class="arcBody"]/div[@id="disappear"]')[0].extract()
			# print content

			# 表格开始
			fh_table_rows = response.xpath('//div[@class="arcBody"]/table/tbody/tr')
			if (len(fh_table_rows) > 0):
				for row in fh_table_rows: # tr
					cols = row.xpath('td/text()') # td
					if len(cols) == 5:
						# print u'番号列表表格匹配到'
						fanhao = cols[0].extract().strip()
						av_name = cols[1].extract().strip()
						av_time = cols[2].extract().strip()
						av_public = cols[3].extract().strip()
						av_other = cols[4].extract().strip()
						print u'%s %s %s %s %s' % (fanhao, av_name, av_time, av_public, av_other)
						db_client.replace_into(
							'movie', 
							'id, star_id, name, movie_length, public_time, other', 
							'"%s", %s, "%s", "%s", "%s", "%s"' % (fanhao, star_id, av_name, av_time, av_public, av_other))		
					else:
						print u'表格不是5列，而是%d列, 非番号列表吧' % len(cols)
						

			db_client.replace_into(
				'star', 
				'id, name, small_pic_url, big_pic_url', 
				'%s, "%s", "%s", "%s"' % (star_id, star_name, small_pic_url, big_pic_url))

			

		yield None
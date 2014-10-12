# -*- coding: utf-8 -*-
import sys    
reload(sys)   
sys.setdefaultencoding('utf8')   

import scrapy
from scrapy.contrib.spiders import CrawlSpider, Rule
from scrapy.contrib.linkextractors import LinkExtractor
import re
# from database_client import *
import os

from sqlobject import *
import db_object

class KTorrent(SQLObject):
	class sqlmeta:
		changeSchema = True
	# 	lazyUpdate = True
        # cacheValues = False
	name = UnicodeCol(default=None)
	size = UnicodeCol(default=None)
	time = UnicodeCol(default=None)
	intro = UnicodeCol(default=None)

KTorrent.createTable(ifNotExists=True)

# class w5u6d_torrent(CrawlSpider):
class w5u6d_torrent(scrapy.Spider):
	name = "w5u6d_torrent"
	allowed_domains = ["www.5u6d.com"]

	start_urls = [
		"http://www.5u6d.com/bt_see.asp?id=34790"
	]

	rules = (
	 	Rule(LinkExtractor(allow=('http://www.5u6d.com/bt_see.asp?id=(.*)')), follow=True, callback='parse_item'),
	)

	def parse(self, response):
		print "catch url: %s" % response.url
		torrent_id = re.match('.*id=(\d+)', response.url).group(1) # Get Id
		torrent_id = int(torrent_id)
		print "The ID: %s" % torrent_id

		torrent_name =  response.css('.blkNav1 h1::text')[0].extract()
		print torrent_name

		# 分類: bt種子-> "亞洲有碼""
		cat_name = response.css('.blkNav1-b span a::text')[2].extract()
		print cat_name

		# 文章 TODO: 內含大小
		response.css('table.blkMovCon tr td')[0].re(u'文件大小：(.*) \| 时间：(.*) <br>')
		content = response.css('table.blkMovCon tr td')[0].extract()

		matchall =  re.findall(u'文件大小：(.*) \| 时间：(.*) <br>([\s\S]*)</td>', content, re.M)
		file_size = matchall[0][0].strip()
		file_time = matchall[0][1].strip()
		file_intro = matchall[0][2]


		try:
			dbObj = KTorrent(id=torrent_id)
		except:
			dbObj = KTorrent.get(torrent_id)

		dbObj.set(name=torrent_name, size=file_size, time=file_time, intro=file_intro)

		print u"文件大小: %s" % file_size
		print u"文件時間: %s" % file_time
		print u'文件描述: %s' % file_intro





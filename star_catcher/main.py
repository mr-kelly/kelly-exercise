import scrapy
import sqlite3

class DatabaseClient:
	def __init__(self, path):
		self.db = sqlite3.connect(path)
		self.client = self.db.cursor() # main cursor
	def execute(self, *args):
		return self.client.execute(*args)

db_client = DatabaseClient("./data.db")
db_client.execute("create table catalog (id integer primary key,pid integer,name varchar(10) UNIQUE,nickname text NULL)")

class StarItem(scrapy.Item):
	name = scrapy.Field()


class StarSpider(scrapy.Spider):
	name = "star"
	allowed_domains = ["www.yanjiuseng.net"]

	start_urls = [
		"http://www.yanjiuseng.net/"
	]

	def parse(self, response):
		print response.xpath('/div[@class="article_img_mid"]/h1/text()').extract()

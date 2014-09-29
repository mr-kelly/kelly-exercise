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

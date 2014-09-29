import scrapy
import sqlite3



class DatabaseClient:
	def __init__(self, path):
		self.db = sqlite3.connect(path)
		self.client = self.db.cursor() # main cursor
	def execute(self, *args):
		return self.client.execute(*args)

	def replace_into(self, table, cols, values):
		self.client.execute("REPLACE into `%s` (%s) VALUES (%s);" % (table, cols, values))
		self.db.commit()


db_client = DatabaseClient("./data.db")
#coding=utf-8
import scrapy
import sqlite3

class DatabaseClient:
	def __init__(self, path):
		self.db = sqlite3.connect(path)
		self.client = self.db.cursor() # main cursor
		self.client.execute("""
		CREATE TABLE IF NOT EXISTS "star" (
		"id"  integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		"name"  TEXT,
		"small_pic_url"  TEXT,
		"big_pic_url"  TEXT,
		"content"  TEXT
		) ;

		""")
		self.client.execute("""
		CREATE TABLE IF NOT EXISTS "movie" (
		"id"  TEXT NOT NULL,
		"star_id"  integer,
		"name"  TEXT,
		"movie_length"  TEXT,
		"public_time"  TEXT,
		"other"  TEXT,
		PRIMARY KEY ("id" ASC)
		);
		""")

	def execute(self, *args):
		return self.client.execute(*args)

	def replace_into(self, table, cols, values):
		sql = "REPLACE into `%s` (%s) VALUES (%s);" % (table, cols, values)
		print 'Execute SQL: %s' % sql
		self.client.execute(sql)
		self.db.commit()


db_client = DatabaseClient("./data.db")
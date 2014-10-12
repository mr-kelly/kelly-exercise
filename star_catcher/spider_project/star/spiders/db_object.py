from sqlobject import *
import sys, os

db_filename = os.path.abspath('./data2.db')
connection_string = 'sqlite:' + db_filename
connection = connectionForURI(connection_string)
sqlhub.processConnection = connection

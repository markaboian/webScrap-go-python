import pandas as pd
import requests
from io import StringIO
from sqlalchemy import create_engine, Table, Column, Integer, String, MetaData, DECIMAL
import mysql.connector

url = 'https://fbref.com/en/comps/9/Premier-League-Stats'
response = requests.get(url)
html_content = StringIO(response.text)
df = pd.read_html(html_content, attrs={'id': 'results2023-202491_overall'})[0]
df.to_csv('/Users/markaboian/Desktop/Go-Python/premier-league-table.csv')


connection = mysql.connector.connect(user='root', password='mceastermako2002', host='localhost', port=3306)
cursor = connection.cursor()
cursor.execute('CREATE DATABASE IF NOT EXISTS premier_league')
cursor.close()
connection.close()


engine = create_engine('mysql+pymysql://root:mceastermako2002@localhost:3306/premier_league')
metadata = MetaData()

table = Table('regular_season', metadata,
              Column('rk', Integer, primary_key=True),
              Column('Squad', String(100)),
              Column('Matches_Played', Integer),
              Column('W', Integer),
              Column('D', Integer),
              Column('L', Integer),
              Column('GF', Integer),
              Column('GA', Integer),
              Column('GD', Integer),
              Column('Pts', Integer),
              Column('Pts/MP', DECIMAL(10, 2)),
              Column('xG', DECIMAL(10, 2)),
              Column('xGA', DECIMAL(10, 2)),
              Column('xGD', DECIMAL(10, 2)),
              Column('xGD/90', DECIMAL(10, 2)),
              Column('Last_Five', String(100)),
              Column('Attendance', Integer),
              Column('Top_Team_Scorer', String(100)),
              Column('Goalkeeper', String(100)),
              Column('Notes', String(100))
)

metadata.create_all(engine)

df = pd.read_csv('premier-league-table.csv')
df.to_sql('regular_season', engine, if_exists='replace', index=False)

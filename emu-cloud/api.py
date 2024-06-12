#!/usr/bin/env python3

import MySQLdb as mysql
import os
import json
from fastapi import FastAPI

app = FastAPI()

def get_config():
    file = ".config.json"
    config_path = os.path.abspath(f'{file}')
    n = open(f'{config_path}')
    data = json.load(n)
    db_host = data['db_host']
    db_port = data['db_port']
    db_user = data['db_user']
    db_password = data['db_password']
    debug = data['debug']
    return db_host, db_port, db_user, db_password, debug



#user = config[0]
#password = config[1]
#vsphere_user = config[2]
#vsphere_password = config[3]
#debug = bool(config[4])

#def mysql_connection(db_host,db_port,db_user,db_password,database):
def mysql_connection():
   config = get_config()
   mydb = mysql.connect(host=config[0],port=config[1],user=,password=f"{db_password}",database=f"{database}")
   return mydb


#@app.get("/api/hosts")
#async def hosts():
#    cursor = mysql_connection()  
#    cursor.execute("""SELECT"""",()
#    return {"hosts": f"{hosts}"}

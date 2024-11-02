#!/usr/bin/env python3

import logging
import sys
import os
from datetime import datetime

try:
    #from app import app
    from flask import Flask, render_template
    from dotenv import load_dotenv, dotenv_values
    from starlette.responses import RedirectResponse
    from sqlalchemy.orm import Session
    import models
    from database import SessionLocal, engine
except ImportError:
    print("fastapi, dotenv, sqlalchemy module missing")
    sys.exit(-1)


models.Base.metadata.create_all(bind=engine)

app = FastAPI()

if os.path.exists(".env"):
    config = dotenv_values(".env")
else:
    config = dotenv_values("/etc/default/app.sh")

debug = config['API_ENVIRONMENT']
api_database = config['API_DATABASE']

def create_logs(*command):
    if debug == "development":
        logging.basicConfig(format='%(asctime)s :: %(levelname)s :: %(message)s', encoding='utf-8',
                                datefmt='%Y-%m-%d %H:%M:%S', level=logging.DEBUG)
        logging.debug(f"{command}")
    if debug == "production":
        logging.basicConfig(format='%(asctime)s :: %(levelname)s :: %(message)s', encoding='utf-8',
                                datefmt='%Y-%m-%d %H:%M:%S', level=logging.INFO)
        logging.info(f"{command}")

# Dependency
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

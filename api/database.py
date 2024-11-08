#!/usr/bin/env python3

from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from dotenv import load_dotenv, dotenv_values
import os

if os.path.exists(".env"):
    config = dotenv_values(".env")
else:
    config = dotenv_values("/etc/default/app.sh")


SQLALCHEMY_DATABASE_URL = "sqlite:///./db.sqlite"
# SQLALCHEMY_DATABASE_URL = "postgresql://user:password@postgresserver/db"

engine = create_engine(
    SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False}
)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base = declarative_base()

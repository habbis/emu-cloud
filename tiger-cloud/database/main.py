#!/usr/bin/env python3

import sqlite3


def database(name):
   connection = sqlite3.connect(f"data/{name}.db")
   return connection




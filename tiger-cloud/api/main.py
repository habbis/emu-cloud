import sqlite3
from fastapi import FastAPI

app = FastAPI()

@app.get("/api/hosts")
async def hosts():
    return {"hosts": f"{hosts}"}

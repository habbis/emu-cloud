#!/usr/bin/env python3

from sqlalchemy import Boolean, Column, Integer, String

from database import Base


class host2hostgroup(Base):
    __tablename__ = "host2hostgroup"
    id = Column(Integer, primary_key=True, index=True)
    id_hosts = Column(Integer, index=True)
    id_hostgroup = Column(Integer, index=True)
    time_created = Column(DateTime(timezone=True), server_default=func.now())
    time_updated = Column(DateTime(timezone=True), onupdate=func.now())

class host_vars(Base):
    __tablename__ = "host_vars"
    id = Column(Integer, primary_key=True, index=True)
    id_hosts = Column(Integer, index=True)
    keyname = Column(String, index=True)
    value = Column(Text, index=True)
    time_created = Column(DateTime(timezone=True), server_default=func.now())
    time_updated = Column(DateTime(timezone=True), onupdate=func.now())

class hosts(Base):
    __tablename__ = "hosts"
    id = Column(Integer, primary_key=True, index=True)
    fqdn = Column(String, index=True)
    env = Column(String, index=True)
    time_created = Column(DateTime(timezone=True), server_default=func.now())
    time_updated = Column(DateTime(timezone=True), onupdate=func.now())

class interface(Base):
    __tablename__ = "interface"
    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, index=True)
    ip = Column(String, index=True)
    id_hosts = Column(Integer, index=True)
    id_vlan = Column(Integer, index=True)
    time_created = Column(DateTime(timezone=True), server_default=func.now())
    time_updated = Column(DateTime(timezone=True), onupdate=func.now())

# examole vlan name
# int_22_test_dc1
class interface(Base):
    __tablename__ = "interface"
    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, index=True)
    vlan_id = Column(Integer, index=True)
    prefix = Column(String, index=True)
    comment = Column(String, index=True)
    time_created = Column(DateTime(timezone=True), server_default=func.now())
    time_updated = Column(DateTime(timezone=True), onupdate=func.now())


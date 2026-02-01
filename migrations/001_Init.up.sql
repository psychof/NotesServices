CREATE TABLE IF NOT EXISTS Notes (
id serial primary key,
title text not null,
description text,
time_stamp timestamp without time zone
)


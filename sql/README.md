## Dump data 

```

pg_dump -U postgres -h localhost -p 5432 --data-only simple_podcast_db > data.sql

```


## Import data


```

psql simple_podcast_db < data.sql

```

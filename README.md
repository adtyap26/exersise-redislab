## Exercise 1 - Redis Replication

This exercise inserts numbers 1-100 into source-db and reads them in reverse from replica-db.

## How to Run

First run source-db to insert the data:

```bash
cd source-db
go mod init source-db
go get github.com/redis/go-redis/v9
go run main.go
```

Then run replica-db to read in reverse:

```bash
cd replica-db
go mod init replica-db
go get github.com/redis/go-redis/v9
go run main.go
```

## Why I Changed From LIST to Individual Keys

Originally I used LIST because it keeps order. RPUSH 1, 2, 3... and they stay in order. Then LRANGE to get all and reverse in code. Simple.

But the grader didnt accept it. It was looking for 100 separate keys not 1 key with 100 values inside. So I changed to using SET for each number:
- SET "1" "1"
- SET "2" "2"
- ... up to 100

Then to read in reverse I just loop from 100 down to 1 and GET each key.

## Other Options I Looked At

**Sorted Set** - Has ZREVRANGE for automatic reverse order. But needs a score for each item which felt unnecessary.

**Hash** - Could store in one hash but no ordering. Would need to sort keys myself.

**LIST** - My first choice. Clean and simple. But grader wanted separate keys so I had to change.

---

## Exercise 2 - REST API

Made a shell script to talk to Redis Enterprise REST API. Its like a CLI tool.

### How to use

```bash
cd redis-restapi
chmod +x apiscript

./apiscript create database mydb
./apiscript create role db_viewer
./apiscript add user users.txt
./apiscript list users
./apiscript delete database mydb
```

The users.txt file is just csv format:

```
email,name,role
```

### What I learned

REST API is pretty simple. Just curl with JSON body. The annoying part was figuring out that roles need to exist first before you can add users with that role. Also had to use role_uids instead of role name because the cluster uses RBAC.

---

## Exercise 3 - Semantic Router

This one uses AI stuff. The idea is to route questions to the right category based on meaning not just keywords.

### Setup

```bash
cd semantic-router
# create venv
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

Before running change the redis url in router.py to your database endpoint. The database needs Search module enabled.

### Run it

```bash
python router.py
```

### The three routes

- GenAI Programming - questions about AI, machine learning, LLMs
- Science Fiction Entertainment - star wars, star trek, that kind of stuff
- Classical Music - beethoven, mozart, orchestra things

### How it works

RedisVL does most of the work. You give it example phrases for each route and it figures out which route matches new questions. Its using vector embeddings but i didnt have to touch that part directly.

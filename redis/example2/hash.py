import redis

redisClient = redis.StrictRedis(host='localhost', port=6379, db=0)

redisClient.hset("NumberVsString", "1", "One")

redisClient.hset("NumberVsString", "2", "Two")

redisClient.hset("NumberVsString", "3", "Three")

print("Value for the key 3 is")

print(redisClient.hget("NumberVsString", "3"))

-- Restricted object
local key = KEYS[1]
-- Window size
local window = tonumber(ARGV[1])
-- Rate
local threshold = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
-- window start time
local min = now - window

redis.call('ZREMRANGEBYSCORE', key, '-inf', min)
local cnt = redis.call('ZCOUNT', key, '-inf', '+inf')

if cnt > threshold then
    -- do rate limit
    return "true"
else
    redis.call('ZADD', key, now, now)
    redis.call('PEXPIRE', key, window)
    return "false"
end
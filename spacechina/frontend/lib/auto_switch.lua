-- 这个脚本演示当请求大于count/second以后，切换emergency_ups
local default_ups = "backend_default"
local emergency_ups = "backend1"
local limit_count = require "resty.limit.count"
local count = 100
local second = 1
-- rate count/second
local lim, err = limit_count.new("limit_conn_store", count, second)
-- 根据并发请求数进行判断设置upstream
if not lim then
    ngx.log(ngx.ERR, "failed to instantiate a resty.limit.count object: ", err)
    return default_ups
end

-- use host as the limiting key
local key = ngx.var.http_host
local delay, err = lim:incoming(key, true)
if not delay then
    if err == "rejected" then
        ngx.header["X-RateLimit-Limit"] = count.."/"..second
        ngx.header["X-RateLimit-Remaining"] = 0
        ngx.header["X-RateLimit-Upstream"] = emergency_ups
        return emergency_ups
    end
    ngx.log(ngx.ERR, "failed to limit count: ", err)
    return default_ups
end

-- the 2nd return value holds the current remaining number
-- of requests for the specified key.
local remaining = err

ngx.header["X-RateLimit-Limit"] = count.."/"..second
ngx.header["X-RateLimit-Remaining"] = remaining
ngx.header["X-RateLimit-Upstream"] = default_ups
return default_ups
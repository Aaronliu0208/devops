local limit_count = require "resty.limit.count"
local rate_pre_second = 100
-- rate: 100 requests per 1s
local lim, err = limit_count.new("limit_conn_store", rate_pre_second, 1)
if not lim then
    ngx.log(ngx.ERR, "failed to instantiate a resty.limit.count object: ", err)
    return ngx.exit(500)
end
-- use the Authorization header as the limiting key
local key = ngx.req.get_headers()["Authorization"] or "public"
local delay, err = lim:incoming(key, true)
if not delay then
    if err == "rejected" then
        ngx.header["X-RateLimit-Limit"] = tostring(rate_pre_second)
        ngx.header["X-RateLimit-Remaining"] = 0
        --每秒请求超过100次的就报错403
        --return ngx.exit(403)
        -- 一旦负载高于每秒100个请求就重定向到百度
        return ngx.redirect("http://www.baidu.com", ngx.HTTP_MOVED_TEMPORARILY)
    end
    ngx.log(ngx.ERR, "failed to limit count: ", err)

    return ngx.exit(500)
end

    -- the 2nd return value holds the current remaining number
-- of requests for the specified key.
local remaining = err

ngx.header["X-RateLimit-Limit"] = tostring(rate_pre_second)
ngx.header["X-RateLimit-Remaining"] = remaining
local map = {}
map[ "fcydswjcz"]="https://video.ht2025.com/sv/2f8a2349-172a1cf88fa/2f8a2349-172a1cf88fa.mp4"
local url=ngx.var.request_uri
for key, value in pairs(map) do
    if string.match(url, key) then
        return ngx.redirect(value, ngx.HTTP_MOVED_TEMPORARILY)
    end
end
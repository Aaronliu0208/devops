local map = {}
map["fcydswjcz"]="https://video.ht2025.com/sv/2f8a2349-172a1cf88fa/2f8a2349-172a1cf88fa.mp4"
map["fyfcyy"]="https://video.ht2025.com/sv/30b68a93-17283a5629f/30b68a93-17283a5629f.mp4"
map["glyy"]="https://video.ht2025.com/sv/50d8cd3e-17283a5631f/50d8cd3e-17283a5631f.mp4"
map["tngy1"]="https://video.ht2025.com/sv/1a1c6e4b-17283a56323/1a1c6e4b-17283a56323.mp4"
map["tngy2"]="https://video.ht2025.com/sv/5652c76f-17283a56323/5652c76f-17283a56323.mp4"
map["/upvideo/1.mp4"]="https://video.ht2025.com/sv/b70e941-172920a0784/b70e941-172920a0784.mp4"
map["/upvideo/2.mp4"]="https://video.ht2025.com/sv/14835758-172920a07ba/14835758-172920a07ba.mp4"
map["/upvideo/3.mp4"]="https://video.ht2025.com/sv/6bbb9a3-172920a0818/6bbb9a3-172920a0818.mp4"
map["/upvideo/4.mp4"]="https://video.ht2025.com/sv/5191be61-172920a082b/5191be61-172920a082b.mp4"
map["/upvideo/5.mp4"]="https://video.ht2025.com/sv/159dd7d9-172a1cf8844/159dd7d9-172a1cf8844.mp4"
map["/upvideo/6.mp4"]="https://video.ht2025.com/sv/2b335323-172a1cf8881/2b335323-172a1cf8881.mp4"
map["/upvideo/7.mp4"]="https://video.ht2025.com/sv/79fb4e5-172a1cf88bb/79fb4e5-172a1cf88bb.mp4"
map["/upvideo/8.mp4"]="https://video.ht2025.com/sv/2f8a2349-172a1cf88fa/2f8a2349-172a1cf88fa.mp4"
map["/upvideo/9.mp4"]="https://video.ht2025.com/sv/95a6125-172a1cf892d/95a6125-172a1cf892d.mp4"
map["/upvideo/10.mp4"]="https://video.ht2025.com/sv/2e499dc6-172a665d72b/2e499dc6-172a665d72b.mp4"
map["/upvideo/11.mp4"]="https://video.ht2025.com/sv/4a04f800-172a79fc541/4a04f800-172a79fc541.mp4"
map["/upvideo/12.mp4"]="https://video.ht2025.com/sv/5e9ba288-172a96a25fe/5e9ba288-172a96a25fe.mp4"
map["/upvideo/zlt.mp4"]="https://video.ht2025.com/f440b055ab0b4111834a1245308df827/4cdc60203c0d69a393611ce281785475-sd-nbv1.mp4"
local url=ngx.var.request_uri
for key, value in pairs(map) do
    if string.match(url, key) then
        return ngx.redirect(value, ngx.HTTP_MOVED_TEMPORARILY)
    end
end
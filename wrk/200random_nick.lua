wrk.headers["Content-Type"]="application/x-www-form-urlencoded"
wrk.headers["Cookie"]="token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjIxMjMyZjI5N2E1N2E1YTc0Mzg5NGEwZTRhODAxZmMzIiwicGFzc3dvcmQiOiI4MWRjOWJkYjUyZDA0ZGMyMDAzNmRiZDgzMTNlZDA1NSIsImV4cCI6MTY2MDY0MDMyMSwiaXNzIjoieWFuZ21pbmdydW4ifQ.4RsyuWb7Q2rKfFRAOaLfZCx4bXwp5zNY0_14o-2BgJE"
local UserCount = 10000000
request = function()
    local count = math.random(1,UserCount)
    body = "nickname=yyy"
    path = "/api/user/nick?username="..tostring(count)
    return wrk.format("POST",path,nil,body)
end

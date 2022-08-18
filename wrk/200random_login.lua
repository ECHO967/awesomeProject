wrk.headers["Content-Type"]="application/x-www-form-urlencoded"
local UserCount = 10000000
request = function()
    local count = math.random(1,UserCount)
    body = "username="..tostring(count).."&&password="..tostring(count)
    path="/api/user/login"
    return wrk.format("POST",path,nil,body)
end

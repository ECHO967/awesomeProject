
wrk.headers["Content-Type"]="application/x-www-form-urlencoded"

local UserCount = 1
request = function()
    body = "username="..tostring(UserCount).."&&password="..tostring(UserCount)
    path="/api/user/login"
    if UserCount >= 200 then
        UserCount = 0
    else
        UserCount= UserCount + 1
    end
    return wrk.format("POST",path,nil,body)
end

-- init random
math.randomseed(os.time())

-- the request function that will run at each request
request = function()
    -- define the path that will search for q=%v 9%v being a random number between 0 and 1000)
    phoneId = math.random(1, 1000000)
    userId = math.random(1, 1000000)
    body = '{"userId": ' .. userId .. ', "phoneId":' .. phoneId .. ','
    body = body .. '"categoryId": 9, "itemId": 1337, "phoneDisplayLoc": "monolith-item"}'
    --print("body: " .. body)
    -- if we want to print the path generated
    --print(url_path)
    -- Return the request object with the current URL path
    return wrk.format(nil, nil, nil, body)
end

-- response = function(status, headers, body)
--    print("status: " .. status)
--    print("body: " .. body)
-- end

wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

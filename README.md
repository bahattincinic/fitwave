# FitWave

Virtualize Your Strava Workouts

## How to get access token

Follow the following document to get Access token

https://developers.strava.com/docs/getting-started/#:~:text=Access%20tokens%20are%20required%20for,tokens%20expire%20every%20six%20hours.

Step 1:
http://www.strava.com/oauth/authorize?client_id=30980&response_type=code&redirect_uri=http://localhost/exchange_token&approval_prompt=force&scope=read,activity:read

Step2:
curl -X POST https://www.strava.com/oauth/token \
-F client_id=30980 \
-F client_secret=d8d5256cc59c20e49d984d412377604e8dfb6052 \
-F code=e80825b36f668a777f4ce9e9960f6c2b6351dd6f \
-F grant_type=authorization_code

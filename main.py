import requests 
import json
import datetime
import click

def getCurrentDate():
    today = datetime.date.today()
    return today

def getAPIKey():
    return "B63dwvvoZp5Kh47nTVcvRg3Mf1SbkBvCkrAkFbSA"

def makeAPIString(date):
    apiCall = "https://www.rescuetime.com/anapi/data?key="
    apiCall += getAPIKey()
    apiCall += "&perspective=interval&restrict_kind=productivity&interval=hour&restrict_begin="
    apiCall += date
    apiCall += "&restrict_end="
    apiCall += date
    apiCall += "&format=csv"
    return apiCall
        
response = requests.get(makeAPIString(datetime.today()))



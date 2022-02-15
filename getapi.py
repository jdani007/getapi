import urllib.request

import json

urlData = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/2.5_day.geojson"

def printResults(data):
    theJSON = json.loads(data)
    if "title" in theJSON["metadata"]:
        print(theJSON["metadata"]["title"])
    print("--------------------\n")

    count = theJSON["metadata"]["count"]
    print(count, "events recorded")
    print("--------------------\n")

    for i in theJSON["features"]:
        print(i["properties"]["place"])
    print("--------------------\n")

    # for i in theJSON["features"]:
    #     if i["properties"]["mag"] >= 4.0:
    #         print(i["properties"]["place"])
    # print("--------------------\n")

    # for i in theJSON["features"]:
    #     felt = i["properties"]["felt"]
    #     if felt != None:
    #         if felt > 0:
    #             print(i["properties"]["place"], felt,"times")
    # print("--------------------\n")

webUrl = urllib.request.urlopen(urlData)
print ("result code:" + str(webUrl.getcode()))

if (webUrl.getcode() == 200):
    data = webUrl.read()
    printResults(data)
else:
    print("Recieved an error from the server, cannot print results", webUrl.getcode())

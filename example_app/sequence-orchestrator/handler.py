import os
import requests
import sys

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    # uses a default of "gateway" for when "gateway_hostname" is not set
    gateway_hostname = os.getenv("gateway_hostname", "gateway") 
    url = "http://" + gateway_hostname + ":8080/function/"

    req = call_function("appender-0", url + "appender-0", req)
    req = call_function("appender-1", url + "appender-1", req)
    req = call_function("appender-2", url + "appender-2", req)

    return req


def call_function(function_name, function_url, data):
    r = requests.get(function_url, data=data)
    if r.status_code != 200:
        sys.exit("Error with %s function, status code: %d\n" % (function_name, 200, req.status_code))

    return r.text


# Author of the file: yodalee
# link to the github repo: https://github.com/yodalee/HDEchallenge

import requests
import hmac
import hashlib
import time
import struct
import json
from requests.auth import HTTPBasicAuth


USERID = "EMAIL"  # Change the "EMAIL" to your email address.
ROOT = "https://api.challenge.hennge.com/challenges/003"
CONTENT_TYPE = "application/json"
SECRET_SUFFIX = "HENNGECHALLENGE003"
SHARED_SECRET = USERID + SECRET_SUFFIX

TIMESTEP = 30
T0 = 0
DIGITS = 10

data = {
    "github_url": "https://gist.github.com/YOUR_ACCOUNT/GIST_ID",
    "contact_email": USERID
}


def HOTP(K, C, digits=10):
    """HTOP:
    K is the shared key
    C is the counter value
    digits control the response length
    """
    K_bytes = str.encode(K)
    C_bytes = struct.pack(">Q", C)
    hmac_sha512 = hmac.new(key=K_bytes, msg=C_bytes,
                           digestmod=hashlib.sha512).hexdigest()
    return Truncate(hmac_sha512)[-digits:]


def Truncate(hmac_sha512):
    """truncate sha512 value"""
    offset = int(hmac_sha512[-1], 16)
    binary = int(hmac_sha512[(offset * 2):((offset*2)+8)], 16) & 0x7FFFFFFF
    return str(binary)


def TOTP(K, digits=10, timeref=0, timestep=30):
    """TOTP, time-based variant of HOTP
    digits control the response length
    the C in HOTP is replaced by ( (currentTime - timeref) / timestep )
    """
    C = int(time.time() - timeref) // timestep
    return HOTP(K, C, digits=digits)


passwd = TOTP(SHARED_SECRET, DIGITS, T0, TIMESTEP).zfill(10)
resp = requests.post(ROOT, auth=HTTPBasicAuth(
    USERID, passwd), data=json.dumps(data))
print(resp.json())

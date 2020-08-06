import requests

import sys


def main(num: int):
    with open("res/my_plugin.zip", "rb") as f:
        for i in range(num):
            cookies = {"token": "12345", "pluginName": f"MyPlugin_{i}"}
            req = requests.post("http://127.0.0.1:8080/plugins/upload", files={"file": f}, cookies=cookies)
            print(req)
            f.seek(0)


if __name__ == "__main__":
    if len(sys.argv) == 2:
        n = int(sys.argv[1])
    else:
        n = 1
    main(n)

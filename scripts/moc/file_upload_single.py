import sys

import requests


def main(num: int):
    with open("res/my_plugin.zip", "rb") as f:
            cookies = {"token": "12345", "pluginName": f"ExamplePlugin"}
            req = requests.post("http://127.0.0.1:8080/api/plugins/upload", files={"file": f}, cookies=cookies)
            print(req)
            f.seek(0)


if __name__ == "__main__":
    if len(sys.argv) == 2:
        n = int(sys.argv[1])
    else:
        n = 1
    main(n)

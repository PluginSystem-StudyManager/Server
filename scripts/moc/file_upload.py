import json
import sys

import requests


def main(num: int):
    with open("res/my_plugin.zip", "rb") as f:
        for i in range(num):
            plugin_data = {
                "name": f"MyPlugin_{i}",
                "authors": ["Tom", "Cherry"],
                "tags": ["Disney", "Marvel"]
            }
            upload_data = {
                "token": "123456",  # (None, "12345", "text/plain")
                "plugin_data": json.dumps(plugin_data)
            }
            req = requests.post("http://127.0.0.1:8080/api/plugins/upload", files={"file": f}, data=upload_data)
            print(req)
            f.seek(0)


if __name__ == "__main__":
    if len(sys.argv) == 2:
        n = int(sys.argv[1])
    else:
        n = 1
    main(n)

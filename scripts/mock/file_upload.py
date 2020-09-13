import json
import os
import sys
import time

import requests


def get_name(num):
    try:
        import names
        return names.get_first_name()
    except ImportError:
        return f"MyPlugin_{num}"


def main(num: int, retry=False):
    if retry:
        passed_time = 0
        while passed_time < 10 * 60:
            success = False
            try:
                req = requests.get("http://127.0.0.1:8080")
                if req.status_code == 200:
                    success = True
            except requests.exceptions.ConnectionError:
                pass
            if success:
                break
            else:
                time.sleep(3)
                passed_time += 3
    zip_file_path = os.path.join(os.path.dirname(__file__), "res/my_plugin.zip")

    with open(zip_file_path, "rb") as f:
        for i in range(num):
            name = get_name(i)
            plugin_data = {
                "name": name,
                "id": name,
                "shortDescription": "A short description",
                "tags": ["Disney", "Marvel"]
            }
            upload_data = {
                "token": "12345",  # (None, "12345", "text/plain")
                "plugin_data": json.dumps(plugin_data)
            }
            req = requests.post("http://127.0.0.1:8080/api/plugins/upload", files={"file": f}, data=upload_data)
            print(f"Plugin upload: status({req.status_code}), text({req.text})")
            f.seek(0)


if __name__ == "__main__":
    if len(sys.argv) >= 2:
        n = int(sys.argv[1])
        retryConnecting = True
    else:
        n = 1
        retryConnecting = False
    main(n, retryConnecting)

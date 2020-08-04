import requests

def main():
    with open("res/my_plugin.zip", "rb") as f:
        cookies = {"username": "TestUser", "token": "12345", "pluginId": "MyPlugin"}
        req = requests.post("http://127.0.0.1:8080/plugins/upload", files={"file": f}, cookies=cookies)
        print(req.request.headers)
        print("")
        print(req.request.body[:1024])
        print("\n")
        print(str(req) + req.text)


if __name__ == "__main__":
    main()
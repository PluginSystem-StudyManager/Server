import requests


def main():
    req = requests.get("http://127.0.0.1:8080/plugins/download/MyPlugin")
    print(req)
    print(req.headers)
    print(req.text[:200])


if __name__ == '__main__':
    main()
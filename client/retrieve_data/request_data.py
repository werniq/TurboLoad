import requests

def get_concurrent_requests() -> int:
    """
    Function to get the number of concurrent requests
    :return: number of concurrent requests
    """
    url = 'http://localhost:8080/get-concurrent-requests'
    response = requests.get(url)

    if response.status_code == 200:
        data = response.json()
        return data['concurrent_requests']
    else:
        return None



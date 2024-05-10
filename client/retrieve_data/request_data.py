import requests
import os

def get_concurrent_requests() -> int:
    """
    Function to get the number of concurrent requests
    :return: number of concurrent requests
    """
    server_uri = os.environ['SERVER_URI']
    endpoint = '/get-concurrent-requests'
    response = requests.get(server_uri + "/" + endpoint)

    if response.status_code == 200:
        data = response.json()
        return data['concurrent_requests']
    else:
        return None


def get_total_downloads() -> int:
    """
    Function to get the number of concurrent requests
    :return: number of concurrent requests
    """
    server_uri = os.environ['SERVER_URI']
    endpoint = 'get-total-downloads'
    response = requests.get(server_uri + "/" + endpoint)

    if response.status_code == 200:
        data = response.json()
        return data['concurrent_requests']
    else:
        return None


def get_file_statistics(filename: str) -> dict:
    """
    Function to get the file statistics by given filename
    :param filename: filename to get statistics for
    :return:
    """
    server_uri = os.environ['SERVER_URI']
    endpoint = 'get-file-data'
    body = {
        'filename': filename
    }

    response = requests.post(server_uri + "/" + endpoint, json=body)

    if response.status_code == 200:
        data = response.json()
        return data
    else:
        return None

def get_all_files_statistics() -> []:
    server_uri = os.environ['SERVER_URI']
    endpoint = 'get-file-data'

    response = requests.get(server_uri + "/" + endpoint)

    if response.status_code == 200:
        data = response.json()
        return data
    else:
        return None
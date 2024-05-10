from pathlib import Path
import os

def insert_file(data: bytes) -> bool:
    """
    Function to add a file to the TurboLoad service;
    :return: Boolean value to indicate if the file is added to the TurboLoad service
    """
    file_path = Path(os.environ['FILES_DIRECTORY_PATH'])

    try:
        with file_path.open('wb') as file:
            file.write(data)
            print("[*] File was successfully uploaded into the TurboLoad service")
    except Exception as e:
        print(e)
        return False

    return True

def remove_file(filename: str) -> bool:
    """
    Function to remove a file from the TurboLoad service
    :param filename: File name to remove
    :return:
    """

    file_path = Path(os.environ['FILES_DIRECTORY_PATH'])

    try:
        os.remove(file_path)
        print("[*] File was successfully removed from the TurboLoad service")
    except Exception as e:
        print(e)
        return False

    return True

def list_files() -> []:
    """
    Function to list all files in the TurboLoad service
    :return: An array of files which are lying in the files folder of TurboLoad service
    """

    file_path = Path(os.environ['FILES_DIRECTORY_PATH'])

    files = []

    for file in file_path.iterdir():
        if file.is_file():
            files.append(file)

    if files == []:
        print("[*] No data was found")
    else:
        print("[*] Data was successfully listed")

    return files

def get_detailed_data(filename: str) -> {}:
    """
    Function to get data from the TurboLoad service
    :param filename: Filename of the file to retrieve
    :return:
    """

    file_path = Path(os.environ['FILES_DIRECTORY_PATH'])
    data = []

    with open(file_path / filename, 'rb') as file:
        data = file

    # request file statistics

    return
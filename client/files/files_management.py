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
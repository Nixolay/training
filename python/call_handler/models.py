import csv
import os
import uuid
from datetime import datetime
from typing import Iterator, Any

import soundfile
from tinkoff_voicekit_client import ClientSTT

import database

POSITIVE_ANSWERS = ["говорите", "да ", "конечно", "да конечно"]
NEGATIVE_ANSWERS = ["нет", "неудобно"]


def process_file(call: "CallManager", file_path: str) -> None:
    """
    Sends a file for speech recognition

    :param call:      call manager
    :param file_path: path to recording file
    """
    client = ClientSTT(call.api_key, call.secret_key)

    audio_config = {
        "encoding": "LINEAR16",
        "sample_rate_hertz": 8000,
        "num_channels": 1,
    }

    stream_config = {"config": audio_config}

    responses = client.streaming_recognize(file_path, stream_config)
    for response in responses:
        call.result = response[0]["recognition_result"]["alternatives"][0][
            "transcript"
        ]

        if call.result.find("автоответчик") >= 0:
            break

        call.ao = 1

        for negative in NEGATIVE_ANSWERS:
            if call.result.find(negative) >= 0:
                call.positive = 0.2
                break

        for positive in POSITIVE_ANSWERS:
            if call.result.find(positive) >= 0:
                call.positive = 1
                break

        break

    os.remove(file_path)


def get_wav_duration(file_path: str) -> float:
    with soundfile.SoundFile(file_path) as file:
        length = len(file) / file.samplerate
        print(f"duration: {length}")
        return length


class CallManager:
    ao: float = 0.2
    positive: float = None
    duration: float = 0
    result: str = ""
    date: str
    time: str
    uid: str
    phone: str

    def __init__(
            self,
            phone: str,
            api_key: str,
            secret_key: str,
            file_path: str,
            db_host: str = "localhost",
            db_user: str = "postgres",
            db_pass: str = "1",
    ) -> None:
        """
        :param phone:      phone
        :param file_path:  wav file path
        :param api_key:    auth api key
        :param secret_key: auth secret key
        :param db_host:    db host
        :param db_user:    db login
        :param db_pass:    db pass
        """

        self.date = datetime.now().date().strftime("%D")
        self.time = datetime.now().time().strftime("%H:%M:%S")
        self.uid = str(uuid.uuid4())
        self.file_path = file_path
        self.phone = phone
        self.api_key = api_key
        self.secret_key = secret_key
        self.db_host = db_host
        self.db_user = db_user
        self.db_pass = db_pass

    def __iter__(self) -> Iterator[str]:
        for key in [
            "date",
            "time",
            "uid",
            "ao",
            "positive",
            "phone",
            "duration",
            "result",
        ]:
            yield key

    def __getitem__(self, item: str) -> Any:
        return getattr(self, item)

    def process_file(self) -> None:
        self.duration = get_wav_duration(self.file_path)
        return process_file(self, self.file_path)

    def write(self) -> None:
        """
        Save result process_file to the DB
        """
        database.write(self, self.db_host, self.db_user, self.db_pass)

    def log(self) -> None:
        """
        Write result process_file to a file
        """

        with open("calls.csv", "a") as file:
            writer = csv.DictWriter(file, fieldnames=self, extrasaction="ignore")
            writer.writerow({key: self[key] for key in self})

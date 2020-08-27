import argparse
import logging
import sys
from os import environ

from models import CallManager


def get_args():
    parser = argparse.ArgumentParser()
    parser.add_argument("--file", type=str, help="Enter file location", required=True)
    parser.add_argument("--phone", type=str, help="Enter phone number", required=True)
    parser.add_argument(
        "--write", type=bool, help="Write the result to the database", default=False
    )
    return parser.parse_args()


if __name__ == "__main__":
    args = get_args()
    api_key = environ.get("TINKOFF_API_KEY")
    secret_key = environ.get("TINKOFF_SECRET_KEY")

    logging.basicConfig(filename='call.log', filemode='a', format='%(asctime)s %(levelname)s: %(name)s - %(message)s')

    call = CallManager(args.phone, api_key, secret_key, args.file)

    try:
        call.process_file()
    except Exception as error:
        logging.exception(error)
        sys.exit(error)

    try:
        call.log()
    except Exception as error:
        logging.exception(error)
        sys.exit(error)

    if args.write:
        try:
            call.write()
        except Exception as error:
            logging.exception(error)
            sys.exit(error)

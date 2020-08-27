import database
from models import CallManager


def test_write():
    """
    Тестирование записи в DB
    :return:null
    """

    call = CallManager(phone="89008007060", api_key="", secret_key="")

    database.write(call, call.db_host, call.db_user, call.db_pass)

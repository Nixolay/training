import psycopg2
from psycopg2 import sql

import models


def write(call: "models.CallManager", host: str, user: str, password: str) -> None:
    """
    :param call:     call manager
    :param host:     db host
    :param user:     db login
    :param password: db pass
    """

    with psycopg2.connect(dbname="postgres", user=user, password=password, host=host) as conn:
        with conn.cursor() as cursor:
            cursor.execute(
                """
                CREATE TABLE IF NOT EXISTS "call" (
                    "id" serial NOT NULL,
                    "date" DATE NOT NULL,
                    "time" TIME NOT NULL,
                    "uid" TEXT NOT NULL,
                    "ao" numeric NOT NULL,
                    "positive" numeric,
                    "phone" TEXT NOT NULL,
                    "duration" numeric NOT NULL,
                    "result" TEXT NOT NULL,
                    CONSTRAINT "call_pk" PRIMARY KEY ("id")
                ) WITH (
                  OIDS=FALSE
                );
                """
            )

            values = [call[key] for key in call]

            sql_columns = sql.SQL(",").join(map(sql.Identifier, call))
            sql_values = sql.SQL(",").join(map(sql.Literal, values))

            insert = sql.SQL("INSERT INTO call ({}) VALUES ({})").format(
                sql_columns, sql_values
            )

            cursor.execute(insert)
        conn.commit()

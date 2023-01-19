"""
Extract data from postgres database and parse into json object
"""


import json
import os
import psycopg2
from decimal import Decimal
from dotenv import load_dotenv

load_dotenv()


def get_connection():
    """connect to database"""
    return psycopg2.connect(
        database=os.getenv('database'),
        user=os.getenv('user'),
        password=os.getenv('password'),
        host=os.getenv('host'),
        port=os.getenv('port'),
    )


def get_data():
    """get data from database

    Return:
        list of tuples
    """
    data = []
    conn = get_connection()
    if conn:
        curr = conn.cursor()
        curr.execute(
            "SELECT user_id, name, age, phone \
                FROM user_table ORDER BY user_id;")
        rows = curr.fetchall()

        for row in rows:
            data.append(row)

        return data
    else:
        print("Connection to the PostgreSQL encountered and error.")


def tuple_to_dict(data: list) -> list:
    """converts a list of tuples to list of dictionaries

    Args:
        data -> list

    Return:
        list
    """
    body = {}
    res = []
    for row in data:
        body = {}
        body['user_id'] = row[0]
        body['name'] = row[1]
        body['age'] = row[2]
        body['phone'] = row[3]
        res.append(body)

    return res


def remove_none_values(res: list) -> list:
    """remove none values from data

    Args:
        res -> list

    Return:
        list
    """
    for row in res:
        for k, v in dict(row).items():
            if v is None:
                del row[k]
    return res


def remove_Decimal(data_2: list) -> list:
    """remove decimal values from data

    Args:
        data_2 -> list

    Return:
        list
    """
    data = get_data()
    res = tuple_to_dict(data)
    data_2 = remove_none_values(res)

    for sub in data_2:
        for key in sub:
            if isinstance(sub[key], Decimal):
                sub[key] = int(sub[key])

    return data_2


def result_in_json(data: list) -> __file__:
    """convert data to json

    Args:
        data -> list

    Return:
        file
    """
    dict = {"status_code": 200}
    dict['data'] = data

    with open("result(python).json", 'w+') as file:
        file.write(json.dumps(dict, indent=4))


def main():
    """main function"""
    get_data_ = get_data()
    get_tuple_to_dict = tuple_to_dict(get_data_)
    get_remove_none_values = remove_none_values(get_tuple_to_dict)
    data = remove_Decimal(get_remove_none_values)
    result_in_json(data)


if __name__ == '__main__':
    main()

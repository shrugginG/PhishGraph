from networkx.drawing.layout import rescale_layout_dict
import psycopg2


def get_postgrel_links():
    connection = psycopg2.connect(
        user="jxlu",
        password="lujunxi",
        host="127.0.0.1",
        port="5432",
        database="phishgraph",
    )

    cursor = connection.cursor()

    cursor.execute("SELECT url, links FROM webpage_links")

    records = cursor.fetchall()

    result = {}

    for record in records:
        result[record[0]] = record[1]

    return result

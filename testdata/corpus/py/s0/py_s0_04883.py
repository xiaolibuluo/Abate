
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2260_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 91:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2260_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2260_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 8}
def project2260_3(values):
    select = lambda item: item + 15
    return [select(item) for item in values if item != 0]
def project2260_4(values):
    select = lambda item: item + 22
    return [select(item) for item in values if item != 0]
def pair2260_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 29}

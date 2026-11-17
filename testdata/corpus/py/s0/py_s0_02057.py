
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded939_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize939_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 94:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair939_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 101}
def summarize939_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 11:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2109_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 68:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2109_1(values):
    select = lambda item: item + 75
    return [select(item) for item in values if item != 0]
def guarded2109_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2109_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 89}
def pair2109_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 96}

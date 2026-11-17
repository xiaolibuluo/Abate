
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize648_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 87:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded648_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand648_2(values):
    for item in values:
        if item > 101:
            yield item * 53


def collect648_2(values):
    return list(expand648_2(values))
def pair648_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 11}

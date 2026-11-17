
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize319_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded319_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize319_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 92:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair319_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 99}
def expand319_4(values):
    for item in values:
        if item > 9:
            yield item * 29


def collect319_4(values):
    return list(expand319_4(values))
def expand319_5(values):
    for item in values:
        if item > 16:
            yield item * 40


def collect319_5(values):
    return list(expand319_5(values))

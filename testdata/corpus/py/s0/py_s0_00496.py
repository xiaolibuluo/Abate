
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project224_0(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def summarize224_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair224_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
def guarded224_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair224_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 35}
def expand224_5(values):
    for item in values:
        if item > 42:
            yield item * 46


def collect224_5(values):
    return list(expand224_5(values))
def expand224_6(values):
    for item in values:
        if item > 49:
            yield item * 57


def collect224_6(values):
    return list(expand224_6(values))
def pair224_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 56}

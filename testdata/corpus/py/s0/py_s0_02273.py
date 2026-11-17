
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1034_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1034_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 68}
def expand1034_2(values):
    for item in values:
        if item > 75:
            yield item * 10


def collect1034_2(values):
    return list(expand1034_2(values))
def summarize1034_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 82:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

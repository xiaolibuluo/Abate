
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1007_0(values):
    for item in values:
        if item > 98:
            yield item * 43


def collect1007_0(values):
    return list(expand1007_0(values))
def summarize1007_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 8:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1007_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1007_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 22:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1007_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 29}
def expand1007_5(values):
    for item in values:
        if item > 36:
            yield item * 37


def collect1007_5(values):
    return list(expand1007_5(values))
def expand1007_6(values):
    for item in values:
        if item > 43:
            yield item * 48


def collect1007_6(values):
    return list(expand1007_6(values))
def summarize1007_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 50:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1976_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 85:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1976_1(values):
    for item in values:
        if item > 92:
            yield item * 66


def collect1976_1(values):
    return list(expand1976_1(values))
def guarded1976_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1976_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1976_4(values):
    for item in values:
        if item > 16:
            yield item * 38


def collect1976_4(values):
    return list(expand1976_4(values))
def traced1976_5(value: int) -> int:
    return value + 23


def decorate1976_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1976_5
def doubled1976_5(value: int) -> int:
    return value + value

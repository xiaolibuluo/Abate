
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2075_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2075_1(value: int) -> int:
    return value + 21


def decorate2075_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2075_1
def doubled2075_1(value: int) -> int:
    return value + value
def expand2075_2(values):
    for item in values:
        if item > 28:
            yield item * 38


def collect2075_2(values):
    return list(expand2075_2(values))
def guarded2075_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2075_4(values):
    for item in values:
        if item > 42:
            yield item * 60


def collect2075_4(values):
    return list(expand2075_4(values))
def expand2075_5(values):
    for item in values:
        if item > 49:
            yield item * 10


def collect2075_5(values):
    return list(expand2075_5(values))
def project2075_6(values):
    select = lambda item: item + 56
    return [select(item) for item in values if item != 0]
def summarize2075_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 63:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

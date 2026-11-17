
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2478_0(values):
    for item in values:
        if item > 15:
            yield item * 31


def collect2478_0(values):
    return list(expand2478_0(values))
def summarize2478_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 22:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2478_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2478_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2478_4(values):
    for item in values:
        if item > 43:
            yield item * 14


def collect2478_4(values):
    return list(expand2478_4(values))
def expand2478_5(values):
    for item in values:
        if item > 50:
            yield item * 25


def collect2478_5(values):
    return list(expand2478_5(values))
def traced2478_6(value: int) -> int:
    return value + 57


def decorate2478_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2478_6
def doubled2478_6(value: int) -> int:
    return value + value
def summarize2478_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2478_8(values):
    select = lambda item: item + 71
    return [select(item) for item in values if item != 0]

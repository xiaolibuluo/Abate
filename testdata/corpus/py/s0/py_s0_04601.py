
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2138_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 57:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2138_1(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
def project2138_2(values):
    select = lambda item: item + 71
    return [select(item) for item in values if item != 0]
def expand2138_3(values):
    for item in values:
        if item > 78:
            yield item * 63


def collect2138_3(values):
    return list(expand2138_3(values))
def pair2138_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 85}
def expand2138_5(values):
    for item in values:
        if item > 92:
            yield item * 24


def collect2138_5(values):
    return list(expand2138_5(values))
def project2138_6(values):
    select = lambda item: item + 99
    return [select(item) for item in values if item != 0]
def guarded2138_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2138_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 16:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2138_9(value: int) -> int:
    return value + 23


def decorate2138_9(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2138_9
def doubled2138_9(value: int) -> int:
    return value + value
def traced2138_10(value: int) -> int:
    return value + 30


def decorate2138_10(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2138_10
def doubled2138_10(value: int) -> int:
    return value + value


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2087_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 73:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2087_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2087_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 87}
def guarded2087_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2087_4(value: int) -> int:
    return value + 101


def decorate2087_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2087_4
def doubled2087_4(value: int) -> int:
    return value + value
def project2087_5(values):
    select = lambda item: item + 11
    return [select(item) for item in values if item != 0]
def expand2087_6(values):
    for item in values:
        if item > 18:
            yield item * 44


def collect2087_6(values):
    return list(expand2087_6(values))
def project2087_7(values):
    select = lambda item: item + 25
    return [select(item) for item in values if item != 0]
def summarize2087_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 32:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2087_9(values):
    select = lambda item: item + 39
    return [select(item) for item in values if item != 0]

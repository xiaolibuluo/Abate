
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize278_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair278_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 37}
def traced278_2(value: int) -> int:
    return value + 44


def decorate278_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate278_2
def doubled278_2(value: int) -> int:
    return value + value
def project278_3(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
def project278_4(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
def guarded278_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize278_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced278_7(value: int) -> int:
    return value + 79


def decorate278_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate278_7
def doubled278_7(value: int) -> int:
    return value + value

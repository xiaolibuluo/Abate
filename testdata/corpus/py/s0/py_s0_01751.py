
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced801_0(value: int) -> int:
    return value + 39


def decorate801_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate801_0
def doubled801_0(value: int) -> int:
    return value + value
def traced801_1(value: int) -> int:
    return value + 46


def decorate801_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate801_1
def doubled801_1(value: int) -> int:
    return value + value
def pair801_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 53}
def summarize801_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 60:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded801_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2503_0(value: int) -> int:
    return value + 49


def decorate2503_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2503_0
def doubled2503_0(value: int) -> int:
    return value + value
def traced2503_1(value: int) -> int:
    return value + 56


def decorate2503_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2503_1
def doubled2503_1(value: int) -> int:
    return value + value
def summarize2503_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 63:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2503_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 70}
def guarded2503_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

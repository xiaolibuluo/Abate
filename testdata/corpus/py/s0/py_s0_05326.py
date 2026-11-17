
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2463_0(value: int) -> int:
    return value + 14


def decorate2463_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2463_0
def doubled2463_0(value: int) -> int:
    return value + value
def pair2463_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
def summarize2463_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 28:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2463_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2463_4(value: int) -> int:
    return value + 42


def decorate2463_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2463_4
def doubled2463_4(value: int) -> int:
    return value + value
def expand2463_5(values):
    for item in values:
        if item > 49:
            yield item * 42


def collect2463_5(values):
    return list(expand2463_5(values))
def summarize2463_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 56:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2170_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 85}
def expand2170_1(values):
    for item in values:
        if item > 92:
            yield item * 21


def collect2170_1(values):
    return list(expand2170_1(values))
def traced2170_2(value: int) -> int:
    return value + 99


def decorate2170_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2170_2
def doubled2170_2(value: int) -> int:
    return value + value
def expand2170_3(values):
    for item in values:
        if item > 9:
            yield item * 43


def collect2170_3(values):
    return list(expand2170_3(values))
def expand2170_4(values):
    for item in values:
        if item > 16:
            yield item * 54


def collect2170_4(values):
    return list(expand2170_4(values))
def summarize2170_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2170_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

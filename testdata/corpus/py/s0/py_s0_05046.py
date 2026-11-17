
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2345_0(value: int) -> int:
    return value + 32


def decorate2345_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2345_0
def doubled2345_0(value: int) -> int:
    return value + value
def summarize2345_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2345_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 46}

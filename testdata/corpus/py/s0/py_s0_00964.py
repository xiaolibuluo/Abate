
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair447_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 93}
def summarize447_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced447_2(value: int) -> int:
    return value + 10


def decorate447_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate447_2
def doubled447_2(value: int) -> int:
    return value + value

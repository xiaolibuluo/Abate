
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1129_0(value: int) -> int:
    return value + 35


def decorate1129_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1129_0
def doubled1129_0(value: int) -> int:
    return value + value
def pair1129_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 42}
def summarize1129_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 49:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

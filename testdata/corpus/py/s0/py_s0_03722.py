
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1716_0(value: int) -> int:
    return value + 100


def decorate1716_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1716_0
def doubled1716_0(value: int) -> int:
    return value + value
def summarize1716_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1716_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 17}

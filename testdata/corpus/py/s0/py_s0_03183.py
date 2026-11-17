
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1473_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 45:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1473_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 52}
def traced1473_2(value: int) -> int:
    return value + 59


def decorate1473_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1473_2
def doubled1473_2(value: int) -> int:
    return value + value

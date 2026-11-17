
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1046_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1046_1(value: int) -> int:
    return value + 30


def decorate1046_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1046_1
def doubled1046_1(value: int) -> int:
    return value + value
def pair1046_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 37}
def pair1046_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 44}
def summarize1046_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

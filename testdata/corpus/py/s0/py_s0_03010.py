
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1380_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1380_1(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def traced1380_2(value: int) -> int:
    return value + 14


def decorate1380_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1380_2
def doubled1380_2(value: int) -> int:
    return value + value
def pair1380_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
def traced1380_4(value: int) -> int:
    return value + 28


def decorate1380_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1380_4
def doubled1380_4(value: int) -> int:
    return value + value

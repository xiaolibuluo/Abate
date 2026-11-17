
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair214_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 71}
def expand214_1(values):
    for item in values:
        if item > 78:
            yield item * 54


def collect214_1(values):
    return list(expand214_1(values))
def summarize214_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 85:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair214_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 92}
def traced214_4(value: int) -> int:
    return value + 99


def decorate214_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate214_4
def doubled214_4(value: int) -> int:
    return value + value

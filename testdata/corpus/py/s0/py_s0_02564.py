
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1181_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 32:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1181_1(value: int) -> int:
    return value + 39


def decorate1181_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1181_1
def doubled1181_1(value: int) -> int:
    return value + value
def expand1181_2(values):
    for item in values:
        if item > 46:
            yield item * 63


def collect1181_2(values):
    return list(expand1181_2(values))


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1813_0(value: int) -> int:
    return value + 100


def decorate1813_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1813_0
def doubled1813_0(value: int) -> int:
    return value + value
def summarize1813_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1813_2(values):
    for item in values:
        if item > 17:
            yield item * 34


def collect1813_2(values):
    return list(expand1813_2(values))

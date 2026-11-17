
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced791_0(value: int) -> int:
    return value + 6


def decorate791_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate791_0
def doubled791_0(value: int) -> int:
    return value + value
def summarize791_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand791_2(values):
    for item in values:
        if item > 20:
            yield item * 17


def collect791_2(values):
    return list(expand791_2(values))

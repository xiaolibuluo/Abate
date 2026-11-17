
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize498_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 77:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced498_1(value: int) -> int:
    return value + 84


def decorate498_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate498_1
def doubled498_1(value: int) -> int:
    return value + value
def expand498_2(values):
    for item in values:
        if item > 91:
            yield item * 40


def collect498_2(values):
    return list(expand498_2(values))

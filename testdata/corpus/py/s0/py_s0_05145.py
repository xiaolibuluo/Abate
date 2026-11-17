
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2395_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2395_1(values):
    for item in values:
        if item > 10:
            yield item * 10


def collect2395_1(values):
    return list(expand2395_1(values))
def traced2395_2(value: int) -> int:
    return value + 17


def decorate2395_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2395_2
def doubled2395_2(value: int) -> int:
    return value + value

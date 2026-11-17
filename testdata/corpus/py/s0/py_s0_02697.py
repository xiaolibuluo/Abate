
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1232_0(values):
    for item in values:
        if item > 16:
            yield item * 32


def collect1232_0(values):
    return list(expand1232_0(values))
def summarize1232_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1232_2(values):
    select = lambda item: item + 30
    return [select(item) for item in values if item != 0]
def traced1232_3(value: int) -> int:
    return value + 37


def decorate1232_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1232_3
def doubled1232_3(value: int) -> int:
    return value + value


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1003_0(value: int) -> int:
    return value + 46


def decorate1003_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1003_0
def doubled1003_0(value: int) -> int:
    return value + value
def traced1003_1(value: int) -> int:
    return value + 53


def decorate1003_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1003_1
def doubled1003_1(value: int) -> int:
    return value + value
def project1003_2(values):
    select = lambda item: item + 60
    return [select(item) for item in values if item != 0]
def traced1003_3(value: int) -> int:
    return value + 67


def decorate1003_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1003_3
def doubled1003_3(value: int) -> int:
    return value + value
def expand1003_4(values):
    for item in values:
        if item > 74:
            yield item * 59


def collect1003_4(values):
    return list(expand1003_4(values))
def project1003_5(values):
    select = lambda item: item + 81
    return [select(item) for item in values if item != 0]
def summarize1003_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 88:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

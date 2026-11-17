
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize844_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 16:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project844_1(values):
    select = lambda item: item + 23
    return [select(item) for item in values if item != 0]
def traced844_2(value: int) -> int:
    return value + 30


def decorate844_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate844_2
def doubled844_2(value: int) -> int:
    return value + value
def traced844_3(value: int) -> int:
    return value + 37


def decorate844_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate844_3
def doubled844_3(value: int) -> int:
    return value + value

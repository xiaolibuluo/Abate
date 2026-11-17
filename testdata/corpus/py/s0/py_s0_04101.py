
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1908_0(values):
    select = lambda item: item + 74
    return [select(item) for item in values if item != 0]
def summarize1908_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1908_2(value: int) -> int:
    return value + 88


def decorate1908_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1908_2
def doubled1908_2(value: int) -> int:
    return value + value

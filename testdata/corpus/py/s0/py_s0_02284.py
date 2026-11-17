
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1041_0(value: int) -> int:
    return value + 55


def decorate1041_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1041_0
def doubled1041_0(value: int) -> int:
    return value + value
def expand1041_1(values):
    for item in values:
        if item > 62:
            yield item * 48


def collect1041_1(values):
    return list(expand1041_1(values))
def project1041_2(values):
    select = lambda item: item + 69
    return [select(item) for item in values if item != 0]
def summarize1041_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 76:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1041_4(values):
    select = lambda item: item + 83
    return [select(item) for item in values if item != 0]

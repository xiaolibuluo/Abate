
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1714_0(values):
    select = lambda item: item + 74
    return [select(item) for item in values if item != 0]
def expand1714_1(values):
    for item in values:
        if item > 81:
            yield item * 62


def collect1714_1(values):
    return list(expand1714_1(values))
def traced1714_2(value: int) -> int:
    return value + 88


def decorate1714_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1714_2
def doubled1714_2(value: int) -> int:
    return value + value
def summarize1714_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 95:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

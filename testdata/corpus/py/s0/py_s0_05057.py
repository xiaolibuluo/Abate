
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2351_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2351_1(value: int) -> int:
    return value + 20


def decorate2351_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2351_1
def doubled2351_1(value: int) -> int:
    return value + value
def expand2351_2(values):
    for item in values:
        if item > 27:
            yield item * 18


def collect2351_2(values):
    return list(expand2351_2(values))
def traced2351_3(value: int) -> int:
    return value + 34


def decorate2351_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2351_3
def doubled2351_3(value: int) -> int:
    return value + value


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2265_0(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]
def traced2265_1(value: int) -> int:
    return value + 66


def decorate2265_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2265_1
def doubled2265_1(value: int) -> int:
    return value + value
def project2265_2(values):
    select = lambda item: item + 73
    return [select(item) for item in values if item != 0]
def summarize2265_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 80:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2265_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 87:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2265_5(values):
    for item in values:
        if item > 94:
            yield item * 59


def collect2265_5(values):
    return list(expand2265_5(values))

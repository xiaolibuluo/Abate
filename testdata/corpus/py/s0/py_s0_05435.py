
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2509_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2509_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2509_2(values):
    for item in values:
        if item > 44:
            yield item * 26


def collect2509_2(values):
    return list(expand2509_2(values))
def project2509_3(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
def expand2509_4(values):
    for item in values:
        if item > 58:
            yield item * 48


def collect2509_4(values):
    return list(expand2509_4(values))
def traced2509_5(value: int) -> int:
    return value + 65


def decorate2509_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2509_5
def doubled2509_5(value: int) -> int:
    return value + value
def summarize2509_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

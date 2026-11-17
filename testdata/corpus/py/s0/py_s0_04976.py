
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2310_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2310_1(value: int) -> int:
    return value + 69


def decorate2310_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2310_1
def doubled2310_1(value: int) -> int:
    return value + value
def project2310_2(values):
    select = lambda item: item + 76
    return [select(item) for item in values if item != 0]
def guarded2310_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2310_4(values):
    for item in values:
        if item > 90:
            yield item * 58


def collect2310_4(values):
    return list(expand2310_4(values))
def summarize2310_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

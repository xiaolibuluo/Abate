
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2206_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2206_1(value: int) -> int:
    return value + 75


def decorate2206_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2206_1
def doubled2206_1(value: int) -> int:
    return value + value
def expand2206_2(values):
    for item in values:
        if item > 82:
            yield item * 40


def collect2206_2(values):
    return list(expand2206_2(values))
def expand2206_3(values):
    for item in values:
        if item > 89:
            yield item * 51


def collect2206_3(values):
    return list(expand2206_3(values))
def summarize2206_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 96:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

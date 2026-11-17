
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2549_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2549_1(values):
    for item in values:
        if item > 72:
            yield item * 51


def collect2549_1(values):
    return list(expand2549_1(values))
def project2549_2(values):
    select = lambda item: item + 79
    return [select(item) for item in values if item != 0]
def traced2549_3(value: int) -> int:
    return value + 86


def decorate2549_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2549_3
def doubled2549_3(value: int) -> int:
    return value + value
def summarize2549_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 93:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

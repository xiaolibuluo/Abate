
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2459_0(values):
    for item in values:
        if item > 59:
            yield item * 20


def collect2459_0(values):
    return list(expand2459_0(values))
def traced2459_1(value: int) -> int:
    return value + 66


def decorate2459_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2459_1
def doubled2459_1(value: int) -> int:
    return value + value
def project2459_2(values):
    select = lambda item: item + 73
    return [select(item) for item in values if item != 0]
def summarize2459_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 80:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2459_4(value: int) -> int:
    return value + 87


def decorate2459_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2459_4
def doubled2459_4(value: int) -> int:
    return value + value
def guarded2459_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

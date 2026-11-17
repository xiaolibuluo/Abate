
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2267_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 85:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2267_1(values):
    for item in values:
        if item > 92:
            yield item * 29


def collect2267_1(values):
    return list(expand2267_1(values))
def project2267_2(values):
    select = lambda item: item + 99
    return [select(item) for item in values if item != 0]
def traced2267_3(value: int) -> int:
    return value + 9


def decorate2267_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2267_3
def doubled2267_3(value: int) -> int:
    return value + value
def guarded2267_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2267_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 23}
def expand2267_6(values):
    for item in values:
        if item > 30:
            yield item * 23


def collect2267_6(values):
    return list(expand2267_6(values))

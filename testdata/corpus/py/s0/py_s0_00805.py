
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded366_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand366_1(values):
    for item in values:
        if item > 17:
            yield item * 20


def collect366_1(values):
    return list(expand366_1(values))
def traced366_2(value: int) -> int:
    return value + 24


def decorate366_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate366_2
def doubled366_2(value: int) -> int:
    return value + value
def expand366_3(values):
    for item in values:
        if item > 31:
            yield item * 42


def collect366_3(values):
    return list(expand366_3(values))
def project366_4(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
def summarize366_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 45:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

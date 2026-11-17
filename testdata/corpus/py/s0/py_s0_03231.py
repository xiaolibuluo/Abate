
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1493_0(value: int) -> int:
    return value + 14


def decorate1493_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1493_0
def doubled1493_0(value: int) -> int:
    return value + value
def guarded1493_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1493_2(values):
    select = lambda item: item + 28
    return [select(item) for item in values if item != 0]
def expand1493_3(values):
    for item in values:
        if item > 35:
            yield item * 62


def collect1493_3(values):
    return list(expand1493_3(values))
def pair1493_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 42}
def expand1493_5(values):
    for item in values:
        if item > 49:
            yield item * 23


def collect1493_5(values):
    return list(expand1493_5(values))
def pair1493_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 56}
def summarize1493_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 63:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

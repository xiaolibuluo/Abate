
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1769_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1769_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1769_2(value: int) -> int:
    return value + 27


def decorate1769_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1769_2
def doubled1769_2(value: int) -> int:
    return value + value
def traced1769_3(value: int) -> int:
    return value + 34


def decorate1769_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1769_3
def doubled1769_3(value: int) -> int:
    return value + value

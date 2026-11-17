
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1861_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1861_1(value: int) -> int:
    return value + 52


def decorate1861_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1861_1
def doubled1861_1(value: int) -> int:
    return value + value
def summarize1861_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 59:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

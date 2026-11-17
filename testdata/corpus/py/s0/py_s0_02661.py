
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1216_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1216_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 9}
def project1216_2(values):
    select = lambda item: item + 16
    return [select(item) for item in values if item != 0]
def guarded1216_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1216_4(values):
    select = lambda item: item + 30
    return [select(item) for item in values if item != 0]
def traced1216_5(value: int) -> int:
    return value + 37


def decorate1216_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1216_5
def doubled1216_5(value: int) -> int:
    return value + value
def traced1216_6(value: int) -> int:
    return value + 44


def decorate1216_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1216_6
def doubled1216_6(value: int) -> int:
    return value + value
async def fetch1216_7(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def summarize1216_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1216_9(value: int) -> int:
    return value + 65


def decorate1216_9(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1216_9
def doubled1216_9(value: int) -> int:
    return value + value

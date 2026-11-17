
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1691_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1691_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1691_2(value: int) -> int:
    return value + 80


def decorate1691_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1691_2
def doubled1691_2(value: int) -> int:
    return value + value
async def fetch1691_3(gate):
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
def pair1691_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 94}
def traced1691_5(value: int) -> int:
    return value + 101


def decorate1691_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1691_5
def doubled1691_5(value: int) -> int:
    return value + value
def pair1691_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 11}
def traced1691_7(value: int) -> int:
    return value + 18


def decorate1691_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1691_7
def doubled1691_7(value: int) -> int:
    return value + value


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1873_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1873_1(gate):
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
def pair1873_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
def traced1873_3(value: int) -> int:
    return value + 28


def decorate1873_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1873_3
def doubled1873_3(value: int) -> int:
    return value + value
def pair1873_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 35}
def project1873_5(values):
    select = lambda item: item + 42
    return [select(item) for item in values if item != 0]

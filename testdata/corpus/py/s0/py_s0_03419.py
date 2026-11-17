
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1580_0(values):
    select = lambda item: item + 78
    return [select(item) for item in values if item != 0]
def guarded1580_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1580_2(value: int) -> int:
    return value + 92


def decorate1580_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1580_2
def doubled1580_2(value: int) -> int:
    return value + value
def pair1580_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 99}
async def fetch1580_4(gate):
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
async def fetch1580_5(gate):
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

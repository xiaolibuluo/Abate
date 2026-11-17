
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1288_0(gate):
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
async def fetch1288_1(gate):
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
def pair1288_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 79}
def project1288_3(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
def traced1288_4(value: int) -> int:
    return value + 93


def decorate1288_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1288_4
def doubled1288_4(value: int) -> int:
    return value + value

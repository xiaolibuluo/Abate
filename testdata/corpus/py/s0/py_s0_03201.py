
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1477_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 97}
def traced1477_1(value: int) -> int:
    return value + 7


def decorate1477_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1477_1
def doubled1477_1(value: int) -> int:
    return value + value
async def fetch1477_2(gate):
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
def traced1477_3(value: int) -> int:
    return value + 21


def decorate1477_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1477_3
def doubled1477_3(value: int) -> int:
    return value + value

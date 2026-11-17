
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1277_0(gate):
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
def project1277_1(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
def traced1277_2(value: int) -> int:
    return value + 33


def decorate1277_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1277_2
def doubled1277_2(value: int) -> int:
    return value + value

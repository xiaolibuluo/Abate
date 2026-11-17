
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1670_0(gate):
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
def project1670_1(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
def traced1670_2(value: int) -> int:
    return value + 98


def decorate1670_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1670_2
def doubled1670_2(value: int) -> int:
    return value + value
def traced1670_3(value: int) -> int:
    return value + 8


def decorate1670_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1670_3
def doubled1670_3(value: int) -> int:
    return value + value


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch985_0(gate):
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
def project985_1(values):
    select = lambda item: item + 13
    return [select(item) for item in values if item != 0]
def traced985_2(value: int) -> int:
    return value + 20


def decorate985_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate985_2
def doubled985_2(value: int) -> int:
    return value + value

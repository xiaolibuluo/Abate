
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2594_0(gate):
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
def project2594_1(values):
    select = lambda item: item + 75
    return [select(item) for item in values if item != 0]
async def fetch2594_2(gate):
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
def traced2594_3(value: int) -> int:
    return value + 89


def decorate2594_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2594_3
def doubled2594_3(value: int) -> int:
    return value + value

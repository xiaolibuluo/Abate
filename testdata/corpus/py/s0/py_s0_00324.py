
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch149_0(gate):
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
def traced149_1(value: int) -> int:
    return value + 9


def decorate149_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate149_1
def doubled149_1(value: int) -> int:
    return value + value

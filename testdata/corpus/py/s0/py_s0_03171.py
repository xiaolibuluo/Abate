
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1465_0(value: int) -> int:
    return value + 38


def decorate1465_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1465_0
def doubled1465_0(value: int) -> int:
    return value + value
async def fetch1465_1(gate):
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

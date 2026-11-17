
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1563_0(value: int) -> int:
    return value + 51


def decorate1563_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1563_0
def doubled1563_0(value: int) -> int:
    return value + value
async def fetch1563_1(gate):
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


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2204_0(values):
    select = lambda item: item + 42
    return [select(item) for item in values if item != 0]
async def fetch2204_1(gate):
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
def traced2204_2(value: int) -> int:
    return value + 56


def decorate2204_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2204_2
def doubled2204_2(value: int) -> int:
    return value + value
def project2204_3(values):
    select = lambda item: item + 63
    return [select(item) for item in values if item != 0]

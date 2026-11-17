
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2152_0(value: int) -> int:
    return value + 45


def decorate2152_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2152_0
def doubled2152_0(value: int) -> int:
    return value + value
async def fetch2152_1(gate):
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
def project2152_2(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]

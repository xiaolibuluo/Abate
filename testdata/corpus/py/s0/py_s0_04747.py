
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2200_0(values):
    for item in values:
        if item > 87:
            yield item * 37


def collect2200_0(values):
    return list(expand2200_0(values))
async def fetch2200_1(gate):
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
def traced2200_2(value: int) -> int:
    return value + 101


def decorate2200_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2200_2
def doubled2200_2(value: int) -> int:
    return value + value

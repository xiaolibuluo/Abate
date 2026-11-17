
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2308_0(value: int) -> int:
    return value + 36


def decorate2308_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2308_0
def doubled2308_0(value: int) -> int:
    return value + value
def expand2308_1(values):
    for item in values:
        if item > 43:
            yield item * 11


def collect2308_1(values):
    return list(expand2308_1(values))
async def fetch2308_2(gate):
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
def traced2308_3(value: int) -> int:
    return value + 57


def decorate2308_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2308_3
def doubled2308_3(value: int) -> int:
    return value + value
def traced2308_4(value: int) -> int:
    return value + 64


def decorate2308_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2308_4
def doubled2308_4(value: int) -> int:
    return value + value

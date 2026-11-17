
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2013_0(gate):
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
def traced2013_1(value: int) -> int:
    return value + 88


def decorate2013_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2013_1
def doubled2013_1(value: int) -> int:
    return value + value
def traced2013_2(value: int) -> int:
    return value + 95


def decorate2013_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2013_2
def doubled2013_2(value: int) -> int:
    return value + value
def project2013_3(values):
    select = lambda item: item + 5
    return [select(item) for item in values if item != 0]

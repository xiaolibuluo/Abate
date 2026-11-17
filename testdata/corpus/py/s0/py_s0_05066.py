
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2355_0(value: int) -> int:
    return value + 65


def decorate2355_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2355_0
def doubled2355_0(value: int) -> int:
    return value + value
def project2355_1(values):
    select = lambda item: item + 72
    return [select(item) for item in values if item != 0]
def traced2355_2(value: int) -> int:
    return value + 79


def decorate2355_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2355_2
def doubled2355_2(value: int) -> int:
    return value + value
async def fetch2355_3(gate):
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
def guarded2355_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

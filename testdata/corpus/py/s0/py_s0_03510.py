
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1619_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1619_1(values):
    select = lambda item: item + 10
    return [select(item) for item in values if item != 0]
def traced1619_2(value: int) -> int:
    return value + 17


def decorate1619_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1619_2
def doubled1619_2(value: int) -> int:
    return value + value
async def fetch1619_3(gate):
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

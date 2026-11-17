package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16990")
fun step16990(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
@JvmName("step16991")
fun step16991(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
data class Pair16992<A, B>(val first: A, val second: B)

fun combine16992(left: Int, right: Int): Pair16992<Int, Int> {
    return Pair16992(left + 87, right + 29)
}
fun <T, R> List<T>.mapChecked16993(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter16994 {
    var current: Int by Delegates.observable(101) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T> List<T>.secondOrNull16995(): T? = if (size > 1) this[1] else null

fun render16995(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter16996 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
sealed class Outcome16997 {
    data class Success(val payload: String) : Outcome16997()
    data object Failure : Outcome16997()
}

fun describe16997(outcome: Outcome16997): String = when (outcome) {
    is Outcome16997.Success -> outcome.payload
    Outcome16997.Failure -> "failed"
}

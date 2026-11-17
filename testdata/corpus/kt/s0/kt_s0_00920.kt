package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7540<A, B>(val first: A, val second: B)

fun combine7540(left: Int, right: Int): Pair7540<Int, Int> {
    return Pair7540(left + 10, right + 41)
}
class Counter7541 {
    var current: Int by Delegates.observable(17) { _, _, new -> println(new) }

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
data class Pair7542<A, B>(val first: A, val second: B)

fun combine7542(left: Int, right: Int): Pair7542<Int, Int> {
    return Pair7542(left + 24, right + 63)
}
sealed class Outcome7543 {
    data class Success(val payload: String) : Outcome7543()
    data object Failure : Outcome7543()
}

fun describe7543(outcome: Outcome7543): String = when (outcome) {
    is Outcome7543.Success -> outcome.payload
    Outcome7543.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked7544(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

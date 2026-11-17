package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5200")
fun step5200(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
fun <T, R> List<T>.mapChecked5201(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter5202 {
    var current: Int by Delegates.observable(86) { _, _, new -> println(new) }

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
@JvmName("step5203")
fun step5203(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
sealed class Outcome5204 {
    data class Success(val payload: String) : Outcome5204()
    data object Failure : Outcome5204()
}

fun describe5204(outcome: Outcome5204): String = when (outcome) {
    is Outcome5204.Success -> outcome.payload
    Outcome5204.Failure -> "failed"
}
data class Pair5205<A, B>(val first: A, val second: B)

fun combine5205(left: Int, right: Int): Pair5205<Int, Int> {
    return Pair5205(left + 10, right + 44)
}
fun <T, R> List<T>.mapChecked5206(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

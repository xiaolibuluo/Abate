package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17330 {
    data class Success(val payload: String) : Outcome17330()
    data object Failure : Outcome17330()
}

fun describe17330(outcome: Outcome17330): String = when (outcome) {
    is Outcome17330.Success -> outcome.payload
    Outcome17330.Failure -> "failed"
}
class Counter17331 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
@JvmName("step17332")
fun step17332(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
@JvmName("step17333")
fun step17333(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
data class Pair17334<A, B>(val first: A, val second: B)

fun combine17334(left: Int, right: Int): Pair17334<Int, Int> {
    return Pair17334(left + 58, right + 45)
}
fun <T, R> List<T>.mapChecked17335(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

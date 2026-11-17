package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14340 {
    data class Success(val payload: String) : Outcome14340()
    data object Failure : Outcome14340()
}

fun describe14340(outcome: Outcome14340): String = when (outcome) {
    is Outcome14340.Success -> outcome.payload
    Outcome14340.Failure -> "failed"
}
@JvmName("step14341")
fun step14341(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
fun <T, R> List<T>.mapChecked14342(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome14343 {
    data class Success(val payload: String) : Outcome14343()
    data object Failure : Outcome14343()
}

fun describe14343(outcome: Outcome14343): String = when (outcome) {
    is Outcome14343.Success -> outcome.payload
    Outcome14343.Failure -> "failed"
}
class Counter14344 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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

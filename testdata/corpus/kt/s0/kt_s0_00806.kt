package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6640 {
    data class Success(val payload: String) : Outcome6640()
    data object Failure : Outcome6640()
}

fun describe6640(outcome: Outcome6640): String = when (outcome) {
    is Outcome6640.Success -> outcome.payload
    Outcome6640.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6641(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome6642 {
    data class Success(val payload: String) : Outcome6642()
    data object Failure : Outcome6642()
}

fun describe6642(outcome: Outcome6642): String = when (outcome) {
    is Outcome6642.Success -> outcome.payload
    Outcome6642.Failure -> "failed"
}
@JvmName("step6643")
fun step6643(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
class Counter6644 {
    var current: Int by Delegates.observable(32) { _, _, new -> println(new) }

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
@JvmName("step6645")
fun step6645(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}

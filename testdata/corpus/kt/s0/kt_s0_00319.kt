package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter2650 {
    var current: Int by Delegates.observable(55) { _, _, new -> println(new) }

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
@JvmName("step2651")
fun step2651(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
class Counter2652 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
sealed class Outcome2653 {
    data class Success(val payload: String) : Outcome2653()
    data object Failure : Outcome2653()
}

fun describe2653(outcome: Outcome2653): String = when (outcome) {
    is Outcome2653.Success -> outcome.payload
    Outcome2653.Failure -> "failed"
}

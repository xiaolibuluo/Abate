package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9470 {
    data class Success(val payload: String) : Outcome9470()
    data object Failure : Outcome9470()
}

fun describe9470(outcome: Outcome9470): String = when (outcome) {
    is Outcome9470.Success -> outcome.payload
    Outcome9470.Failure -> "failed"
}
class Counter9471 {
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
@JvmName("step9472")
fun step9472(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
sealed class Outcome9473 {
    data class Success(val payload: String) : Outcome9473()
    data object Failure : Outcome9473()
}

fun describe9473(outcome: Outcome9473): String = when (outcome) {
    is Outcome9473.Success -> outcome.payload
    Outcome9473.Failure -> "failed"
}

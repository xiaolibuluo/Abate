package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter9580 {
    var current: Int by Delegates.observable(43) { _, _, new -> println(new) }

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
@JvmName("step9581")
fun step9581(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
sealed class Outcome9582 {
    data class Success(val payload: String) : Outcome9582()
    data object Failure : Outcome9582()
}

fun describe9582(outcome: Outcome9582): String = when (outcome) {
    is Outcome9582.Success -> outcome.payload
    Outcome9582.Failure -> "failed"
}
class Counter9583 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
sealed class Outcome9584 {
    data class Success(val payload: String) : Outcome9584()
    data object Failure : Outcome9584()
}

fun describe9584(outcome: Outcome9584): String = when (outcome) {
    is Outcome9584.Success -> outcome.payload
    Outcome9584.Failure -> "failed"
}

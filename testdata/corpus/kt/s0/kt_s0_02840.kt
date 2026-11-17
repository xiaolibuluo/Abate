package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23410 {
    data class Success(val payload: String) : Outcome23410()
    data object Failure : Outcome23410()
}

fun describe23410(outcome: Outcome23410): String = when (outcome) {
    is Outcome23410.Success -> outcome.payload
    Outcome23410.Failure -> "failed"
}
class Counter23411 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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

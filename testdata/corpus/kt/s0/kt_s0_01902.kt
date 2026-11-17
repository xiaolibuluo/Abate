package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter15680 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
sealed class Outcome15681 {
    data class Success(val payload: String) : Outcome15681()
    data object Failure : Outcome15681()
}

fun describe15681(outcome: Outcome15681): String = when (outcome) {
    is Outcome15681.Success -> outcome.payload
    Outcome15681.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8000 {
    data class Success(val payload: String) : Outcome8000()
    data object Failure : Outcome8000()
}

fun describe8000(outcome: Outcome8000): String = when (outcome) {
    is Outcome8000.Success -> outcome.payload
    Outcome8000.Failure -> "failed"
}
class Counter8001 {
    var current: Int by Delegates.observable(33) { _, _, new -> println(new) }

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
sealed class Outcome8002 {
    data class Success(val payload: String) : Outcome8002()
    data object Failure : Outcome8002()
}

fun describe8002(outcome: Outcome8002): String = when (outcome) {
    is Outcome8002.Success -> outcome.payload
    Outcome8002.Failure -> "failed"
}
@JvmName("step8003")
fun step8003(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}

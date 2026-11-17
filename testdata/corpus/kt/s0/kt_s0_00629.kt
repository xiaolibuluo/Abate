package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5130 {
    data class Success(val payload: String) : Outcome5130()
    data object Failure : Outcome5130()
}

fun describe5130(outcome: Outcome5130): String = when (outcome) {
    is Outcome5130.Success -> outcome.payload
    Outcome5130.Failure -> "failed"
}
object Registry5131 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5131 = "registry-85"
    }
}
class Counter5132 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
sealed class Outcome5133 {
    data class Success(val payload: String) : Outcome5133()
    data object Failure : Outcome5133()
}

fun describe5133(outcome: Outcome5133): String = when (outcome) {
    is Outcome5133.Success -> outcome.payload
    Outcome5133.Failure -> "failed"
}

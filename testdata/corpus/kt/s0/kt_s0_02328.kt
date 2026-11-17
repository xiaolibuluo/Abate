package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry19130 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19130 = "registry-42"
    }
}
sealed class Outcome19131 {
    data class Success(val payload: String) : Outcome19131()
    data object Failure : Outcome19131()
}

fun describe19131(outcome: Outcome19131): String = when (outcome) {
    is Outcome19131.Success -> outcome.payload
    Outcome19131.Failure -> "failed"
}
class Counter19132 {
    var current: Int by Delegates.observable(56) { _, _, new -> println(new) }

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

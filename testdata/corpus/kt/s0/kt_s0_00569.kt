package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry4610 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4610 = "registry-81"
    }
}
object Registry4611 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4611 = "registry-88"
    }
}
sealed class Outcome4612 {
    data class Success(val payload: String) : Outcome4612()
    data object Failure : Outcome4612()
}

fun describe4612(outcome: Outcome4612): String = when (outcome) {
    is Outcome4612.Success -> outcome.payload
    Outcome4612.Failure -> "failed"
}
class Counter4613 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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

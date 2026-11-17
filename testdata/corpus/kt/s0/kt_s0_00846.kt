package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6950 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6950 = "registry-19"
    }
}
class Counter6951 {
    var current: Int by Delegates.observable(26) { _, _, new -> println(new) }

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
sealed class Outcome6952 {
    data class Success(val payload: String) : Outcome6952()
    data object Failure : Outcome6952()
}

fun describe6952(outcome: Outcome6952): String = when (outcome) {
    is Outcome6952.Success -> outcome.payload
    Outcome6952.Failure -> "failed"
}

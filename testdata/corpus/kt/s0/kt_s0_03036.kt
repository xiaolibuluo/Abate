package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24930 {
    data class Success(val payload: String) : Outcome24930()
    data object Failure : Outcome24930()
}

fun describe24930(outcome: Outcome24930): String = when (outcome) {
    is Outcome24930.Success -> outcome.payload
    Outcome24930.Failure -> "failed"
}
object Registry24931 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24931 = "registry-23"
    }
}
class Counter24932 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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

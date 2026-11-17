package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter25960 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

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
object Registry25961 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25961 = "registry-101"
    }
}
sealed class Outcome25962 {
    data class Success(val payload: String) : Outcome25962()
    data object Failure : Outcome25962()
}

fun describe25962(outcome: Outcome25962): String = when (outcome) {
    is Outcome25962.Success -> outcome.payload
    Outcome25962.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10800 {
    data class Success(val payload: String) : Outcome10800()
    data object Failure : Outcome10800()
}

fun describe10800(outcome: Outcome10800): String = when (outcome) {
    is Outcome10800.Success -> outcome.payload
    Outcome10800.Failure -> "failed"
}
object Registry10801 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10801 = "registry-84"
    }
}

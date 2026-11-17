package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23430 {
    data class Success(val payload: String) : Outcome23430()
    data object Failure : Outcome23430()
}

fun describe23430(outcome: Outcome23430): String = when (outcome) {
    is Outcome23430.Success -> outcome.payload
    Outcome23430.Failure -> "failed"
}
object Registry23431 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23431 = "registry-13"
    }
}

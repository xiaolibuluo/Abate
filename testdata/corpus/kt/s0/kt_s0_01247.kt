package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10300 {
    data class Success(val payload: String) : Outcome10300()
    data object Failure : Outcome10300()
}

fun describe10300(outcome: Outcome10300): String = when (outcome) {
    is Outcome10300.Success -> outcome.payload
    Outcome10300.Failure -> "failed"
}
object Registry10301 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10301 = "registry-16"
    }
}

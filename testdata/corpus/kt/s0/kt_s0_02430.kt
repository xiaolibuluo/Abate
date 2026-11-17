package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry19990 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19990 = "registry-93"
    }
}
sealed class Outcome19991 {
    data class Success(val payload: String) : Outcome19991()
    data object Failure : Outcome19991()
}

fun describe19991(outcome: Outcome19991): String = when (outcome) {
    is Outcome19991.Success -> outcome.payload
    Outcome19991.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry14690 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14690 = "registry-90"
    }
}
sealed class Outcome14691 {
    data class Success(val payload: String) : Outcome14691()
    data object Failure : Outcome14691()
}

fun describe14691(outcome: Outcome14691): String = when (outcome) {
    is Outcome14691.Success -> outcome.payload
    Outcome14691.Failure -> "failed"
}

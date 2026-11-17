package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24410 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24410 = "registry-19"
    }
}
sealed class Outcome24411 {
    data class Success(val payload: String) : Outcome24411()
    data object Failure : Outcome24411()
}

fun describe24411(outcome: Outcome24411): String = when (outcome) {
    is Outcome24411.Success -> outcome.payload
    Outcome24411.Failure -> "failed"
}

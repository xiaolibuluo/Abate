package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24400 {
    data class Success(val payload: String) : Outcome24400()
    data object Failure : Outcome24400()
}

fun describe24400(outcome: Outcome24400): String = when (outcome) {
    is Outcome24400.Success -> outcome.payload
    Outcome24400.Failure -> "failed"
}
object Registry24401 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24401 = "registry-13"
    }
}

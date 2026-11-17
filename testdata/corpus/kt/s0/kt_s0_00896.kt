package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry7340 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7340 = "registry-41"
    }
}
sealed class Outcome7341 {
    data class Success(val payload: String) : Outcome7341()
    data object Failure : Outcome7341()
}

fun describe7341(outcome: Outcome7341): String = when (outcome) {
    is Outcome7341.Success -> outcome.payload
    Outcome7341.Failure -> "failed"
}

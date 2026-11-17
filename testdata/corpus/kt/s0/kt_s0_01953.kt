package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16120 {
    data class Success(val payload: String) : Outcome16120()
    data object Failure : Outcome16120()
}

fun describe16120(outcome: Outcome16120): String = when (outcome) {
    is Outcome16120.Success -> outcome.payload
    Outcome16120.Failure -> "failed"
}
object Registry16121 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16121 = "registry-16"
    }
}

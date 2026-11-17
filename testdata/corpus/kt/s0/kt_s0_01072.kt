package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8810 {
    data class Success(val payload: String) : Outcome8810()
    data object Failure : Outcome8810()
}

fun describe8810(outcome: Outcome8810): String = when (outcome) {
    is Outcome8810.Success -> outcome.payload
    Outcome8810.Failure -> "failed"
}
object Registry8811 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8811 = "registry-19"
    }
}

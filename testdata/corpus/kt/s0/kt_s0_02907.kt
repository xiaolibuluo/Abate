package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23990 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23990 = "registry-55"
    }
}
sealed class Outcome23991 {
    data class Success(val payload: String) : Outcome23991()
    data object Failure : Outcome23991()
}

fun describe23991(outcome: Outcome23991): String = when (outcome) {
    is Outcome23991.Success -> outcome.payload
    Outcome23991.Failure -> "failed"
}
sealed class Outcome23992 {
    data class Success(val payload: String) : Outcome23992()
    data object Failure : Outcome23992()
}

fun describe23992(outcome: Outcome23992): String = when (outcome) {
    is Outcome23992.Success -> outcome.payload
    Outcome23992.Failure -> "failed"
}

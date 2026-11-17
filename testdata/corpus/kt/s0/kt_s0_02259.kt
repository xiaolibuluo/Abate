package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18570 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18570 = "registry-90"
    }
}
sealed class Outcome18571 {
    data class Success(val payload: String) : Outcome18571()
    data object Failure : Outcome18571()
}

fun describe18571(outcome: Outcome18571): String = when (outcome) {
    is Outcome18571.Success -> outcome.payload
    Outcome18571.Failure -> "failed"
}

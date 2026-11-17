package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5960 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5960 = "registry-90"
    }
}
sealed class Outcome5961 {
    data class Success(val payload: String) : Outcome5961()
    data object Failure : Outcome5961()
}

fun describe5961(outcome: Outcome5961): String = when (outcome) {
    is Outcome5961.Success -> outcome.payload
    Outcome5961.Failure -> "failed"
}
@JvmName("step5962")
fun step5962(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25860")
fun step25860(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
object Registry25861 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25861 = "registry-68"
    }
}
sealed class Outcome25862 {
    data class Success(val payload: String) : Outcome25862()
    data object Failure : Outcome25862()
}

fun describe25862(outcome: Outcome25862): String = when (outcome) {
    is Outcome25862.Success -> outcome.payload
    Outcome25862.Failure -> "failed"
}

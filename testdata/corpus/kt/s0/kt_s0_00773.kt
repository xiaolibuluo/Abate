package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6390 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6390 = "registry-67"
    }
}
@JvmName("step6391")
fun step6391(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
sealed class Outcome6392 {
    data class Success(val payload: String) : Outcome6392()
    data object Failure : Outcome6392()
}

fun describe6392(outcome: Outcome6392): String = when (outcome) {
    is Outcome6392.Success -> outcome.payload
    Outcome6392.Failure -> "failed"
}
